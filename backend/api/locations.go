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
	defaultInput

	Name string `json:"name" nullable:"false"`

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

func (a *API) CreateCity() usecase.Interactor {
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

			if userData["role"] != adminRole {
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

			conn, err := a.dbConn(ctx)

			if err != nil {
				return err
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

				ID: newUUID,

				Name: input.Name,

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

// GetCity takes a location's uuid as a query parameter and returns its

// details in the response body.

func (a *API) GetCity() usecase.Interactor {
	response := usecase.NewInteractor(

		func(ctx context.Context, input uuidInput, output *db.LocationsCity) error {
			conn, err := a.dbConn(ctx)

			if err != nil {
				return err
			}

			defer conn.Release()

			queries := db.New(conn)

			*output, err = queries.GetCity(ctx, input.ID)

			if err != nil {
				if strings.Contains(err.Error(), "no rows") {
					return status.Wrap(fmt.Errorf("no city with that id"), status.NotFound)
				}

				log.Println(fmt.Errorf("could not get city: %w", err))

				return status.Wrap(fmt.Errorf(internalErrMsg), status.Internal)
			}

			return nil
		})

	response.SetTitle("Get Location")

	response.SetDescription("Get the details of a city.")

	response.SetTags("Locations")

	response.SetExpectedErrors(status.InvalidArgument, status.NotFound)

	return response
}

// GetAllCitiesInState takes a state code as a query parameter and returns

// a list of all cities within the given state.

func (a *API) GetAllCitiesInState() usecase.Interactor {
	response := usecase.NewInteractor(

		func(ctx context.Context, input stateInput, output *citiesOutput) error {
			conn, err := a.dbConn(ctx)

			if err != nil {
				return err
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

func (a *API) GetAllCities() usecase.Interactor {
	response := usecase.NewInteractor(

		func(ctx context.Context, input defaultInput, output *citiesOutput) error {
			conn, err := a.dbConn(ctx)

			if err != nil {
				return err
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

	response.SetTitle("Get All Locations")

	response.SetDescription("Get a list of every city in the database.")

	response.SetTags("Locations")

	return response
}

// UpdateCity takes a location's uuid as a query parameter and a new city name

// in the request body, then returns the updated location details in the

// response body. Only accessible to admins.

func (a *API) UpdateCity() usecase.Interactor {
	type cityName struct {
		uuidInput

		Name string `json:"name" nullable:"false"`
	}

	response := usecase.NewInteractor(

		func(ctx context.Context, input cityName, output *db.LocationsCity) error {
			userData := auth.ValidateJWT(input.AccessToken)

			if userData == nil {
				return status.Wrap(fmt.Errorf("you must be logged in to perform this action"), status.Unauthenticated)
			}

			if userData["role"] != adminRole {
				return status.Wrap(fmt.Errorf("you do not have permission to perform this action"), status.PermissionDenied)
			}

			conn, err := a.dbConn(ctx)

			if err != nil {
				return err
			}

			defer conn.Release()

			queries := db.New(conn)

			// Get original city details

			city, err := queries.GetCity(ctx, input.ID)

			if err != nil {
				if strings.Contains(err.Error(), "no rows") {
					return status.Wrap(fmt.Errorf("no city with that id"), status.NotFound)
				}

				log.Println(fmt.Errorf("could not get city: %w", err))

				return status.Wrap(fmt.Errorf(internalErrMsg), status.Internal)
			}

			// Check that the new city name won't conflict

			_, err = queries.GetCityID(ctx, db.GetCityIDParams{Name: input.Name, State: city.State})

			if err == nil {
				// If err is nil then a city with the new name already exists

				return status.Wrap(fmt.Errorf("could not update: a city with that name already exists"), status.AlreadyExists)
			}

			*output, err = queries.UpdateCityName(ctx, db.UpdateCityNameParams{ID: input.ID, Name: input.Name})

			if err != nil {
				log.Println("could not update city name: %w", err)

				return status.Wrap(fmt.Errorf(internalErrMsg), status.Internal)
			}

			return nil
		})

	response.SetTitle("Update Location")

	response.SetDescription("Change the name of an existing location.")

	response.SetTags("Locations")

	response.SetExpectedErrors(

		status.InvalidArgument,

		status.Unauthenticated,

		status.PermissionDenied,

		status.AlreadyExists,

		status.NotFound,
	)

	return response
}

// DeleteCity takes a city uuid input through query parameters and returns

// a success message if the city does not exist after running the deletion

// query (404 is never returned). Only accessible to admins.

func (a *API) DeleteCity() usecase.Interactor {
	response := usecase.NewInteractor(

		func(ctx context.Context, input uuidInput, output *genericOutput) error {
			userData := auth.ValidateJWT(input.AccessToken)

			if userData == nil {
				return status.Wrap(fmt.Errorf("you must be logged in to perform this action"), status.Unauthenticated)
			}

			if userData["role"] != adminRole {
				return status.Wrap(fmt.Errorf("you do not have permission to perform this action"), status.PermissionDenied)
			}

			conn, err := a.dbConn(ctx)

			if err != nil {
				return err
			}

			defer conn.Release()

			queries := db.New(conn)

			err = queries.DeleteCity(ctx, input.ID)

			if err != nil {
				if strings.Contains(err.Error(), "fkey") {
					return status.Wrap(fmt.Errorf("cannot delete a location that is in use by other data"), status.Aborted)
				}

				log.Println(fmt.Errorf("could not delete city: %w", err))

				return status.Wrap(fmt.Errorf(internalErrMsg), status.Internal)
			}

			output.Message = "success: location deleted"

			return nil
		})

	response.SetTitle("Delete Location")

	response.SetDescription("Remove a location.")

	response.SetTags("Locations")

	response.SetExpectedErrors(

		status.InvalidArgument,

		status.Unauthenticated,

		status.PermissionDenied,

		status.Aborted,
	)

	return response
}

// GetAllStates takes no input parameters and returns a list of every state

// in the database.

func (a *API) GetAllStates() usecase.Interactor {
	type statesOutput struct {
		States []db.LocationsState `json:"states"`
	}

	response := usecase.NewInteractor(

		func(ctx context.Context, input defaultInput, output *statesOutput) error {
			conn, err := a.dbConn(ctx)

			if err != nil {
				return err
			}

			defer conn.Release()

			queries := db.New(conn)

			output.States, err = queries.GetAllStates(ctx)

			if err != nil {
				log.Println(fmt.Errorf("could not get all states: %w", err))

				return status.Wrap(fmt.Errorf(internalErrMsg), status.Internal)
			}

			return nil
		})

	response.SetTitle("Get All States")

	response.SetDescription("Get a list of every US state in the database.")

	response.SetTags("Locations")

	return response
}
