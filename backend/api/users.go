package api

import (
	"context"
	"fmt"
	"log"

	"github.com/CommuniTEAM/CommuniTEA/auth"
	db "github.com/CommuniTEAM/CommuniTEA/db/sqlc"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/swaggest/usecase"
	"github.com/swaggest/usecase/status"
	"golang.org/x/crypto/bcrypt"
)

type newUserInput struct {
	Role string `enum:"user, business" json:"role" nullable:"false"`

	Username string `json:"username" nullable:"false"`

	FirstName string `json:"first_name"`

	LastName string `json:"last_name"`

	Email string `json:"email"`

	Password string `json:"password" nullable:"false"`

	LocationCity string `json:"city" nullable:"false"`

	LocationState string `json:"state" nullable:"false"`
}

type loginInput struct {
	Username string `json:"username"`

	Password string `json:"password"`
}

type logoutInput struct {
	Cookie string `cookie:"bearer-token" json:"-"`
}

type logoutOutput struct {
	Message string `json:"message"`

	Cookie string `cookie:"bearer-token,httponly,secure,samesite=strict,path=/,max-age:3600" json:"-"`
}

// UserLogin takes an inputted username and password and, if the credentials

// are valid, returns the user's data along with an authenticating jwt cookie.

func UserLogin(dbPool *pgxpool.Pool) usecase.Interactor {
	response := usecase.NewInteractor(

		func(ctx context.Context, input loginInput, output *auth.TokenData) error {
			conn, err := dbPool.Acquire(ctx)

			if err != nil {
				log.Println(fmt.Errorf("could not acquire db connection: %w", err))

				return status.Wrap(fmt.Errorf("could not process request, please try again"), status.Internal)
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

			userUUID, err := uuid.FromBytes(userData.ID.Bytes[:])

			if err != nil {
				log.Println(fmt.Errorf("could not convert user uuid to uuid type: %w", err))

				return status.Wrap(fmt.Errorf("could not process request, please try again"), status.Internal)
			}

			output.ID = userUUID

			locationUUID, err := uuid.FromBytes(userData.Location.Bytes[:])

			if err != nil {
				log.Println(fmt.Errorf("could not convert location uuid to uuid type: %w", err))

				return status.Wrap(fmt.Errorf("could not process request, please try again"), status.Internal)
			}

			output.Location = locationUUID

			output.Role = userData.Role

			output.Username = userData.Username

			output.FirstName = userData.FirstName.String

			output.LastName = userData.LastName.String

			output, err = auth.GenerateNewJWT(output, false)

			if err != nil {
				log.Println(fmt.Errorf("could not generate new jwt: %w", err))

				return status.Wrap(fmt.Errorf("could not process request, please try again"), status.Internal)
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

func UserLogout() usecase.Interactor {
	response := usecase.NewInteractor(

		func(ctx context.Context, input logoutInput, output *logoutOutput) error {
			userData := auth.ValidateJWT(input.Cookie)

			if userData == nil {
				output.Message = "success"

				output.Cookie = ""

				return nil
			}

			token, err := auth.GenerateNewJWT(

				&auth.TokenData{},

				true,
			)

			if err != nil {
				log.Println(fmt.Errorf("could not generate new jwt: %w", err))

				return status.Wrap(fmt.Errorf("could not process request, please try again"), status.Internal)
			}

			output.Cookie = token.Token

			output.Message = "success"

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

func CreateUser(dbPool *pgxpool.Pool) usecase.Interactor {
	response := usecase.NewInteractor(

		func(ctx context.Context, input newUserInput, output *auth.TokenData) error {
			conn, err := dbPool.Acquire(ctx)

			if err != nil {
				log.Println(fmt.Errorf("could not acquire db connection: %w", err))

				return status.Wrap(fmt.Errorf("could not process request, please try again"), status.Internal)
			}

			defer conn.Release()

			queries := db.New(conn)

			newUUID, err := uuid.NewRandom()

			if err != nil {
				log.Println(fmt.Errorf("could not generate new uuid: %w", err))

				return status.Wrap(fmt.Errorf("could not process request, please try again"), status.Internal)
			}

			hashPass, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)

			if err != nil {
				log.Println(fmt.Errorf("could not hash inputted password: %w", err))

				return status.Wrap(fmt.Errorf("could not process request, please try again"), status.Internal)
			}

			locationInput := db.GetCityParams{Name: input.LocationCity, State: input.LocationState}

			locationID, locationErr := queries.GetCity(ctx, locationInput)

			if locationErr != nil {
				return fmt.Errorf("failed to get locationUUID, invalid city or state: %w", locationErr)
			}

			inputArgs := db.CreateUserParams{

				Column1: pgtype.UUID{Bytes: newUUID, Valid: true},

				Column2: pgtype.Text{String: input.Role, Valid: true},

				Column3: pgtype.Text{String: input.Username, Valid: true},

				Column4: pgtype.Text{String: input.FirstName, Valid: (input.FirstName != "")},

				Column5: pgtype.Text{String: input.LastName, Valid: (input.LastName != "")},

				Column6: pgtype.Text{String: input.Email, Valid: (input.Email != "")},

				Column7: hashPass,

				Column8: pgtype.UUID{Bytes: locationID.Bytes, Valid: (locationErr == nil)},
			}

			userData, err := queries.CreateUser(ctx, inputArgs)

			if err != nil {
				log.Println(fmt.Errorf("failed to create user: %w", err))

				return status.Wrap(fmt.Errorf("could not process request, please try again"), status.Internal)
			}

			output.ID = userData.ID.Bytes

			output.Role = userData.Role

			output.Username = userData.Username

			output.FirstName = userData.FirstName.String

			output.LastName = userData.LastName.String

			output.Location = userData.Location.Bytes

			output, err = auth.GenerateNewJWT(output, false)

			if err != nil {
				log.Println(fmt.Errorf("could not generate new jwt: %w", err))

				return status.Wrap(fmt.Errorf("could not process request, please try again"), status.Internal)
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
