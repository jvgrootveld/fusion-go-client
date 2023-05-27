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

// NewDataModel api group with connection
func NewDataModel(con *connection.Connection, application string) *API {
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

type DataModel struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	// The name of the index pipeline used for the data model.
	IndexPipeline string `json:"indexPipeline"`
	// The name of the query pipeline used for the data model.
	QueryPipeline string           `json:"queryPipeline,omitempty"`
	Fields        []DataModelField `json:"fields,omitempty"`
}

type DataModelField struct {
	Name        string                  `json:"name"`
	Description string                  `json:"description"`
	Required    bool                    `json:"required,omitempty"`
	Mappings    []DataModelFieldMapping `json:"mappings"`
}

type DataModelFieldMapping struct {
	SolrField   string `json:"solrField"`
	QueryField  bool   `json:"queryField,omitempty"`
	PhraseMatch bool   `json:"phraseMatch,omitempty"`
	// The amount of boost to give to the query.
	// If this is a query field, apply this boost to matches
	BoostValue float32 `json:"boostValue,omitempty"`
	// The amount of boost to give to the query, if it matches as an exact phrase.
	// If this is a phrase match field, apply this boost to matches
	PhraseBoost float32 `json:"phraseBoost,omitempty"`
}
