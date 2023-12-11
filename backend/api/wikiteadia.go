package api

import (
	"context"
	"fmt"
	"os"

	db "github.com/CommuniTEAM/CommuniTEA/db/sqlc"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/swaggest/usecase"
	"github.com/swaggest/usecase/status"
)

type getTeasInput struct {
	Published bool `path:"published"`
}

type teaInput struct {
	ID          pgtype.UUID   `json:"id"`
	Name        string        `json:"name"`
	ImgURL      pgtype.Text   `json:"img_url"`
	Description string        `json:"description"`
	BrewTime    pgtype.Text   `json:"brew_time"`
	BrewTemp    pgtype.Float8 `json:"brew_temp"`
	Published   bool          `json:"published"`
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

		teaParams := db.CreateTeaParams{
			ID:          input.ID,
			Name:        input.Name,
			ImgUrl:      input.ImgURL,
			Description: input.Description,
			BrewTime:    input.BrewTime,
			BrewTemp:    input.BrewTemp,
			Published:   input.Published,
		}
		// db.CreateTeaParams.Name = fmt.Sprintf("%s", input.Name)

		// db.CreateTeaParams.ImgUrl = input.ImgUrl

		// db.CreateTeaParams.Description = fmt.Sprintf("%s", input.Description)

		// db.CreateTeaParams.BrewTime = input.BrewTime

		// db.CreateTeaParams.BrewTemp = input.BrewTemp

		// db.CreateTeaParams.Published = false

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
