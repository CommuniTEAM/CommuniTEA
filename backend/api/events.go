package api

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	db "github.com/CommuniTEAM/CommuniTEA/db/sqlc"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/swaggest/usecase"
	"github.com/swaggest/usecase/status"
)

const timeFormat = "01-02-2006 15:04"

type eventOutput struct {
	ID               uuid.UUID     `json:"id"`
	Name             string        `json:"name"`
	Host             uuid.UUID     `json:"host"`
	Location         eventLocation `json:"location"`
	StartTime        string        `example:"01-02-2006 15:04 MST" json:"start_time"`
	EndTime          string        `example:"01-02-2006 15:04 MST" json:"end_time"`
	HTMLDescription  string        `json:"html_description"`
	Rsvps            bool          `json:"rsvps"`
	Capacity         int32         `json:"capacity"`
	TimezoneLocation string        `json:"timezone_location"`
	Visible          bool          `json:"visible"`
}

type eventLocation struct {
	Name          string    `example:"The Tea House" json:"name"`
	StreetAddress string    `example:"123 45th St."  json:"street_address"`
	City          string    `example:"Los Angeles"   json:"city"`
	CityID        uuid.UUID `json:"city_id"`
	StateCode     string    `example:"CA"            json:"state_code"`
	Zipcode       string    `example:"12345"         json:"zipcode"`
}

func (a *API) GetEvent() usecase.Interactor {
	response := usecase.NewInteractor(
		func(ctx context.Context, input uuidInput, output *eventOutput) error {
			conn, err := a.dbConn(ctx)
			if err != nil {
				return err
			}
			defer conn.Release()

			queries := db.New(conn)

			eventDetails, err := queries.GetEventByID(ctx, input.ID)
			if err != nil {
				if strings.Contains(err.Error(), "no rows") {
					return status.Wrap(errors.New("event not found"), status.NotFound)
				}
				log.Println(fmt.Errorf("error getting event by id: %w", err))
				return status.Wrap(errors.New(internalErrMsg), status.Internal)
			}

			eventCity, err := a.getLocationDetails(eventDetails.City)
			if err != nil {
				return err
			}

			output.ID = eventDetails.ID
			output.Name = eventDetails.Name
			output.Host = eventDetails.Host
			output.Location.Name = eventDetails.LocationName.String
			output.Location.StreetAddress = eventDetails.StreetAddress
			output.Location.CityID = eventDetails.City
			output.Location.City = eventCity.Name
			output.Location.StateCode = eventCity.State
			output.Location.Zipcode = eventDetails.Zipcode
			output.HTMLDescription = eventDetails.HtmlDescription.String
			output.Rsvps = eventDetails.Rsvps
			output.Capacity = eventDetails.Capacity.Int32
			output.TimezoneLocation = eventDetails.TimezoneLocation
			output.Visible = eventDetails.Visible

			output.StartTime, err = getTimeWithTimezone(eventDetails.TimezoneLocation, eventDetails.StartTime.Time)
			if err != nil {
				return err
			}

			output.EndTime, err = getTimeWithTimezone(eventDetails.TimezoneLocation, eventDetails.EndTime.Time)
			if err != nil {
				return err
			}

			return nil
		})

	response.SetTitle("Get Event")
	response.SetDescription("Get the details of an event.")
	response.SetTags("Events")
	response.SetExpectedErrors(status.InvalidArgument)

	return response
}

