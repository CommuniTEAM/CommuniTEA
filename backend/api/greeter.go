// This is a test function for the Swagger UI

package api

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/swaggest/usecase"
	"github.com/swaggest/usecase/status"
)

// Declare input port type.
type helloInput struct {
	Locale string `default:"en-US" enum:"ru-RU,en-US" pattern:"^[a-z]{2}-[A-Z]{2}$" query:"locale"`
	Name   string `minLength:"3"   path:"name"` // Field tags define parameter location and JSON schema constraints.

	// Field tags of unnamed fields are applied to parent schema.
	// they are optional and can be used to disallow unknown parameters.
	// For non-body params, name tag must be provided explicitly.
	// E.g. here no unknown `query` and `cookie` parameters allowed,
	// unknown `header` params are ok.
	_ struct{} `query:"_" cookie:"_" additionalProperties:"false"`
}

// Declare output port type.
type helloOutput struct {
	Now     time.Time `header:"X-Now" json:"-"`
	Message string    `json:"message"`
}

func Greet() usecase.IOInteractorOf[helloInput, helloOutput] {
	messages := map[string]string{
		"en-US": "Hello, %s!",
		"ru-RU": "Привет, %s!",
	}

	// Create use case interactor with references to input/output types and interaction function.
	u := usecase.NewInteractor(func(ctx context.Context, input helloInput, output *helloOutput) error {
		msg, available := messages[input.Locale]
		if !available {
			return status.Wrap(errors.New("unknown locale"), status.InvalidArgument)
		}

		output.Message = fmt.Sprintf(msg, input.Name)
		output.Now = time.Now()

		return nil
	})

	// Describe use case interactor.
	u.SetTitle("Greeter")
	u.SetDescription("Greeter greets you.")

	u.SetExpectedErrors(status.InvalidArgument)

	return u
}
