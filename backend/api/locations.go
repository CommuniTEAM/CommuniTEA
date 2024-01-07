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
)

type cityInput struct {
	AccessToken string `cookie:"bearer-token" json:"-"`

	Name string `json:"name" required:"true"`

	State string `json:"state" required:"true"`
}

// Example endpoint that is not yet fully-functional, but serves to

// showcase authentication in action.

// TEA-62 will flesh it out with data validation and error handling

func CreateCity(dbPool PgxPoolIface) usecase.Interactor {
	response := usecase.NewInteractor(

		func(ctx context.Context, input cityInput, output *db.LocationsCity) error {
			// Validate the access token sent with the request, if valid then

			// the user's data will be stored in userData for easy access.

			userData := auth.ValidateJWT(input.AccessToken)

			// If the token was invalid or nonexistent then userData will be nil

			if userData == nil {
				return status.Wrap(fmt.Errorf("you must be logged in to perform this action"), status.Unauthenticated)
			}

			// For this authentication test, only proceed with the POST request

			// if the user's role is admin

			// Note: currently the admin role is only possible via an UPDATE SQL

			// request to change the role of an existing user through Beekeeper

			if userData["role"] != "admin" {
				return status.Wrap(fmt.Errorf("you do not have permission to perform this action"), status.PermissionDenied)
			}

			// Authenticate complete; now carry out the request

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

			inputArgs := db.CreateCityParams{

				ID: pgtype.UUID{Bytes: newUUID, Valid: true},

				Name: input.Name,

				State: input.State,
			}

			*output, err = queries.CreateCity(ctx, inputArgs)

			if err != nil {
				log.Println(fmt.Errorf("failed to create city: %w", err))

				return status.Wrap(fmt.Errorf("could not process request, please try again"), status.Internal)
			}

			return nil
		})

	response.SetTitle("Create Location")

	response.SetDescription("Make a new US city.")

	response.SetTags("Locations")

	response.SetExpectedErrors(

		status.InvalidArgument,

		status.Unauthenticated,

		status.PermissionDenied,
	)

	return response
}

// ! ONLY A TEMP FUNCTION -- DELETE FOR PROD

// Returns the first city in the database with the name "string"

// func GetCity(dbPool PgxPoolIface) pgtype.UUID {

// 	conn, _ := dbPool.Acquire(context.Background())

// 	defer conn.Release()

// 	queries := db.New(conn)

// 	city, _ := queries.GetCity(context.Background())

// 	return city

// }
