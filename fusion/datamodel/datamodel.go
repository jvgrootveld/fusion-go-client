package datamodel

import (
	"github.com/jvgrootveld/fusion-go-client/fusion/connection"
	"github.com/jvgrootveld/fusion-go-client/fusion/generic"
)

const ApiName = "data-models"

// API Contains all the builders required to access the Fusion DataModel
type API struct {
	connection  *connection.Connection
	application string
}

// NewDataModelApi api group with connection
func NewDataModelApi(con *connection.Connection, application string) *API {
	return &API{connection: con, application: application}
}

// Deleter new builder to delete a DataModel's
func (api *API) Deleter() *generic.Deleter {
	return generic.NewDeleter(api.connection, ApiName).
		ForApplication(api.application)
}

// Getter new builder to retrieve one or all DataModel's
func (api *API) Getter() *Getter {
	return NewDataModelGetter(api.connection).
		ForApplication(api.application)
}

// Creator builder to create a new DataModel's
func (api *API) Creator() *Creator {
	return NewDataModelCreator(api.connection).
		ForApplication(api.application)
}
