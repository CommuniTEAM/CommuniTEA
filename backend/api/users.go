package api

import (
	"context"
	"fmt"
	"log"

	"github.com/CommuniTEAM/CommuniTEA/auth"
	db "github.com/CommuniTEAM/CommuniTEA/db/sqlc"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/swaggest/usecase"
	"github.com/swaggest/usecase/status"
	"golang.org/x/crypto/bcrypt"
)

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
			output.Username = userData.Username
			output.FirstName = userData.FirstName.String
			output.LastName = userData.LastName.String

			output, err = a.Auth.GenerateNewJWT(output, false)
			if err != nil {
				log.Println(fmt.Errorf("could not generate new jwt: %w", err))
				return status.Wrap(fmt.Errorf(internalErrMsg), status.Internal)
			}

			output.ExpiresIn = 3600
			output.TokenType = "bearer"

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
	type logoutOutput struct {
		genericOutput
		auth.TokenCookie
	}

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
		Role         string `enum:"user,business"         json:"role"      nullable:"false"`
		Username     string `json:"username"              nullable:"false"`
		FirstName    string `json:"first_name"`
		LastName     string `json:"last_name"`
		Email        string `json:"email"`
		Password     string `json:"password"              nullable:"false"`
		PasswordConf string `json:"password_confirmation" nullable:"false"`
	}

	response := usecase.NewInteractor(
		func(ctx context.Context, input newUserInput, output *auth.TokenData) error {
			if input.Role != "user" && input.Role != "business" {
				return status.Wrap(fmt.Errorf("role must be either 'user' or 'business'"), status.InvalidArgument)
			}

			if input.Password != input.PasswordConf {
				return status.Wrap(fmt.Errorf("passwords do not match"), status.InvalidArgument)
			}

			conn, err := a.dbConn(ctx)
			if err != nil {
				return err
			}
			defer conn.Release()

			queries := db.New(conn)

			locationID, err := queries.GetCityID(ctx, db.GetCityIDParams{
				Name:  input.CityName,
				State: input.StateCode,
			})
			if err != nil {
				return status.Wrap(fmt.Errorf("location does not exist"), status.InvalidArgument)
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
			output.Location = locationDetails

			output, err = a.Auth.GenerateNewJWT(output, false)
			if err != nil {
				log.Println(fmt.Errorf("could not generate new jwt: %w", err))
				return status.Wrap(fmt.Errorf(internalErrMsg), status.Internal)
			}

			output.ExpiresIn = 3600
			output.TokenType = "bearer"

			return nil
		})

	response.SetTitle("Create User")
	response.SetDescription("Make a new user account.")
	response.SetTags("Users")
	response.SetExpectedErrors(status.InvalidArgument)

	return response
}
