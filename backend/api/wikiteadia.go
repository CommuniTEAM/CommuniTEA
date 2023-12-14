package api

import (
	"context"
	"fmt"

	db "github.com/CommuniTEAM/CommuniTEA/db/sqlc"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/swaggest/usecase"
	"github.com/swaggest/usecase/status"
)

type getTeasInput struct {
	Published bool `path:"published"`
}

type teaInput struct {
	ID pgtype.UUID

	Name string `json:"name" minLength:"1"`

	ImgURL string `default:"" json:"img_url ,omitempty"`

	Description string `json:"description" minLength:"1"`

	BrewTime string `default:"" json:"brew_time ,omitempty"`

	BrewTemp float64 `json:"brew_temp ,omitempty"`

	Published bool
}

func CreateTea(dbPool *pgxpool.Pool) usecase.Interactor {
	response := usecase.NewInteractor(func(ctx context.Context, input teaInput, output *db.Tea) error {
		// dbCtx := context.Background()

		conn, err := dbPool.Acquire(ctx)

		// conn, err := pgx.Connect(dbCtx, (os.Getenv("DB_URI")))

		if err != nil {
			return fmt.Errorf("failed to connect to DB: %w", err)
		}

		defer conn.Release()

		// defer conn.Close(dbCtx)

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

			ID: pgtype.UUID{Bytes: newUUID, Valid: true},

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

func GetAllTeas(dbPool *pgxpool.Pool) usecase.Interactor {
	response := usecase.NewInteractor(func(ctx context.Context, input getTeasInput, output *[]db.Tea) error {
		// dbCtx := context.Background()

		// conn, err := pgx.Connect(dbCtx, (os.Getenv("DB_URI")))

		conn, err := dbPool.Acquire(ctx)

		if err != nil {
			return fmt.Errorf("failed to connect to DB: %w", err)
		}

		// defer conn.Close(dbCtx)

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
