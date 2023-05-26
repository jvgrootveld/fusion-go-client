package index

import (
	"context"
	"fmt"
	"net/http"

	"github.com/jvgrootveld/fusion-go-client/fusion/connection"
	"github.com/jvgrootveld/fusion-go-client/fusion/except"
	"github.com/jvgrootveld/fusion-go-client/fusion/pathbuilder"
)

// ProfileCreator builder to create new Index Profile's
type ProfileCreator struct {
	connection           *connection.Connection
	application          string
	id                   string
	indexPipeline        string
	collection           string
	parser               string
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

func (creator *ProfileCreator) WithIndexPipeline(indexPipeline string) *ProfileCreator {
	creator.indexPipeline = indexPipeline
	return creator
}

func (creator *ProfileCreator) WithCollection(collection string) *ProfileCreator {
	creator.collection = collection
	return creator
}

// WithParser is optional
func (creator *ProfileCreator) WithParser(parser string) *ProfileCreator {
	creator.parser = parser
	return creator
}

// WithAdditionalProperties is optional
func (creator *ProfileCreator) WithAdditionalProperties(additionalProperties map[string]string) *ProfileCreator {
	creator.additionalProperties = additionalProperties
	return creator
}

// WithAdditionalProperty adds a single property on the new Index Profile. Is optional
func (creator *ProfileCreator) WithAdditionalProperty(key, value string) *ProfileCreator {
	if creator.additionalProperties == nil {
		creator.additionalProperties = make(map[string]string, 1)
	}

	creator.additionalProperties[key] = value
	return creator
}

// Do create the specified Index Profile in Fusion
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

type ProfileCreatorRequestData struct {
	Id                   string            `json:"id"`
	IndexPipeline        string            `json:"indexPipeline"`
	Collection           string            `json:"collection"`
	Parser               string            `json:"parser,omitempty"`
	AdditionalProperties map[string]string `json:"additionalProperties,omitempty"`
}

func (creator *ProfileCreator) CreateRequestObject() (*ProfileCreatorRequestData, error) {
	err := creator.checkRequired()
	if err != nil {
		return nil, err
	}

	return &ProfileCreatorRequestData{
		Id:                   creator.id,
		IndexPipeline:        creator.indexPipeline,
		Collection:           creator.collection,
		Parser:               creator.parser,
		AdditionalProperties: creator.additionalProperties,
	}, err
}

func (creator *ProfileCreator) checkRequired() error {
	typeName := fmt.Sprint(ProfileApiName, "Creator")

	if creator.id == "" {
		return except.NewRequiredError(typeName, "id")
	}
	if creator.indexPipeline == "" {
		return except.NewRequiredError(typeName, "indexPipeline")
	}
	if creator.collection == "" {
		return except.NewRequiredError(typeName, "collection")
	}

	return nil
}
