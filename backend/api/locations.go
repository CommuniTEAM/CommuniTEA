package api

import (
	"context"
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/CommuniTEAM/CommuniTEA/auth"
	db "github.com/CommuniTEAM/CommuniTEA/db/sqlc"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/swaggest/usecase"
	"github.com/swaggest/usecase/status"
)

type cityInput struct {
	AccessToken string `cookie:"bearer-token" json:"-"`
	Name        string `json:"name"           required:"true"`
	State       string `json:"state"          required:"true"`
}

// CreateCity adds a city name and its state code to the database.
// Only accessible to admins.
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

			// Verify that the user has the 'admin' role
			if userData["role"] != "admin" {
				return status.Wrap(fmt.Errorf("you do not have permission to perform this action"), status.PermissionDenied)
			}

			// Verify that input matches required pattern
			input.State = strings.ToUpper(input.State)
			match, err := regexp.MatchString("^[A-Z]{2}$", input.State)
			if err != nil {
				log.Println(fmt.Errorf("could not match regex: %w", err))
				return status.Wrap(fmt.Errorf(internalErrMsg), status.Internal)
			}
			if !match {
				return status.Wrap(fmt.Errorf("invalid state code"), status.InvalidArgument)
			}

			conn, err := dbPool.Acquire(ctx)
			if err != nil {
				log.Println(fmt.Errorf("could not acquire db connection: %w", err))
				return status.Wrap(fmt.Errorf(internalErrMsg), status.Internal)
			}

			defer conn.Release()

			queries := db.New(conn)

			// Check that the city isn't already in the database
			_, err = queries.GetCityID(ctx, db.GetCityIDParams{Name: input.Name, State: input.State})
			if err == nil {
				// If err is nil then the location already exists
				return status.Wrap(fmt.Errorf("location already exists"), status.AlreadyExists)
			}

			newUUID, err := uuid.NewRandom()
			if err != nil {
				log.Println(fmt.Errorf("could not generate new uuid: %w", err))
				return status.Wrap(fmt.Errorf(internalErrMsg), status.Internal)
			}

			inputArgs := db.CreateCityParams{
				ID:    pgtype.UUID{Bytes: newUUID, Valid: true},
				Name:  input.Name,
				State: input.State,
			}

			*output, err = queries.CreateCity(ctx, inputArgs)
			if err != nil {
				if strings.Contains(err.Error(), "fkey") {
					return status.Wrap(fmt.Errorf("invalid state code"), status.InvalidArgument)
				}
				log.Println(fmt.Errorf("failed to create city: %w", err))
				return status.Wrap(fmt.Errorf(internalErrMsg), status.Internal)
			}

			return nil
		})

	response.SetTitle("Create Location")
	response.SetDescription("Add a new US city.")
	response.SetTags("Locations")
	response.SetExpectedErrors(
		status.InvalidArgument,
		status.Unauthenticated,
		status.PermissionDenied,
		status.AlreadyExists,
	)

	return response
}
