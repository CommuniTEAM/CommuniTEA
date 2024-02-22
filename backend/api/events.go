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

type eventInput struct {
	defaultInput
	ID uuid.UUID `nullable:"false" path:"id" required:"true"`
}

// type eventOutput struct {
// 	ID              uuid.UUID `json:"id"               required:"true"`
// 	Name            string    `json:"name"             required:"true"`
// 	Host            uuid.UUID `json:"host"             required:"true"`
// 	LocationName    string    `json:"location_name"    required:"true"`
// 	StreetAddress   string    `json:"street_address"   required:"true"`
// 	City            uuid.UUID `json:"city"             required:"true"`
// 	State           string    `json:"state"            required:"true"`
// 	Zipcode         string    `json:"zipcode"          required:"true"`
// 	Date            int64     `json:"date"             required:"true"`
// 	StartTime       int64     `json:"start_time"       nullable:"false" required:"true"`
// 	EndTime         int64     `json:"end_time"         nullable:"true"  required:"true"`
// 	MdDescription   string    `json:"md_description"`
// 	HTMLDescription string    `json:"html_description"`
// 	Rsvps           bool      `json:"rsvps"            required:"true"`
// 	Capacity        int       `json:"capacity"         required:"true"`
// }

func (a *API) GetEvent() usecase.Interactor {
	response := usecase.NewInteractor(
		func(ctx context.Context, input eventInput, output *db.Event) error {
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

			return nil
		})

	response.SetTitle("Get Event")
	response.SetDescription("Get the details of an event.")
	response.SetTags("Events")
	response.SetExpectedErrors(status.InvalidArgument)

	return response
}