func (a *API) CreateEvent() usecase.Interactor {
	type eventInput struct {
		cityInput
		Name             string    `json:"name"                    nullable:"false"`
		Host             uuid.UUID `json:"host"                    nullable:"false"`
		LocationName     string    `json:"location_name,omitempty"`
		StreetAddress    string    `json:"street_address"          nullable:"false"`
		Zipcode          string    `json:"zipcode"                 nullable:"false"         pattern:"^[0-9]{5}$"`
		StartTime        string    `example:"MM-DD-YYYY HH:MM"     json:"start_time"        nullable:"false"`
		EndTime          string    `example:"MM-DD-YYYY HH:MM"     json:"end_time"          nullable:"true"`
		HTMLDescription  string    `json:"html_description"`
		Rsvps            bool      `json:"rsvps"                   nullable:"false"`
		Capacity         int32     `json:"capacity"                nullable:"false"`
		TimezoneLocation string    `example:"America/Los_Angeles"  json:"timezone_location" nullable:"false"`
		Visible          bool      `json:"visible"                 nullable:"false"`
	}

	response := usecase.NewInteractor(
		func(ctx context.Context, input eventInput, output *eventOutput) error {
			// userData := a.Auth.ValidateJWT(input.AccessToken)

			// if userData == nil {
			// 	return status.Wrap(errors.New("you must be logged in to perform this action"), status.Unauthenticated)
			// }

			// if userData.Role != adminRole || userData.Role != businessRole {
			// 	return status.Wrap(errors.New("you do not have permission to perform this action"), status.PermissionDenied)
			// }

			parsedStart, err := time.Parse(timeFormat, input.StartTime)
			if err != nil {
				log.Println(fmt.Errorf("could not parse start time: %w", err))
				return status.Wrap(errors.New("could not parse start time"), status.InvalidArgument)
			}

			parsedEnd, err := time.Parse(timeFormat, input.EndTime)
			if err != nil {
				log.Println(fmt.Errorf("could not parse start time: %w", err))
				return status.Wrap(errors.New("could not parse start time"), status.InvalidArgument)
			}

			newUUID, err := uuid.NewRandom()
			if err != nil {
				log.Println(fmt.Errorf("could not generate new uuid: %w", err))
				return status.Wrap(errors.New(internalErrMsg), status.Internal)
			}

			city, err := a.getLocationID(input.CityName, input.StateCode)
			if err != nil {
				return err
			}

			conn, err := a.dbConn(ctx)
			if err != nil {
				return err
			}
			defer conn.Release()

			queries := db.New(conn)

			inputArgs := db.CreateEventParams{
				ID:               newUUID,
				Name:             input.Name,
				Host:             input.Host,
				LocationName:     pgtype.Text{String: input.LocationName, Valid: len(input.LocationName) != 0},
				StreetAddress:    input.StreetAddress,
				City:             city,
				Zipcode:          input.Zipcode,
				StartTime:        pgtype.Timestamp{Time: parsedStart, Valid: true},
				EndTime:          pgtype.Timestamp{Time: parsedEnd, Valid: true},
				HtmlDescription:  pgtype.Text{String: input.HTMLDescription, Valid: len(input.HTMLDescription) != 0},
				Rsvps:            input.Rsvps,
				Capacity:         pgtype.Int4{Int32: input.Capacity, Valid: true},
				TimezoneLocation: input.TimezoneLocation,
				Visible:          input.Visible,
			}

			newEvent, err := queries.CreateEvent(ctx, inputArgs)
			if err != nil {
				if strings.Contains(err.Error(), "host_fkey") {
					return status.Wrap(errors.New("invalid host id"), status.InvalidArgument)
				}
				if strings.Contains(err.Error(), "timezone_location_fkey") {
					return status.Wrap(errors.New("invalid IANA timezone location"), status.InvalidArgument)
				}
				log.Println(fmt.Errorf("failed to create event: %w", err))
				return status.Wrap(errors.New(internalErrMsg), status.Internal)
			}

			output.ID = newEvent.ID
			output.Name = newEvent.Name
			output.Host = newEvent.Host
			output.Location.Name = newEvent.LocationName.String
			output.Location.StreetAddress = newEvent.StreetAddress
			output.Location.CityID = newEvent.City
			output.Location.City = input.CityName
			output.Location.StateCode = input.StateCode
			output.Location.Zipcode = newEvent.Zipcode
			output.HTMLDescription = newEvent.HtmlDescription.String
			output.Rsvps = newEvent.Rsvps
			output.Capacity = newEvent.Capacity.Int32
			output.TimezoneLocation = newEvent.TimezoneLocation
			output.Visible = newEvent.Visible

			output.StartTime, err = getTimeWithTimezone(newEvent.TimezoneLocation, newEvent.StartTime.Time)
			if err != nil {
				return err
			}

			output.EndTime, err = getTimeWithTimezone(newEvent.TimezoneLocation, newEvent.EndTime.Time)
			if err != nil {
				return err
			}

			return nil
		})
	return response
}

// getTimeWithTimezone is a helper function that takes in an IANA timezone
// location and a time.Time timestamp, then returns a formatted timestamp with
// the correct time zone attached.
// Ex. "02/26/2020 16:04 PST" from "America/Los_Angeles"
func getTimeWithTimezone(timezoneLocation string, timestamp time.Time) (string, error) {
	// load the applicable IANA timezone
	loadedTimezone, err := time.LoadLocation(timezoneLocation)
	if err != nil {
		log.Println(fmt.Errorf("could not load IANA timezone: %w", err))
		return "", status.Wrap(errors.New(internalErrMsg), status.Internal)
	}

	// apply the timezone to grab the shorthand version
	timezone, _ := timestamp.In(loadedTimezone).Zone()

	return timestamp.Format(timeFormat) + " " + timezone, nil
}
