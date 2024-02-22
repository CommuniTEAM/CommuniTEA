package api

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"

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

	ImgURL string `default:"" json:"img_url"`

	Description string `json:"description" minLength:"1"`

	BrewTime string `default:"" json:"brew_time"`

	BrewTemp float64 `json:"brew_temp"`

	Published bool
}

type updateTeaInput struct {
	defaultInput

	ID uuid.UUID `path:"id"`

	Name string `json:"name" minLength:"1"`

	ImgURL string `json:"img_url"`

	Description string `json:"description" minLength:"1"`

	BrewTime string `json:"brew_time"`

	BrewTemp float64 `json:"brew_temp"`

	Published bool `json:"published"`
}

type teaPublishedPatch struct {
	defaultInput
	ID        uuid.UUID `path:"id"`
	Published bool      `json:"published"`
}

func (a *API) CreateTea() usecase.Interactor {
	response := usecase.NewInteractor(func(ctx context.Context, input teaInput, output *db.Tea) error {
		userData := a.Auth.ValidateJWT(input.AccessToken)

		// If the token was invalid or nonexistent then userData will be nil
		if userData == nil {
			return status.Wrap(errors.New("you must be logged in to perform this action"), status.Unauthenticated)
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
			ID:          newUUID,
			Name:        input.Name,
			ImgUrl:      pgtype.Text{String: input.ImgURL, Valid: isImgURLValid},
			Description: input.Description,
			BrewTime:    pgtype.Text{String: input.BrewTime, Valid: isBrewTimeValid},
			BrewTemp:    pgtype.Float8{Float64: input.BrewTemp, Valid: isBrewTempValid},
			Published:   false,
		}

		*output, err = queries.CreateTea(ctx, teaParams)
		if err != nil {
			log.Println(fmt.Errorf("failed to create tea: %w", err))

			return status.Wrap(errors.New(internalErrMsg), status.Internal)
		}
		return nil
	})

	response.SetTitle("Create Tea")
	response.SetDescription("Make a new tea")
	response.SetTags("Teas")
	response.SetExpectedErrors(

		status.InvalidArgument,

		status.Unauthenticated,
	)
	return response
}

func (a *API) UpdateTea() usecase.Interactor {
	response := usecase.NewInteractor(func(ctx context.Context, input updateTeaInput, output *db.Tea) error {
		userData := a.Auth.ValidateJWT(input.AccessToken)

		// If the token was invalid or nonexistent then userData will be nil
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

		// Get original tea details
		_, errCheck := queries.GetTea(ctx, input.ID)
		if errCheck != nil {
			if strings.Contains(errCheck.Error(), "no rows") {
				return status.Wrap(errors.New("no tea with that id"), status.NotFound)
			}
			log.Println(fmt.Errorf("could not get tea: %w", errCheck))
			return status.Wrap(errors.New(internalErrMsg), status.Internal)
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

		teaParams := db.UpdateTeaParams{
			ID:          input.ID,
			Name:        input.Name,
			ImgUrl:      pgtype.Text{String: input.ImgURL, Valid: isImgURLValid},
			Description: input.Description,
			BrewTime:    pgtype.Text{String: input.BrewTime, Valid: isBrewTimeValid},
			BrewTemp:    pgtype.Float8{Float64: input.BrewTemp, Valid: isBrewTempValid},
			Published:   input.Published,
		}

		*output, err = queries.UpdateTea(ctx, teaParams)

		if err != nil {
			log.Println("could not update tea information: %w", err)
			return status.Wrap(errors.New(internalErrMsg), status.Internal)
		}
		return nil
	})

	response.SetTitle("Update Tea")
	response.SetDescription("Update a new tea")
	response.SetTags("Teas")
	response.SetExpectedErrors(
		status.InvalidArgument,
		status.Unauthenticated,
		status.PermissionDenied,
		status.AlreadyExists,
		status.NotFound,
	)

	return response
}

func (a *API) PublishTea() usecase.Interactor {
	response := usecase.NewInteractor(func(ctx context.Context, input teaPublishedPatch, output *db.Tea) error {
		userData := a.Auth.ValidateJWT(input.AccessToken)

		// If the token was invalid or nonexistent then userData will be nil
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

		// Get original tea details
		_, errCheck := queries.GetTea(ctx, input.ID)
		if errCheck != nil {
			if strings.Contains(errCheck.Error(), "no rows") {
				return status.Wrap(errors.New("no tea with that id"), status.NotFound)
			}
			log.Println(fmt.Errorf("could not get tea: %w", errCheck))
			return status.Wrap(errors.New(internalErrMsg), status.Internal)
		}

		teaParams := db.PublishTeaParams{
			ID:        input.ID,
			Published: input.Published,
		}

		*output, err = queries.PublishTea(ctx, teaParams)

		if err != nil {
			log.Println("could not publish tea information: %w", err)
			return status.Wrap(errors.New(internalErrMsg), status.Internal)
		}
		return nil
	})

	response.SetTitle("Publish Tea")
	response.SetDescription("Publish a new tea")
	response.SetTags("Teas")
	response.SetExpectedErrors(
		status.InvalidArgument,
		status.Unauthenticated,
		status.PermissionDenied,
		status.NotFound,
	)

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

func (a *API) DeleteTea() usecase.Interactor {
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

			err = queries.DeleteTea(ctx, input.ID)
			if err != nil {
				log.Println(fmt.Errorf("could not delete tea: %w", err))
				return status.Wrap(errors.New(internalErrMsg), status.Internal)
			}

			output.Message = successMsg

			return nil
		})

	response.SetTitle("Delete Tea")
	response.SetDescription("Delete Tea Data.")
	response.SetTags("Teas")
	response.SetExpectedErrors(
		status.InvalidArgument,
		status.Unauthenticated,
		status.PermissionDenied,
	)

	return response
}
