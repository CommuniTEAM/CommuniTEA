package api

import (
	"context"
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/CommuniTEAM/CommuniTEA/auth"
	db "github.com/CommuniTEAM/CommuniTEA/db/sqlc"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/swaggest/usecase"
	"github.com/swaggest/usecase/status"
	"golang.org/x/crypto/bcrypt"
)

type userInput struct {
	defaultInput
	ID uuid.UUID `nullable:"false" path:"id" required:"true"`
}

type logoutOutput struct {
	genericOutput
	auth.TokenCookie
}

type userOutput struct {
	ID        uuid.UUID        `json:"id"                   required:"true"`
	Role      string           `json:"role"                 required:"true"`
	Username  string           `json:"username"             required:"true"`
	FirstName string           `json:"first_name,omitempty"`
	LastName  string           `json:"last_name,omitempty"`
	Email     string           `json:"email,omitempty"`
	Location  db.LocationsCity `json:"location"             required:"true"`
}

// UserLogin takes an inputted username and password and, if the credentials
// are valid, returns the user's data along with an authenticating jwt cookie.
func (a *API) UserLogin() usecase.Interactor {
	type loginInput struct {
		Username string `json:"username" nullable:"false" required:"true"`
		Password string `json:"password" nullable:"false" required:"true"`
	}

	response := usecase.NewInteractor(
		func(ctx context.Context, input loginInput, output *auth.TokenData) error {
			conn, err := a.dbConn(ctx)
			if err != nil {
				return err
			}
			defer conn.Release()

			queries := db.New(conn)

			userData, err := queries.Login(ctx, input.Username)
			if err != nil {
				log.Println(fmt.Errorf("could not get user from database: %w", err))
				return status.Wrap(fmt.Errorf("invalid credentials"), status.InvalidArgument)
			}

			if bcrypt.CompareHashAndPassword(userData.Password, []byte(input.Password)) != nil {
				return status.Wrap(fmt.Errorf("invalid credentials"), status.InvalidArgument)
			}

			locationDetails, err := a.getLocationDetails(userData.Location)
			if err != nil {
				log.Println(fmt.Errorf("could not get location details: %w", err))
				return status.Wrap(fmt.Errorf(internalErrMsg), status.Internal)
			}

			output.ID = userData.ID
			output.Location = locationDetails
			output.Role = userData.Role
			output.Email = userData.Email.String
			output.Username = userData.Username
			output.FirstName = userData.FirstName.String
			output.LastName = userData.LastName.String

			output, err = a.Auth.GenerateNewJWT(output, false)
			if err != nil {
				log.Println(fmt.Errorf("could not generate new jwt: %w", err))
				return status.Wrap(fmt.Errorf(internalErrMsg), status.Internal)
			}

			output.ExpiresIn = 3600
			return nil
		})

	response.SetTitle("Login")
	response.SetDescription("Log in to an existing account.")
	response.SetTags("Auth")
	response.SetExpectedErrors(status.InvalidArgument)

	return response
}

// UserLogout takes in an optional auth cookie and, if valid, returns a new
// auth cookie with an expiry time one hour in the past, rendering it invalid.
// If the passed-in cookie is already invalid or there is no cookie, an empty
// 200 response is returned as the user is already not logged in.
func (a *API) UserLogout() usecase.Interactor {
	response := usecase.NewInteractor(
		func(ctx context.Context, input defaultInput, output *logoutOutput) error {
			userData := a.Auth.ValidateJWT(input.AccessToken)
			if userData == nil {
				output.Message = successMsg
				output.Token = ""
				return nil
			}

			token, err := a.Auth.GenerateNewJWT(
				&auth.TokenData{},
				true,
			)
			if err != nil {
				log.Println(fmt.Errorf("could not generate new jwt: %w", err))
				return status.Wrap(fmt.Errorf(internalErrMsg), status.Internal)
			}

			output.Token = token.Token
			output.Message = successMsg

			return nil
		})

	response.SetTitle("Logout")
	response.SetDescription("Log out of an account.")
	response.SetTags("Auth")
	response.SetExpectedErrors(status.InvalidArgument)

	return response
}

