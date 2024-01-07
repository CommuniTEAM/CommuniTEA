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
	"github.com/swaggest/usecase"
	"github.com/swaggest/usecase/status"
)

type cityInput struct {
	genericInput
	Name  string `json:"name"  nullable:"false"`
	State string `json:"state" nullable:"false"`
}

type stateInput struct {
	State string `maxLength:"2" minLength:"2" path:"state"`
}

type citiesOutput struct {
	Cities []db.LocationsCity `json:"cities"`
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
				ID:    newUUID,
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

func GetCity(dbPool PgxPoolIface) usecase.Interactor {
	response := usecase.NewInteractor(
		func(ctx context.Context, input uuidInput, output *db.LocationsCity) error {
			conn, err := dbPool.Acquire(ctx)
			if err != nil {
				log.Println(fmt.Errorf("could not acquire db connection: %w", err))
				return status.Wrap(fmt.Errorf(internalErrMsg), status.Internal)
			}

			defer conn.Release()

			queries := db.New(conn)

			*output, err = queries.GetCity(ctx, input.ID)
			if err != nil {
				if strings.Contains(err.Error(), "no rows") {
					return status.Wrap(fmt.Errorf("invalid city id"), status.InvalidArgument)
				}
				log.Println(fmt.Errorf("could not get city: %w", err))
				return status.Wrap(fmt.Errorf(internalErrMsg), status.Internal)
			}

			return nil
		})

	response.SetTitle("Get City")
	response.SetDescription("Get the details of a city.")
	response.SetTags("Locations")
	response.SetExpectedErrors(status.InvalidArgument)

	return response
}

// GetAllCitiesInState takes a state code as a query parameter and returns
// a list of all cities within the given state.
func GetAllCitiesInState(dbPool PgxPoolIface) usecase.Interactor {
	response := usecase.NewInteractor(
		func(ctx context.Context, input stateInput, output *citiesOutput) error {
			conn, err := dbPool.Acquire(ctx)
			if err != nil {
				log.Println(fmt.Errorf("could not acquire db connection: %w", err))
				return status.Wrap(fmt.Errorf(internalErrMsg), status.Internal)
			}

			defer conn.Release()

			queries := db.New(conn)

			input.State = strings.ToUpper(input.State)
			output.Cities, err = queries.GetAllCitiesInState(ctx, input.State)
			if err != nil {
				log.Println(fmt.Errorf("could not get all cities in state: %w", err))
				return status.Wrap(fmt.Errorf(internalErrMsg), status.Internal)
			}

			return nil
		})

	response.SetTitle("Get All Cities in a State")
	response.SetDescription("Get a list of all cities in a given state.")
	response.SetTags("Locations")
	response.SetExpectedErrors(status.InvalidArgument)

	return response
}

// GetAllCities takes no input parameters and returns a list of every city
// in the database.
func GetAllCities(dbPool PgxPoolIface) usecase.Interactor {
	response := usecase.NewInteractor(
		func(ctx context.Context, input genericInput, output *citiesOutput) error {
			conn, err := dbPool.Acquire(ctx)
			if err != nil {
				log.Println(fmt.Errorf("could not acquire db connection: %w", err))
				return status.Wrap(fmt.Errorf(internalErrMsg), status.Internal)
			}

			defer conn.Release()

			queries := db.New(conn)

			output.Cities, err = queries.GetAllCities(ctx)
			if err != nil {
				log.Println(fmt.Errorf("could not get all cities: %w", err))
				return status.Wrap(fmt.Errorf(internalErrMsg), status.Internal)
			}

			return nil
		})

	response.SetTitle("Get All Cities")
	response.SetDescription("Get a list of every city in the database.")
	response.SetTags("Locations")

	return response
}
