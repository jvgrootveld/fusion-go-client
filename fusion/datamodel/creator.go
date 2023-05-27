package datamodel

import (
	"context"
	"fmt"
	"net/http"

	"github.com/jvgrootveld/fusion-go-client/fusion/connection"
	"github.com/jvgrootveld/fusion-go-client/fusion/except"
	"github.com/jvgrootveld/fusion-go-client/fusion/fault"
	"github.com/jvgrootveld/fusion-go-client/fusion/pathbuilder"
)

// Creator builder to create a new DataModel's
type Creator struct {
	connection  *connection.Connection
	application string
	model       DataModel
}

// NewDataModelCreator with Connection
func NewDataModelCreator(connection *connection.Connection) *Creator {
	return &Creator{
		connection: connection,
		model:      DataModel{},
	}
}

// ForApplication specifies the application the Data Model should be created.
// Note: This overrides the default in the Connection
func (creator *Creator) ForApplication(application string) *Creator {
	creator.application = application
	return creator
}

func (creator *Creator) WithModel(model DataModel) *Creator {
	creator.model = model
	return creator
}

func (creator *Creator) WithID(id string) *Creator {
	creator.model.Id = id
	return creator
}

func (creator *Creator) WithName(name string) *Creator {
	creator.model.Name = name
	return creator
}

func (creator *Creator) WithDescription(description string) *Creator {
	creator.model.Description = description
	return creator
}

func (creator *Creator) WithIndexPipeline(indexPipeline string) *Creator {
	creator.model.IndexPipeline = indexPipeline
	return creator
}

// WithQueryPipeline is optional
func (creator *Creator) WithQueryPipeline(queryPipeline string) *Creator {
	creator.model.QueryPipeline = queryPipeline
	return creator
}

// WithFields is optional
func (creator *Creator) WithFields(fields []Field) *Creator {
	creator.model.Fields = fields
	return creator
}

// WithField adds a single Field on the new DataModel. Is optional
func (creator *Creator) WithField(field Field) *Creator {
	if creator.model.Fields == nil {
		creator.model.Fields = []Field{}
	}

	creator.model.Fields = append(creator.model.Fields, field)
	return creator
}

// Do create the specified DataModel in Fusion
func (creator *Creator) Do(ctx context.Context) error {
	err := creator.checkRequired()
	if err != nil {
		return err
	}

	body := creator.model

	path := pathbuilder.AppsPath(pathbuilder.Components{
		Application: creator.application,
		ApiName:     ApiName,
	})
	responseData, err := creator.connection.RunREST(ctx, path, http.MethodPost, body)
	return except.CheckResponseDataErrorAndStatusCode(responseData, err, 200, 201)
}

func (creator *Creator) checkRequired() error {
	typeName := fmt.Sprint(ApiName, "Creator")

	if creator.model.Id == "" {
		return fault.NewRequiredError(typeName, "object.Id")
	}
	if creator.model.Name == "" {
		return fault.NewRequiredError(typeName, "object.Name")
	}
	if creator.model.Description == "" {
		return fault.NewRequiredError(typeName, "object.Description")
	}
	if creator.model.IndexPipeline == "" {
		return fault.NewRequiredError(typeName, "object.IndexPipeline")
	}

	return nil
}
