package api

import (
	"context"
	"fmt"

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
	Role      string `enum:"user, business" json:"role"      nullable:"false"`
	Username  string `json:"username"       nullable:"false"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"       nullable:"false"`
	Location  string `json:"location"       nullable:"false" pattern:"[A-Z][a-zA-Z]+,[ ]?[A-Z]{2}"`
}

type loginInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func UserLogin(dbPool *pgxpool.Pool) usecase.Interactor {
	response := usecase.NewInteractor(func(ctx context.Context, input loginInput, output *auth.TokenData) error {
		conn, err := dbPool.Acquire(ctx)
		if err != nil {
			return fmt.Errorf("could not acquire db connection: %w", err)
		}
		defer conn.Release()

		queries := db.New(conn)

		userData, err := queries.Login(ctx, input.Username)
		if err != nil {
			return fmt.Errorf("could not get user from database: %w", err)
		}

		if bcrypt.CompareHashAndPassword(userData.Password, []byte(input.Password)) != nil {
			return fmt.Errorf("could not authenticate: passwords do not match")
		}

		userUUID, err := uuid.FromBytes(userData.ID.Bytes[:])
		if err != nil {
			return fmt.Errorf("could not convert user uuid to uuid type: %w", err)
		}
		output.ID = userUUID

		locationUUID, err := uuid.FromBytes(userData.Location.Bytes[:])
		if err != nil {
			return fmt.Errorf("could not convert location uuid to uuid type: %w", err)
		}
		output.Location = locationUUID

		output.Role = userData.Role
		output.Username = userData.Username
		output.FirstName = userData.FirstName.String
		output.LastName = userData.LastName.String

		output, err = auth.GenerateNewJWT(output)
		if err != nil {
			return fmt.Errorf("could not generate new JWT: %w", err)
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

func CreateUser(dbPool *pgxpool.Pool) usecase.Interactor {
	response := usecase.NewInteractor(func(ctx context.Context, input newUserInput, output *db.User) error {
		conn, err := dbPool.Acquire(ctx)
		if err != nil {
			return fmt.Errorf("could not acquire db connection: %w", err)
		}
		defer conn.Release()

		queries := db.New(conn)

		newUUID, err := uuid.NewRandom()
		if err != nil {
			return fmt.Errorf("could not generate new uuid: %w", err)
		}

		hashPass, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
		if err != nil {
			return fmt.Errorf("could not hash inputted password: %w", err)
		}

		inputArgs := db.CreateUserParams{
			Column1: pgtype.UUID{Bytes: newUUID, Valid: true},
			Column2: pgtype.Text{String: input.Role, Valid: true},
			Column3: pgtype.Text{String: input.Username, Valid: true},
			Column4: pgtype.Text{String: input.FirstName, Valid: (input.FirstName != "")},
			Column5: pgtype.Text{String: input.LastName, Valid: (input.LastName != "")},
			Column6: pgtype.Text{String: input.Email, Valid: (input.Email != "")},
			Column7: hashPass,
			Column8: GetCity(dbPool),
		}

		*output, err = queries.CreateUser(ctx, inputArgs)
		if err != nil {
			return fmt.Errorf("failed to create user: %w", err)
		}

		return nil
	})

	response.SetTitle("Create User")
	response.SetDescription("Make a new user account.")
	response.SetTags("Accounts")
	response.SetExpectedErrors(status.InvalidArgument)

	return response
}
