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
)

type cityInput struct {
	AccessToken string `cookie:"bearer-token"                                      json:"-"`
	Name        string `description:"The name of the city"                         json:"name"  required:"true"`
	State       string `description:"Abbreviated state the city is located within" json:"state" required:"true"`
}

func CreateCity(dbPool *pgxpool.Pool) usecase.Interactor {
	response := usecase.NewInteractor(
		func(ctx context.Context, input cityInput, output *db.LocationsCity) error {
			// Validate the access token sent with the request, if valid then
			// the user's data will be stored in userData for easy access.
			userData := auth.ValidateJWT(input.AccessToken)

			// If the token was invalid or nonexistent then userData will be nil
			if userData == nil {
				return fmt.Errorf("401: unauthenticated user")
			}

			// For this authentication test, only proceed with the POST request
			// if the user's role is admin
			// Note: currently the admin role is only possible via an UPDATE SQL
			// request to change the role of an existing user through Beekeeper
			if userData["role"] != "admin" {
				return fmt.Errorf("403: user unauthorized")
			}

			// Authenticate complete; now carry out the request
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

			inputArgs := db.CreateCityParams{
				Column1: pgtype.UUID{Bytes: newUUID, Valid: true},
				Column2: pgtype.Text{String: input.Name, Valid: true},
				Column3: pgtype.Text{String: input.State, Valid: true},
			}

			*output, err = queries.CreateCity(ctx, inputArgs)
			if err != nil {
				return fmt.Errorf("failed to create city: %w", err)
			}

			return nil
		})

	response.SetTitle("Create Location")
	response.SetDescription("Make a new US city.")
	response.SetTags("Locations")
	response.SetExpectedErrors(status.InvalidArgument)

	return response
}

// ! ONLY A TEMP FUNCTION -- DELETE FOR PROD
// Returns the first city in the database with the name "string"
func GetCity(dbPool *pgxpool.Pool) pgtype.UUID {
	conn, _ := dbPool.Acquire(context.Background())

	defer conn.Release()

	queries := db.New(conn)

	city, _ := queries.GetCity(context.Background())

	return city
}
