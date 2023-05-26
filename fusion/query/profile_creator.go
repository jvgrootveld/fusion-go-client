package query

import (
	"context"
	"fmt"
	"net/http"

	"github.com/jvgrootveld/fusion-go-client/fusion/connection"
	"github.com/jvgrootveld/fusion-go-client/fusion/except"
	"github.com/jvgrootveld/fusion-go-client/fusion/pathbuilder"
)

// ProfileCreator builder to create new Query Profile's
type ProfileCreator struct {
	connection           *connection.Connection
	application          string
	id                   string
	queryPipeline        string
	collection           string
	searchHandler        string
	params               []Param
	additionalProperties map[string]string
}

// NewProfileCreator with Connection
func NewProfileCreator(connection *connection.Connection) *ProfileCreator {
	return &ProfileCreator{
		connection: connection,
	}
}

// ForApplication specifies the application the pipeline should be created.
// Note: This overrides the default in the Connection
func (creator *ProfileCreator) ForApplication(application string) *ProfileCreator {
	creator.application = application
	return creator
}

func (creator *ProfileCreator) WithID(id string) *ProfileCreator {
	creator.id = id
	return creator
}

func (creator *ProfileCreator) WithQueryPipeline(queryPipeline string) *ProfileCreator {
	creator.queryPipeline = queryPipeline
	return creator
}

func (creator *ProfileCreator) WithCollection(collection string) *ProfileCreator {
	creator.collection = collection
	return creator
}

// WithParams is optional
func (creator *ProfileCreator) WithParams(params []Param) *ProfileCreator {
	creator.params = params
	return creator
}

// WithParam adds a single param on the new Query Profile. Is optional
func (creator *ProfileCreator) WithParam(param Param) *ProfileCreator {
	creator.params = append(creator.params, param)
	return creator
}

// WithAdditionalProperties is optional
func (creator *ProfileCreator) WithAdditionalProperties(additionalProperties map[string]string) *ProfileCreator {
	creator.additionalProperties = additionalProperties
	return creator
}

// WithAdditionalProperty adds a single property on the new Query Profile. Is optional
func (creator *ProfileCreator) WithAdditionalProperty(key, value string) *ProfileCreator {
	if creator.additionalProperties == nil {
		creator.additionalProperties = make(map[string]string, 1)
	}

	creator.additionalProperties[key] = value
	return creator
}

// Do create the specified Query Profile in Fusion
func (creator *ProfileCreator) Do(ctx context.Context) error {
	body, err := creator.CreateRequestObject()
	if err != nil {
		return err
	}

	path := pathbuilder.AppsPath(pathbuilder.Components{
		Application: creator.application,
		ApiName:     ProfileApiName,
	})
	responseData, err := creator.connection.RunREST(ctx, path, http.MethodPost, body)
	return except.CheckResponseDataErrorAndStatusCode(responseData, err, 200, 201)
}

type CreatorRequestData struct {
	Id                   string            `json:"id"`
	QueryPipeline        string            `json:"queryPipeline"`
	Collection           string            `json:"collection"`
	SearchHandler        string            `json:"searchHandler,omitempty"`
	Params               []Param           `json:"params,omitempty"`
	AdditionalProperties map[string]string `json:"additionalProperties,omitempty"`
}

func (creator *ProfileCreator) CreateRequestObject() (*CreatorRequestData, error) {
	err := creator.checkRequired()
	if err != nil {
		return nil, err
	}

	return &CreatorRequestData{
		Id:                   creator.id,
		QueryPipeline:        creator.queryPipeline,
		Collection:           creator.collection,
		SearchHandler:        creator.searchHandler,
		Params:               creator.params,
		AdditionalProperties: creator.additionalProperties,
	}, err
}

func (creator *ProfileCreator) checkRequired() error {
	typeName := fmt.Sprint(ProfileApiName, "ProfileCreator")

	if creator.id == "" {
		return except.NewRequiredError(typeName, "id")
	}
	if creator.queryPipeline == "" {
		return except.NewRequiredError(typeName, "queryPipeline")
	}
	if creator.collection == "" {
		return except.NewRequiredError(typeName, "collection")
	}

	return nil
}
