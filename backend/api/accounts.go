package api

import (
	"context"
	"fmt"

	db "github.com/CommuniTEAM/CommuniTEA/db/sqlc"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/swaggest/usecase"
	"github.com/swaggest/usecase/status"
	"golang.org/x/crypto/bcrypt"
)

type newUserInput struct {
	Role      string `json:"role"`
	Username  string `json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Location  string `json:"location"`
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
