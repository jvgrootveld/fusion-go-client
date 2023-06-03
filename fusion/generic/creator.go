package generic

import (
	"context"
	"net/http"

	"github.com/jvgrootveld/fusion-go-client/fusion/connection"
	"github.com/jvgrootveld/fusion-go-client/fusion/except"
	"github.com/jvgrootveld/fusion-go-client/fusion/pathbuilder"
)

// Creator builder to create a configured objects
type Creator[T interface{}] struct {
	connection     *connection.Connection
	apiName        string
	application    string
	model          T
	modelValidator func(model T) error
}

// NewCreator with Connection and apiName of the configured object to create like index-pipeline or query-profile
func NewCreator[T interface{}](connection *connection.Connection, apiName string) *Creator[T] {
	return &Creator[T]{
		connection: connection,
		apiName:    apiName,
	}
}

// ForApplication specifies the application the configured object should be created.
// When set it's interpreted as an application api e.g. `/api/apps/acme/index-pipelines`
// When empty it's interpreted as a generic api e.g. `/api/collections'
// Note: This overrides the default in the Api
func (creator *Creator[T]) ForApplication(application string) *Creator[T] {
	creator.application = application
	return creator
}

// WithModelValidator validated the model before sending it to Fusion
func (creator *Creator[T]) WithModelValidator(modelValidator func(model T) error) *Creator[T] {
	creator.modelValidator = modelValidator
	return creator
}

// WithModel to create
func (creator *Creator[T]) WithModel(model T) *Creator[T] {
	creator.model = model
	return creator
}

// Do create the specified configured object in Fusion
func (creator *Creator[T]) Do(ctx context.Context) error {
	err := creator.modelValidator(creator.model)
	if err != nil {
		return err
	}

	body := creator.model

	path := pathbuilder.ApiPath(pathbuilder.Components{
		Application: creator.application,
		ApiName:     creator.apiName,
	})
	responseData, err := creator.connection.RunREST(ctx, path, http.MethodPost, body)
	return except.CheckResponseDataErrorAndStatusCode(responseData, err, 200, 201)
}