// CreateUser takes in a user's information, saves it to the database, then
// logs them in and returns the user's data and an authenticating jwt cookie.
func (a *API) CreateUser() usecase.Interactor {
	type newUserInput struct {
		cityInput
		Role         string `enum:"user,business"         json:"role"      nullable:"false" required:"true"`
		Username     string `json:"username"              nullable:"false" required:"true"`
		FirstName    string `json:"first_name"`
		LastName     string `json:"last_name"`
		Email        string `json:"email"`
		Password     string `json:"password"              nullable:"false" required:"true"`
		PasswordConf string `json:"password_confirmation" nullable:"false" required:"true"`
	}

	response := usecase.NewInteractor(
		func(ctx context.Context, input newUserInput, output *auth.TokenData) error {
			if input.Password != input.PasswordConf {
				return status.Wrap(fmt.Errorf("passwords do not match"), status.InvalidArgument)
			}

			conn, err := a.dbConn(ctx)
			if err != nil {
				return err
			}
			defer conn.Release()

			queries := db.New(conn)

			_, err = queries.GetUserByUsername(ctx, input.Username)
			if err == nil {
				return status.Wrap(fmt.Errorf("username taken"), status.AlreadyExists)
			}

			if input.Email != "" {
				emailErr := verifyEmail(input.Email, "", queries)
				if emailErr != nil {
					return emailErr
				}
			}

			locationID, err := a.getLocationID(input.CityName, input.StateCode)
			if err != nil {
				return err
			}

			newUUID, err := uuid.NewRandom()
			if err != nil {
				log.Println(fmt.Errorf("could not generate new uuid: %w", err))
				return status.Wrap(fmt.Errorf(internalErrMsg), status.Internal)
			}

			hashPass, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
			if err != nil {
				log.Println(fmt.Errorf("could not hash inputted password: %w", err))
				return status.Wrap(fmt.Errorf(internalErrMsg), status.Internal)
			}

			inputArgs := db.CreateUserParams{
				ID:        newUUID,
				Role:      input.Role,
				Username:  input.Username,
				FirstName: pgtype.Text{String: input.FirstName, Valid: (input.FirstName != "")},
				LastName:  pgtype.Text{String: input.LastName, Valid: (input.LastName != "")},
				Email:     pgtype.Text{String: input.Email, Valid: (input.Email != "")},
				Password:  hashPass,
				Location:  locationID,
			}

			userData, err := queries.CreateUser(ctx, inputArgs)
			if err != nil {
				log.Println(fmt.Errorf("failed to create user: %w", err))
				return status.Wrap(fmt.Errorf(internalErrMsg), status.Internal)
			}

			locationDetails, err := a.getLocationDetails(locationID)
			if err != nil {
				log.Println(fmt.Errorf("could not get location details: %w", err))
				return status.Wrap(fmt.Errorf(internalErrMsg), status.Internal)
			}

			output.ID = userData.ID
			output.Role = userData.Role
			output.Username = userData.Username
			output.FirstName = userData.FirstName.String
			output.LastName = userData.LastName.String
			output.Email = userData.Email.String
			output.Location = locationDetails

			output, err = a.Auth.GenerateNewJWT(output, false)
			if err != nil {
				log.Println(fmt.Errorf("could not generate new jwt: %w", err))
				return status.Wrap(fmt.Errorf(internalErrMsg), status.Internal)
			}

			output.ExpiresIn = 3600
			return nil
		})

	response.SetTitle("Create User")
	response.SetDescription("Make a new user account.")
	response.SetTags("Users")
	response.SetExpectedErrors(status.InvalidArgument, status.AlreadyExists)

	return response
}

