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

type newUserInput struct {
	Role          string `enum:"user,business" json:"role"      nullable:"false"`
	Username      string `json:"username"      nullable:"false"`
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	Email         string `json:"email"`
	Password      string `json:"password"      nullable:"false"`
	LocationCity  string `json:"city"          nullable:"false"`
	LocationState string `json:"state"         nullable:"false"`
}

type loginInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type logoutOutput struct {
	genericOutput
	auth.TokenCookie
}

// UserLogin takes an inputted username and password and, if the credentials
// are valid, returns the user's data along with an authenticating jwt cookie.
func (a *API) UserLogin() usecase.Interactor {
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

			output.ID = userData.ID
			output.Location = userData.Location
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
	response := usecase.NewInteractor(
		func(ctx context.Context, input newUserInput, output *auth.TokenData) error {
			conn, err := a.dbConn(ctx)
			if err != nil {
				return err
			}
			defer conn.Release()

			queries := db.New(conn)

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

			locationID, err := queries.GetCityID(ctx, db.GetCityIDParams{
				Name:  input.LocationCity,
				State: input.LocationState,
			})
			if err != nil {
				return status.Wrap(fmt.Errorf("location does not exist"), status.InvalidArgument)
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

			output.ID = userData.ID
			output.Role = userData.Role
			output.Username = userData.Username
			output.FirstName = userData.FirstName.String
			output.LastName = userData.LastName.String
			output.Location = userData.Location

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

func (a *API) CreateTempAdmin(ctx context.Context) auth.TokenCookie {
	hashPass, err := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("could not hash inputted password: %v", err)
	}

	conn, err := a.dbConn(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Release()

	queries := db.New(conn)

	id := uuid.MustParse("006de673-e8b1-47d3-b3b7-3d266b6452c3")

	location := uuid.MustParse("4c33e0bc-3d43-4e77-aed0-b7aff09bb689")

	admin, err := queries.CreateUser(ctx, db.CreateUserParams{ //nolint: exhaustruct // temp admin has no need for names or emails
		ID:       id,
		Role:     adminRole,
		Username: "admin",
		Password: hashPass,
		Location: location,
	})
	if err != nil {
		log.Panicf("could not create temp admin: %v", err)
	}

	token, err := a.Auth.GenerateNewJWT(&auth.TokenData{
		ID:       admin.ID,
		Role:     admin.Role,
		Username: admin.Username,
		Location: admin.Location}, false)
	if err != nil {
		log.Panicf("could not generate new jwt: %v", err)
	}

	return token.TokenCookie
}
