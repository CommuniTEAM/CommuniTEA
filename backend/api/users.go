package api

import (
	"context"
	"fmt"

	db "github.com/CommuniTEAM/CommuniTEA/db/sqlc"
	"github.com/jackc/pgx/v5"
	"github.com/swaggest/usecase"
	"github.com/swaggest/usecase/status"
)

// Declare input port type.

type userInput struct {
	Name string `json:"name"` // Field tags define parameter location and JSON schema constraints.

	// Field tags of unnamed fields are applied to parent schema.

	// they are optional and can be used to disallow unknown parameters.

	// For non-body params, name tag must be provided explicitly.

	// E.g. here no unknown `query` and `cookie` parameters allowed,

	// unknown `header` params are ok.

	// _ struct{} `query:"_" cookie:"_" additionalProperties:"false"`

}

// Declare output port type.

// type userOutput struct {
// 	Now time.Time `header:"X-Now" json:"-"`

// 	Name string `json:"name"`
// }

// Displays New User Response

// type newUserResponse(user db.User) struct {

// 	return userOutput{

// 		Now: user.Now,

// 		Name: user.Name,

// 	}

// }

func CreateUser() usecase.Interactor {
	// Create use case interactor with references to input/output types and interaction function.

	response := usecase.NewInteractor(func(ctx context.Context, input userInput, output *db.User) error {
		dbCtx := context.Background()

		conn, err := pgx.Connect(dbCtx, "postgresql://admin:secret@postgres/communitea-db")

		if err != nil {
			return fmt.Errorf("failed to connect to DB: %w", err)
		}

		defer conn.Close(dbCtx)

		queries := db.New(conn)

		name := "%s"

		input.Name = fmt.Sprintf(name, input.Name)

		*output, err = queries.CreateUser(ctx, input.Name)

		if err != nil {
			return fmt.Errorf("failed to create user: %w", err)
		}

		return nil

		// output.Now = time.Now()
	})

	// Describe use case interactor.

	response.SetTitle("Create User")

	response.SetDescription("Make a new user (requires a name).")

	response.SetExpectedErrors(status.InvalidArgument)

	return response
}

func GetAllUsers() usecase.Interactor {
	response := usecase.NewInteractor(func(ctx context.Context, _ struct{}, output *[]db.User) error {
		dbCtx := context.Background()

		conn, err := pgx.Connect(dbCtx, "postgresql://admin:secret@postgres/communitea-db")

		if err != nil {
			return fmt.Errorf("failed to connect to DB: %w", err)
		}

		defer conn.Close(dbCtx)

		queries := db.New(conn)

		*output, err = queries.GetUsers(ctx)

		if err != nil {
			return fmt.Errorf("failed to get all user from DB: %w", err)
		}

		return nil
	})

	response.SetTags("Users")

	return response
}