func (a *API) GetUser() usecase.Interactor {
	response := usecase.NewInteractor(
		func(ctx context.Context, input userInput, output *userOutput) error {
			conn, err := a.dbConn(ctx)
			if err != nil {
				return err
			}
			defer conn.Release()

			queries := db.New(conn)

			userDetails, err := queries.GetUserByID(ctx, input.ID)
			if err != nil {
				if strings.Contains(err.Error(), "no rows") {
					return status.Wrap(fmt.Errorf("no user with that id"), status.NotFound)
				}
				log.Println(fmt.Errorf("could not get user details: %w", err))
				return status.Wrap(fmt.Errorf(internalErrMsg), status.Internal)
			}

			output.ID = userDetails.ID
			output.Username = userDetails.Username
			output.FirstName = userDetails.FirstName.String
			output.LastName = userDetails.LastName.String
			output.Email = userDetails.Email.String
			output.Role = userDetails.Role
			output.Location, err = a.getLocationDetails(userDetails.Location)
			if err != nil {
				return err
			}

			return nil
		})

	response.SetTitle("Get User")
	response.SetDescription("Get the details of a user.")
	response.SetTags("Users")
	response.SetExpectedErrors(status.InvalidArgument)

	return response
}

// UpdateUser accepts a first name, last name, email, role, and location as
// optional inputs for a logged-in user and updates the database with the new
// variables. Returns the updated user details and a new auth token.
func (a *API) UpdateUser() usecase.Interactor {
	type userUpdateInput struct {
		defaultInput
		ID        uuid.UUID `path:"id"`
		FirstName string    `json:"first_name"`
		LastName  string    `json:"last_name"`
		Email     string    `json:"email"`
		Role      string    `enum:"user,business" json:"role"`
		StateCode string    `json:"state_code"    maxLength:"2" minLength:"2"    nullable:"false" pattern:"^(A[KLRZ]|C[AOT]|D[CE]|FL|GA|HI|I[ADLN]|K[SY]|LA|M[ADEINOST]|N[CDEHJMVY]|O[HKR]|PA|RI|S[CD]|T[NX]|UT|V[AT]|W[AIVY])$"`
		CityName  string    `json:"city_name"     minLength:"1" nullable:"false"`
	}

	response := usecase.NewInteractor(
		func(ctx context.Context, input userUpdateInput, output *auth.TokenData) error {
			userData := a.Auth.ValidateJWT(input.AccessToken)
			if userData == nil {
				return status.Wrap(fmt.Errorf("you must be logged in to perform this action"), status.Unauthenticated)
			}

			if userData.ID != input.ID {
				return status.Wrap(fmt.Errorf("you do not have permission to perform this action"), status.PermissionDenied)
			}

			conn, err := a.dbConn(ctx)
			if err != nil {
				return err
			}
			defer conn.Release()

			queries := db.New(conn)

			// Check for differences in input vs userData

			role := userData.Role
			if input.Role != role && input.Role != "" {
				role = input.Role
			}

			firstName := userData.FirstName
			if input.FirstName != "" {
				firstName = input.FirstName
			}

			lastName := userData.LastName
			if input.LastName != "" {
				lastName = input.LastName
			}

			email := userData.Email
			if input.Email != "" {
				emailErr := verifyEmail(input.Email, input.ID.String(), queries)
				if emailErr != nil {
					return emailErr
				}
			}

			location := userData.Location

			if input.CityName == location.Name && input.StateCode == location.State {
				location = userData.Location
			} else if input.CityName != "" || input.StateCode != "" {
				cityName := input.CityName
				stateCode := input.StateCode
				if cityName == "" {
					cityName = location.Name
				}
				if stateCode == "" {
					stateCode = location.State
				}

				locationID, dbErr := a.getLocationID(cityName, stateCode)
				if dbErr != nil {
					return status.Wrap(fmt.Errorf("location does not exist"), status.InvalidArgument)
				}
				location.ID = locationID
				location.Name = cityName
				location.State = stateCode
			}

			// Send updated data to database

			inputArgs := db.UpdateUserParams{
				ID:        userData.ID,
				Role:      role,
				FirstName: pgtype.Text{String: firstName, Valid: (firstName != "")},
				LastName:  pgtype.Text{String: lastName, Valid: (lastName != "")},
				Email:     pgtype.Text{String: email, Valid: (email != "")},
				Location:  location.ID,
			}

			updatedUser, err := queries.UpdateUser(ctx, inputArgs)
			if err != nil {
				log.Println(fmt.Errorf("failed to update user: %w", err))
				return status.Wrap(fmt.Errorf(internalErrMsg), status.Internal)
			}

			output.ID = updatedUser.ID
			output.Role = updatedUser.Role
			output.Username = updatedUser.Username
			output.FirstName = updatedUser.FirstName.String
			output.LastName = updatedUser.LastName.String
			output.Email = updatedUser.Email.String
			output.Location = location

			output, err = a.Auth.GenerateNewJWT(output, false)
			if err != nil {
				log.Println(fmt.Errorf("could not generate new jwt: %w", err))
				return status.Wrap(fmt.Errorf(internalErrMsg), status.Internal)
			}

			output.ExpiresIn = 3600
			return nil
		})

	response.SetTitle("Update User")
	response.SetDescription("Change one or some variables of your user account.")
	response.SetTags("Users")
	response.SetExpectedErrors(status.InvalidArgument, status.Unauthenticated, status.PermissionDenied)

	return response
}

