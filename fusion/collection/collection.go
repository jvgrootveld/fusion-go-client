package collection

import (
	"github.com/jvgrootveld/fusion-go-client/fusion/connection"
	"github.com/jvgrootveld/fusion-go-client/fusion/generic"
)

const ApiName = "collections"

// API Contains all the builders required to access the Fusion Collection
// Also see Fusion docs: https://doc.lucidworks.com/fusion/5.8/307/collections-api
type API struct {
	connection  *connection.Connection
	application string
}

// NewCollectionApi api group with connection
func NewCollectionApi(con *connection.Connection, application string) *API {
	return &API{connection: con, application: application}
}

// Deleter new builder to delete a Collection
func (api *API) Deleter() *generic.Deleter {
	return generic.NewDeleter(api.connection, ApiName).
		ForApplication(api.application)
}

// Getter new builder to retrieve one or all Collections
func (api *API) Getter() *generic.Getter[Collection] {
	return generic.NewGetter[Collection](api.connection, ApiName).
		ForApplication(api.application)
}

// Creator builder to create a new Collections
func (api *API) Creator() *generic.Creator[CreateCollection] {
	return generic.NewCreator[CreateCollection](api.connection, ApiName).
		ForApplication(api.application).
		WithModelValidator(createValidator)
}

// Clearer builder to clear a collection
func (api *API) Clearer() *Clearer {
	return NewClearer(api.connection)
}
