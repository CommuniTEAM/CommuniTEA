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

// type eventOutput struct {
// 	ID              uuid.UUID `json:"id"              `
// 	Name            string    `json:"name"            `
// 	Host            uuid.UUID `json:"host"            `
// 	LocationName    string    `json:"location_name"   `
// 	StreetAddress   string    `json:"street_address"  `
// 	City            uuid.UUID `json:"city"            `
// 	Zipcode         string    `json:"zipcode"         `
// 	StartTime       int64     `json:"start_time"       nullable:"false"`
// 	EndTime         int64     `json:"end_time"         nullable:"true" `
// 	MdDescription   string    `json:"md_description"`
// 	HTMLDescription string    `json:"html_description"`
// 	Rsvps           bool      `json:"rsvps"           `
// 	Capacity        int       `json:"capacity"        `
// }

func (a *API) GetEvent() usecase.Interactor {
	response := usecase.NewInteractor(
		func(ctx context.Context, input uuidInput, output *db.Event) error {
			conn, err := a.dbConn(ctx)
			if err != nil {
				return err
			}
			defer conn.Release()

			queries := db.New(conn)

			*output, err = queries.GetEventByID(ctx, input.ID)
			if err != nil {
				if strings.Contains(err.Error(), "no rows") {
					return status.Wrap(errors.New("event not found"), status.NotFound)
				}
				log.Println(fmt.Errorf("error getting event by id: %w", err))
				return status.Wrap(errors.New(internalErrMsg), status.Internal)
			}

			timezone, _ := time.LoadLocation(output.TimezoneLocation)
			actualTimezone, _ := output.StartTime.Time.In(timezone).Zone()

			// TODO: Use this string as start time in output
			log.Println(output.StartTime.Time.Format(timeFormat), actualTimezone)

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
		Name             string    `json:"name"                   nullable:"false"`
		Host             uuid.UUID `json:"host"                   nullable:"false"`
		LocationName     string    `json:"location_name"`
		StreetAddress    string    `json:"street_address"         nullable:"false"`
		Zipcode          string    `json:"zipcode"                nullable:"false"         pattern:"^[0-9]{5}$"`
		StartTime        string    `example:"01-02-2006 15:04"    json:"start_time"        nullable:"false"`
		EndTime          string    `example:"01-02-2006 15:04"    json:"end_time"          nullable:"true"`
		MdDescription    string    `json:"md_description"`
		Rsvps            bool      `json:"rsvps"                  nullable:"false"`
		Capacity         int32     `json:"capacity"               nullable:"false"`
		TimezoneLocation string    `example:"America/Los_Angeles" json:"timezone_location" nullable:"false"`
	}

	response := usecase.NewInteractor(
		func(ctx context.Context, input eventInput, output *db.Event) error {
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

			timezone := parsedStart.Location().String()
			if timezone != parsedEnd.Location().String() {
				return status.Wrap(errors.New("start and end timezones do not match"), status.InvalidArgument)
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
				MdDescription:    pgtype.Text{String: input.MdDescription, Valid: len(input.MdDescription) != 0},
				HtmlDescription:  pgtype.Text{String: input.MdDescription, Valid: len(input.MdDescription) != 0},
				Rsvps:            input.Rsvps,
				Capacity:         pgtype.Int4{Int32: input.Capacity, Valid: true},
				TimezoneLocation: timezone,
			}

			*output, err = queries.CreateEvent(ctx, inputArgs)
			if err != nil {
				log.Println(fmt.Errorf("failed to create event: %w", err))
				return status.Wrap(errors.New(internalErrMsg), status.Internal)
			}

			return nil
		})
	return response
}
