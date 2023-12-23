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
)

type cityInput struct {
	Name string `description:"The name of the city" json:"name"`

	State string `description:"Abbreviated state the city is located within" json:"state"`
}

func CreateCity(dbPool *pgxpool.Pool) usecase.Interactor {
	response := usecase.NewInteractor(func(ctx context.Context, input cityInput, output *db.LocationsCity) error {
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
			return fmt.Errorf("failed to create user: %w", err)
		}

		return nil

		// output.Now = time.Now()
	})

	// Describe use case interactor.

	response.SetTitle("Create Location")

	response.SetDescription("Make a new US city.")

	response.SetTags("Locations")

	response.SetExpectedErrors(status.InvalidArgument)

	return response
}

// ! ONLY A TEMP FUNCTION -- DELETE FOR PROD
func GetCity(dbPool *pgxpool.Pool) pgtype.UUID {
	conn, _ := dbPool.Acquire(context.Background())

	defer conn.Release()

	queries := db.New(conn)

	city, _ := queries.GetCity(context.Background())

	return city
}
