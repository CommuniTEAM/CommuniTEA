package api

import (
	"context"
	"time"

	// "github.com/swaggest/usecase/status"
	db "github.com/CommuniTEAM/CommuniTEA/db/sqlc"
	"github.com/jackc/pgx/v5"
	"github.com/swaggest/usecase"
)

// Declare input port type.
type userInput struct {
	Name string `minLength:"3" path:"name"` // Field tags define parameter location and JSON schema constraints.

	// Field tags of unnamed fields are applied to parent schema.
	// they are optional and can be used to disallow unknown parameters.
	// For non-body params, name tag must be provided explicitly.
	// E.g. here no unknown `query` and `cookie` parameters allowed,
	// unknown `header` params are ok.
	_ struct{} `query:"_" cookie:"_" additionalProperties:"false"`
}

// Declare output port type.
type userOutput struct {
	Now     time.Time `header:"X-Now" json:"-"`
	Message string    `json:"message"`
}

// func CreateUser() usecase.IOInteractorOf[userInput, userOutput] {
// 	// Create use case interactor with references to input/output types and interaction function.
// 	u := usecase.NewInteractor(func(ctx context.Context, input userInput, output *userOutput) error {

// 		output.Message = fmt.Sprintf(msg, input.Name)
// 		output.Now = time.Now()

// 		return nil
// 	})

// 	// Describe use case interactor.
// 	u.SetTitle("Create User")
// 	u.SetDescription("Make a new user (requires a name).")

// 	u.SetExpectedErrors(status.InvalidArgument)

// 	return u
// }

func GetAllUsers() usecase.Interactor {
	response := usecase.NewInteractor(func(ctx context.Context, _ struct{}, output *[]db.User) error {
		dbCtx := context.Background()

		conn, err := pgx.Connect(dbCtx, "postgresql://admin:secret@postgres/communitea-db")
		if err != nil {
			return err
		}
		defer conn.Close(dbCtx)
		queries := db.New(conn)

		*output, err = queries.GetUsers(ctx)
		if err != nil {
			return err
		}
		return nil
	})
	response.SetTags("Users")

	return response
}
