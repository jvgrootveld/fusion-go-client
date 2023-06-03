package generic

import (
	"context"
	"github.com/jvgrootveld/fusion-go-client/fusion/connection"
	"github.com/jvgrootveld/fusion-go-client/fusion/except"
	"github.com/jvgrootveld/fusion-go-client/fusion/pathbuilder"
	"net/http"
)

// Getter builder to get one or all configured objects
type Getter[T interface{}] struct {
	connection  *connection.Connection
	apiName     string
	application string
	id          string
}

// NewGetter with connection and apiName of the configured object to get like index-pipeline or query-profile
func NewGetter[T interface{}](connection *connection.Connection, apiName string) *Getter[T] {
	return &Getter[T]{
		connection: connection,
		apiName:    apiName,
	}
}

// ForApplication specifies the application the configured object is.
// When set it's interpreted as an application api e.g. `/api/apps/acme/index-pipelines`
// When empty it's interpreted as a generic api e.g. `/api/collections'
func (getter *Getter[T]) ForApplication(application string) *Getter[T] {
	getter.application = application
	return getter
}

// WithID specifies a specific configured object to get.
// When empty, all are returned.
func (getter *Getter[T]) WithID(id string) *Getter[T] {
	getter.id = id
	return getter
}

// Do get one or all configured objects
func (getter *Getter[T]) Do(ctx context.Context) ([]T, error) {
	path := pathbuilder.ApiPath(pathbuilder.Components{
		ApiName:     getter.apiName,
		Application: getter.application,
		ObjectId:    getter.id,
	})

	responseData, err := getter.connection.RunREST(ctx, path, http.MethodGet, nil)
	err = except.CheckResponseDataErrorAndStatusCode(responseData, err, 200)
	if err != nil {
		return nil, err
	}

	// If id is set, unmarshal to a single instance
	if getter.id != "" {
		var resultCollection T
		parseErr := responseData.DecodeBodyIntoTarget(&resultCollection)
		return []T{resultCollection}, parseErr
	}

	var resultCollections []T
	parseErr := responseData.DecodeBodyIntoTarget(&resultCollections)
	return resultCollections, parseErr
}
