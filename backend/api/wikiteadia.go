package api

import (
	"context"
	"fmt"
	"os"

	db "github.com/CommuniTEAM/CommuniTEA/db/sqlc"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/swaggest/usecase"
	"github.com/swaggest/usecase/status"
)

type getTeasInput struct {
	Published bool `path:"published"`
}

type teaInput struct {
	ID          pgtype.UUID
	Name        string      `json:"name"                 minLength:"1"`
	ImgURL      pgtype.Text `json:"img_url ,omitempty"`
	Description string      `json:"description"          minLength:"1"`
	BrewTime    pgtype.Text `json:"brew_time ,omitempty"`
	BrewTemp    float64     `json:"brew_temp ,omitempty"`
	Published   bool
}

func CreateTea() usecase.Interactor {
	response := usecase.NewInteractor(func(ctx context.Context, input teaInput, output *db.Tea) error {
		dbCtx := context.Background()

		conn, err := pgx.Connect(dbCtx, (os.Getenv("DB_URI")))

		if err != nil {
			return fmt.Errorf("failed to connect to DB: %w", err)
		}

		defer conn.Close(dbCtx)

		queries := db.New(conn)

		newUUID, err := uuid.NewRandom()

		if err != nil {
			return fmt.Errorf("failed to create UUID: %w", err)
		}

		isValid := true

		if input.BrewTemp == 0 {
			isValid = false
		}

		teaParams := db.CreateTeaParams{
			ID:          pgtype.UUID{Bytes: newUUID, Valid: true},
			Name:        input.Name,
			ImgUrl:      input.ImgURL,
			Description: input.Description,
			BrewTime:    input.BrewTime,
			BrewTemp:    pgtype.Float8{Float64: input.BrewTemp, Valid: isValid},
			Published:   false,
		}

		*output, err = queries.CreateTea(ctx, teaParams)

		if err != nil {
			return fmt.Errorf("failed to create tea: %w", err)
		}

		return nil
	})
	response.SetTitle("Create Tea")
	response.SetDescription("Make a new tea")

	response.SetExpectedErrors(status.InvalidArgument)

	return response
}

func GetAllTeas() usecase.Interactor {
	response := usecase.NewInteractor(func(ctx context.Context, input getTeasInput, output *[]db.Tea) error {
		dbCtx := context.Background()

		conn, err := pgx.Connect(dbCtx, (os.Getenv("DB_URI")))

		if err != nil {
			return fmt.Errorf("failed to connect to DB: %w", err)
		}

		defer conn.Close(dbCtx)

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
