package api

import (
	"context"
	"errors"
	"fmt"
	"log"

	db "github.com/CommuniTEAM/CommuniTEA/db/sqlc"
	"github.com/google/uuid"
	"github.com/swaggest/usecase"
	"github.com/swaggest/usecase/status"
)

type businessInput struct {
	Name            string    `json:"name"              nullable:"false" required:"true"`
	StreetAddress   string    `json:"street_address"    nullable:"false" required:"true"`
	City            uuid.UUID `json:"city"              nullable:"false" required:"true"`
	State           string    `json:"state"             nullable:"false" required:"true"`
	Zipcode         string    `json:"zipcode"           nullable:"false" required:"true"`
	BusinessOwnerID uuid.UUID `json:"business_owner_id" nullable:"false" required:"true"`
}

type businessOutput struct {
	Businesses []db.Business `json:"businesses"`
}

// GetAllBusinesses returns all businesses in the database
func (a *API) GetAllBusinesses() usecase.Interactor {
	response := usecase.NewInteractor(
		func(ctx context.Context, _defaultInput, output *businessOutput) error {
			conn, err := a.dbConn(ctx)
			if err != nil {
				return err
			}
			defer conn.Release()

			queries := db.New(conn)

			output.Businesses, err = queries.GetAllBusinesses(ctx)
			if err != nil {
				log.Println(fmt.Errorf("error getting all businesses: %w", err))
				return status.Wrap(errors.New(internalErrMsg), status.Internal)
			}

			return nil
		})

	response.SetTitle("Get All Businesses")
	response.SetDescription("Get all businesses in the database")
	response.SetTags("Businesses")

	return response
}

// CreateBusinesses creates a new business in the database
func (a *API) CreateBusinesses() usecase.Interactor {
	response := usecase.NewInteractor(func(ctx context.Context, input businessInput, output *db.Business) error {
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

		businessParams := db.CreateBusinessParams{
			ID:              newUUID,
			Name:            input.Name,
			StreetAddress:   input.StreetAddress,
			City:            input.City,
			State:           input.State,
			Zipcode:         input.Zipcode,
			BusinessOwnerID: input.BusinessOwnerID,
		}

		*output, err = queries.CreateBusiness(ctx, businessParams)
		if err != nil {
			return fmt.Errorf("failed to create business: %w", err)
		}

		return nil
	})

	response.SetTitle("Create Business")
	response.SetDescription("Make a new business")
	response.SetTags("Businesses")
	response.SetExpectedErrors(status.InvalidArgument)

	return response
}