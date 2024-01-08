package api

import (
	"context"
	"fmt"

	"github.com/CommuniTEAM/CommuniTEA/auth"
	db "github.com/CommuniTEAM/CommuniTEA/db/sqlc"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/swaggest/usecase"
	"github.com/swaggest/usecase/status"
)

type getTeasInput struct {
	Published bool `path:"published"`
}

type teaInput struct {
	defaultInput

	ID uuid.UUID

	Name string `json:"name" minLength:"1"`

	ImgURL string `default:"" json:"img_url ,omitempty"`

	Description string `json:"description" minLength:"1"`

	BrewTime string `default:"" json:"brew_time ,omitempty"`

	BrewTemp float64 `json:"brew_temp ,omitempty"`

	Published bool
}

func (a *API) CreateTea() usecase.Interactor {
	response := usecase.NewInteractor(func(ctx context.Context, input teaInput, output *db.Tea) error {
		userData := auth.ValidateJWT(input.AccessToken)

		// If the token was invalid or nonexistent then userData will be nil

		if userData == nil {
			return status.Wrap(fmt.Errorf("you must be logged in to perform this action"), status.Unauthenticated)
		}

		conn, err := a.dbConn(ctx)

		if err != nil {
			return err
		}

		defer conn.Release()

		queries := db.New(conn)

		newUUID, err := uuid.NewRandom()

		if err != nil {
			return fmt.Errorf("failed to create UUID: %w", err)
		}

		isImgURLValid := true

		if input.ImgURL == "" {
			isImgURLValid = false
		}

		isBrewTimeValid := true

		if input.BrewTime == "" {
			isBrewTimeValid = false
		}

		isBrewTempValid := true

		if input.BrewTemp == 0 {
			isBrewTempValid = false
		}

		teaParams := db.CreateTeaParams{

			ID: newUUID,

			Name: input.Name,

			ImgUrl: pgtype.Text{String: input.ImgURL, Valid: isImgURLValid},

			Description: input.Description,

			BrewTime: pgtype.Text{String: input.BrewTime, Valid: isBrewTimeValid},

			BrewTemp: pgtype.Float8{Float64: input.BrewTemp, Valid: isBrewTempValid},

			Published: false,
		}

		*output, err = queries.CreateTea(ctx, teaParams)

		if err != nil {
			return fmt.Errorf("failed to create tea: %w", err)
		}

		return nil
	})

	response.SetTitle("Create Tea")

	response.SetDescription("Make a new tea")

	response.SetTags("Teas")

	response.SetExpectedErrors(status.InvalidArgument)

	return response
}

func (a *API) GetAllTeas() usecase.Interactor {
	response := usecase.NewInteractor(func(ctx context.Context, input getTeasInput, output *[]db.Tea) error {
		conn, err := a.dbConn(ctx)

		if err != nil {
			return err
		}

		defer conn.Release()

		queries := db.New(conn)

		*output, err = queries.GetTeas(ctx, input.Published)

		if err != nil {
			return fmt.Errorf("failed to get all teas from DB: %w", err)
		}

		return nil
	})

	response.SetTags("Teas")

	return response
}