func (a *API) ChangePassword() usecase.Interactor {
	type passwordInput struct {
		defaultInput
		ID              uuid.UUID `path:"id"`
		OldPassword     string    `json:"old_password"      nullable:"false" required:"true"`
		NewPassword     string    `json:"new_password"      nullable:"false" required:"true"`
		NewPasswordConf string    `json:"new_password_conf" nullable:"false" required:"true"`
	}
	response := usecase.NewInteractor(
		func(ctx context.Context, input passwordInput, output *genericOutput) error {
			userData := a.Auth.ValidateJWT(input.AccessToken)
			if userData == nil {
				return status.Wrap(fmt.Errorf("you must be logged in to perform this action"), status.Unauthenticated)
			}

			if userData.ID != input.ID {
				return status.Wrap(fmt.Errorf("you do not have permission to perform this action"), status.PermissionDenied)
			}

			if input.NewPassword != input.NewPasswordConf {
				return status.Wrap(fmt.Errorf("passwords do not match"), status.InvalidArgument)
			}

			conn, err := a.dbConn(ctx)
			if err != nil {
				return err
			}
			defer conn.Release()

			queries := db.New(conn)

			_, err = queries.Login(ctx, userData.Username)
			if err != nil {
				log.Println(fmt.Errorf("could not get user from database: %w", err))
				return status.Wrap(fmt.Errorf("invalid credentials"), status.InvalidArgument)
			}

			hashedPass, err := bcrypt.GenerateFromPassword([]byte(input.NewPassword), bcrypt.DefaultCost)
			if err != nil {
				log.Println(fmt.Errorf("could not hash inputted password: %w", err))
				return status.Wrap(fmt.Errorf(internalErrMsg), status.Internal)
			}

			err = queries.ChangePassword(ctx, db.ChangePasswordParams{ID: userData.ID, Password: hashedPass})
			if err != nil {
				log.Println(fmt.Errorf("could not change password: %w", err))
				return status.Wrap(fmt.Errorf(internalErrMsg), status.Internal)
			}

			output.Message = successMsg
			return nil
		})

	response.SetTitle("Change Password")
	response.SetDescription("Change your account password.")
	response.SetTags("Users")
	response.SetExpectedErrors(status.InvalidArgument, status.Unauthenticated, status.PermissionDenied)

	return response
}

