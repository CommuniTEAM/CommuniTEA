package api

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"

	db "github.com/CommuniTEAM/CommuniTEA/db/sqlc"
	"github.com/google/uuid"
	"github.com/swaggest/usecase"
	"github.com/swaggest/usecase/status"
)

type cityInput struct {
	defaultInput
	StateCode string `json:"state_code" maxLength:"2" minLength:"2"    nullable:"false" pattern:"^(A[KLRZ]|C[AOT]|D[CE]|FL|GA|HI|I[ADLN]|K[SY]|LA|M[ADEINOST]|N[CDEHJMVY]|O[HKR]|PA|RI|S[CD]|T[NX]|UT|V[AT]|W[AIVY])$"`
	CityName  string `json:"city_name"  minLength:"1" nullable:"false" required:"true"`
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
			userData := a.Auth.ValidateJWT(input.AccessToken)

			// If the token was invalid or nonexistent then userData will be nil
			if userData == nil {
				return status.Wrap(errors.New("you must be logged in to perform this action"), status.Unauthenticated)
			}

			// Verify that the user has the 'admin' role
			if userData.Role != adminRole {
				return status.Wrap(errors.New("you do not have permission to perform this action"), status.PermissionDenied)
			}

			conn, err := a.dbConn(ctx)
			if err != nil {
				return err
			}
			defer conn.Release()

			queries := db.New(conn)

			// Check that the city isn't already in the database
			_, err = a.getLocationID(input.CityName, input.StateCode)
			if err == nil {
				// If err is nil then the location already exists
				return status.Wrap(errors.New("location already exists"), status.AlreadyExists)
			}

			newUUID, err := uuid.NewRandom()
			if err != nil {
				log.Println(fmt.Errorf("could not generate new uuid: %w", err))
				return status.Wrap(errors.New(internalErrMsg), status.Internal)
			}

			inputArgs := db.CreateCityParams{
				ID:    newUUID,
				Name:  input.CityName,
				State: input.StateCode,
			}

			*output, err = queries.CreateCity(ctx, inputArgs)

			if err != nil {
				log.Println(fmt.Errorf("failed to create city: %w", err))
				return status.Wrap(errors.New(internalErrMsg), status.Internal)
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
					return status.Wrap(errors.New("no city with that id"), status.NotFound)
				}
				log.Println(fmt.Errorf("could not get city: %w", err))
				return status.Wrap(errors.New(internalErrMsg), status.Internal)
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
	type stateInput struct {
		StateCode string `maxLength:"2" minLength:"2" path:"state-code" pattern:"^(A[KLRZ]|C[AOT]|D[CE]|FL|GA|HI|I[ADLN]|K[SY]|LA|M[ADEINOST]|N[CDEHJMVY]|O[HKR]|PA|RI|S[CD]|T[NX]|UT|V[AT]|W[AIVY])$"`
	}

	response := usecase.NewInteractor(
		func(ctx context.Context, input stateInput, output *citiesOutput) error {
			conn, err := a.dbConn(ctx)
			if err != nil {
				return err
			}
			defer conn.Release()

			queries := db.New(conn)

			input.StateCode = strings.ToUpper(input.StateCode)
			output.Cities, err = queries.GetAllCitiesInState(ctx, input.StateCode)
			if err != nil {
				log.Println(fmt.Errorf("could not get all cities in state: %w", err))
				return status.Wrap(errors.New(internalErrMsg), status.Internal)
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
		func(ctx context.Context, _ defaultInput, output *citiesOutput) error {
			conn, err := a.dbConn(ctx)
			if err != nil {
				return err
			}
			defer conn.Release()

			queries := db.New(conn)

			output.Cities, err = queries.GetAllCities(ctx)
			if err != nil {
				log.Println(fmt.Errorf("could not get all cities: %w", err))
				return status.Wrap(errors.New(internalErrMsg), status.Internal)
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
		Name string `json:"name" minLength:"1" nullable:"false" required:"true"`
	}

	response := usecase.NewInteractor(
		func(ctx context.Context, input cityName, output *db.LocationsCity) error {
			userData := a.Auth.ValidateJWT(input.AccessToken)

			if userData == nil {
				return status.Wrap(errors.New("you must be logged in to perform this action"), status.Unauthenticated)
			}

			if userData.Role != adminRole {
				return status.Wrap(errors.New("you do not have permission to perform this action"), status.PermissionDenied)
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
					return status.Wrap(errors.New("no city with that id"), status.NotFound)
				}
				log.Println(fmt.Errorf("could not get city: %w", err))
				return status.Wrap(errors.New(internalErrMsg), status.Internal)
			}

			// Check that the new city name won't conflict
			_, err = a.getLocationID(input.Name, city.State)
			if err == nil {
				// If err is nil then a city with the new name already exists
				return status.Wrap(errors.New("could not update: a city with that name already exists"), status.AlreadyExists)
			}

			*output, err = queries.UpdateCityName(ctx, db.UpdateCityNameParams{ID: input.ID, Name: input.Name})
			if err != nil {
				log.Println("could not update city name: %w", err)
				return status.Wrap(errors.New(internalErrMsg), status.Internal)
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
			userData := a.Auth.ValidateJWT(input.AccessToken)
			if userData == nil {
				return status.Wrap(errors.New("you must be logged in to perform this action"), status.Unauthenticated)
			}

			if userData.Role != adminRole {
				return status.Wrap(errors.New("you do not have permission to perform this action"), status.PermissionDenied)
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
					return status.Wrap(errors.New("cannot delete a location that is in use by other data"), status.Aborted)
				}
				log.Println(fmt.Errorf("could not delete city: %w", err))
				return status.Wrap(errors.New(internalErrMsg), status.Internal)
			}

			output.Message = successMsg
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
		func(ctx context.Context, _ defaultInput, output *statesOutput) error {
			conn, err := a.dbConn(ctx)
			if err != nil {
				return err
			}
			defer conn.Release()

			queries := db.New(conn)

			output.States, err = queries.GetAllStates(ctx)
			if err != nil {
				log.Println(fmt.Errorf("could not get all states: %w", err))
				return status.Wrap(errors.New(internalErrMsg), status.Internal)
			}

			return nil
		})

	response.SetTitle("Get All States")
	response.SetDescription("Get a list of every US state in the database.")
	response.SetTags("Locations")

	return response
}

// GetIANATimezones takes no input parameters and returns a list of all the
// IANA timezone locations in the database.
func (a *API) GetIANATimezones() usecase.Interactor {
	type timezonesOutput struct {
		Timezones []string `json:"timezones"`
	}

	response := usecase.NewInteractor(
		func(ctx context.Context, _ defaultInput, output *timezonesOutput) error {
			conn, err := a.dbConn(ctx)
			if err != nil {
				return err
			}
			defer conn.Release()

			queries := db.New(conn)

			output.Timezones, err = queries.GetAllTimezones(ctx)
			if err != nil {
				log.Println(fmt.Errorf("could not get all IANA timezone locations: %w", err))
				return status.Wrap(errors.New(internalErrMsg), status.Internal)
			}

			return nil
		})

	response.SetTitle("Get All IANA Timezone Locations")
	response.SetDescription("Get a list of every IANA timezone location in the database.")
	response.SetTags("Locations")

	return response
}

// getLocationID is a helper function that returns the ID of a location, given
// a city name and a state code as strings, if it exists. Any errors returned
// are wrapped.
func (a *API) getLocationID(cityName string, stateCode string) (uuid.UUID, error) {
	conn, err := a.dbConn(context.Background())
	if err != nil {
		return uuid.UUID{}, err
	}
	defer conn.Release()

	queries := db.New(conn)

	locationID, err := queries.GetCityID(context.Background(), db.GetCityIDParams{
		Name:  cityName,
		State: stateCode,
	})
	if err != nil {
		return uuid.UUID{}, status.Wrap(errors.New("location does not exist"), status.InvalidArgument)
	}

	return locationID, nil
}

// getLocationDetails is a helper function that provides the full database
// row for a location, given its ID. Any errors returned are wrapped.
func (a *API) getLocationDetails(locationID uuid.UUID) (db.LocationsCity, error) {
	conn, err := a.dbConn(context.Background())
	if err != nil {
		return db.LocationsCity{}, err
	}
	defer conn.Release()

	queries := db.New(conn)

	locationDetails, err := queries.GetCity(context.Background(), locationID)
	if err != nil {
		log.Println(fmt.Errorf("could not get city: %w", err))
		return db.LocationsCity{}, status.Wrap(errors.New(internalErrMsg), status.Internal)
	}

	return locationDetails, nil
}
