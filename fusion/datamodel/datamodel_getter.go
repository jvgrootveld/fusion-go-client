package datamodel

import (
	"context"
	"net/http"

	"github.com/jvgrootveld/fusion-go-client/fusion/connection"
	"github.com/jvgrootveld/fusion-go-client/fusion/except"
	"github.com/jvgrootveld/fusion-go-client/fusion/pathbuilder"
)

// Getter builder to get a Data Model
type Getter struct {
	connection  *connection.Connection
	application string
	id          string
}

// NewDataModelGetter with connection
func NewDataModelGetter(connection *connection.Connection) *Getter {
	return &Getter{
		connection: connection,
	}
}

// ForApplication specifies the application the Data Model is configured.
// Note: This overrides the default in the Connection
func (getter *Getter) ForApplication(application string) *Getter {
	getter.application = application
	return getter
}

// WithID specifies a specific DataModel to get.
// When empty, all are returned.
func (getter *Getter) WithID(id string) *Getter {
	getter.id = id
	return getter
}

// Do get one or all Data Model's
func (getter *Getter) Do(ctx context.Context) ([]DataModel, error) {
	path := pathbuilder.AppsPath(pathbuilder.Components{
		Application: getter.application,
		ApiName:     ApiName,
		ObjectId:    getter.id,
	})
	responseData, err := getter.connection.RunREST(ctx, path, http.MethodGet, nil)
	err = except.CheckResponseDataErrorAndStatusCode(responseData, err, 200)
	if err != nil {
		return nil, err
	}

	// If id is set, unmarshal to a single instance
	if getter.id != "" {
		var resultDataModel DataModel
		parseErr := responseData.DecodeBodyIntoTarget(&resultDataModel)
		return []DataModel{resultDataModel}, parseErr
	}

	var resultDataModels []DataModel
	parseErr := responseData.DecodeBodyIntoTarget(&resultDataModels)
	return resultDataModels, parseErr
}
