package datamodel

import (
	"github.com/jvgrootveld/fusion-go-client/fusion/connection"
	"github.com/jvgrootveld/fusion-go-client/fusion/generic"
)

const ApiName = "data-models"

// API Contains all the builders required to access the Fusion DataModel
// Also see Fusion docs: https://doc.lucidworks.com/fusion/5.8/8824/data-models-api
type API struct {
	connection  *connection.Connection
	application string
}

// NewDataModelApi api group with connection
func NewDataModelApi(con *connection.Connection, application string) *API {
	return &API{connection: con, application: application}
}

// Deleter new builder to delete a DataModel
func (api *API) Deleter() *generic.Deleter {
	return generic.NewDeleter(api.connection, ApiName).
		ForApplication(api.application)
}

// Getter new builder to retrieve one or all DataModels
func (api *API) Getter() *generic.Getter[DataModel] {
	return generic.NewGetter[DataModel](api.connection, ApiName).
		ForApplication(api.application)
}

// Creator builder to create a new DataModels
func (api *API) Creator() *generic.Creator[DataModel] {
	return generic.NewCreator[DataModel](api.connection, ApiName).
		ForApplication(api.application).
		WithModelValidator(createValidator)
}