// PromoteToAdmin takes a user ID and, if the requester is a logged-in admin,
// promotes the user with the given ID to the admin role.
func (a *API) PromoteToAdmin() usecase.Interactor {
	response := usecase.NewInteractor(
		func(ctx context.Context, input userInput, output *userOutput) error {
			userData := a.Auth.ValidateJWT(input.AccessToken)

			if userData == nil {
				return status.Wrap(fmt.Errorf("you must be logged in to perform this action"), status.Unauthenticated)
			}

			if userData.Role != adminRole {
				return status.Wrap(fmt.Errorf("you do not have permission to perform this action"), status.PermissionDenied)
			}

			conn, err := a.dbConn(ctx)
			if err != nil {
				return err
			}
			defer conn.Release()

			queries := db.New(conn)

			user, err := queries.GetUserByID(ctx, input.ID)
			if err != nil {
				log.Println(fmt.Errorf("could not get user: %w", err))
				return status.Wrap(fmt.Errorf("user does not exist"), status.InvalidArgument)
			}

			promotedUser, err := queries.PromoteToAdmin(ctx, user.ID)
			if err != nil {
				log.Println(fmt.Errorf("could not promote user to admin role: %w", err))
				return status.Wrap(fmt.Errorf(internalErrMsg), status.Internal)
			}

			output.ID = promotedUser.ID
			output.Role = promotedUser.Role
			output.Username = promotedUser.Username
			output.FirstName = promotedUser.FirstName.String
			output.LastName = promotedUser.LastName.String
			output.Email = promotedUser.Email.String
			output.Location, err = a.getLocationDetails(promotedUser.Location)
			if err != nil {
				log.Println(fmt.Errorf("could not get location details: %w", err))
				return status.Wrap(fmt.Errorf(internalErrMsg), status.Internal)
			}

			return nil
		})
	response.SetTitle("Promote to Admin")
	response.SetDescription("Promote a user account to the \"admin\" role.")
	response.SetTags("Users")
	response.SetExpectedErrors(status.InvalidArgument, status.Unauthenticated, status.PermissionDenied)

	return response
}

// DeleteUser takes a user ID and deletes the associated account from the
// database IF the request comes from the user with the same ID.
func (a *API) DeleteUser() usecase.Interactor {
	type deleteUserInput struct {
		defaultInput
		ID uuid.UUID `path:"id"`
	}

	response := usecase.NewInteractor(
		func(ctx context.Context, input deleteUserInput, output *logoutOutput) error {
			userData := a.Auth.ValidateJWT(input.AccessToken)
			if userData == nil {
				return status.Wrap(fmt.Errorf("you must be logged in to perform this action"), status.Unauthenticated)
			}

			if userData.ID != input.ID {
				return status.Wrap(fmt.Errorf("you do not have permission to perform this action"), status.PermissionDenied)
			}

			conn, err := a.dbConn(ctx)
			if err != nil {
				return err
			}
			defer conn.Release()

			queries := db.New(conn)

			err = queries.DeleteUser(ctx, input.ID)
			if err != nil {
				log.Println(fmt.Errorf("could not delete user: %w", err))
				return status.Wrap(fmt.Errorf(internalErrMsg), status.Internal)
			}

			token, err := a.Auth.GenerateNewJWT(
				&auth.TokenData{},
				true,
			)
			if err != nil {
				log.Println(fmt.Errorf("could not generate new jwt: %w", err))
				return status.Wrap(fmt.Errorf(internalErrMsg), status.Internal)
			}

			output.Token = token.Token
			output.Message = successMsg

			return nil
		})

	response.SetTitle("Delete User")
	response.SetDescription("Delete your user account.")
	response.SetTags("Users")
	response.SetExpectedErrors(
		status.InvalidArgument,
		status.Unauthenticated,
		status.PermissionDenied,
	)

	return response
}

// verifyEmail is a helper function that checks a given string against email
// regex, then checks it against the database to verify that it is not in use
// by another user. Errors are returned wrapped.
func verifyEmail(email string, userID string, queries *db.Queries) error {
	match, regexpErr := regexp.MatchString("(^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+.[a-zA-Z0-9-.]+$)", email)
	if regexpErr != nil {
		log.Println(fmt.Errorf("could not match regex: %w", regexpErr))
		return status.Wrap(fmt.Errorf(internalErrMsg), status.Internal)
	}
	if !match {
		return status.Wrap(fmt.Errorf("invalid email address"), status.InvalidArgument)
	}

	userData, err := queries.GetUserByEmail(context.Background(), pgtype.Text{String: email, Valid: true})
	if err != nil {
		if strings.Contains(err.Error(), "no rows") {
			return nil
		}
		log.Println(fmt.Errorf("failed to get user by email: %w", err))
		return status.Wrap(fmt.Errorf(internalErrMsg), status.Internal)
	}

	if userID != "" || userID != userData.ID.String() {
		return status.Wrap(fmt.Errorf("email already in use"), status.AlreadyExists)
	}

	return nil
}
