package query

import (
	"context"
	"fmt"
	"github.com/jvgrootveld/fusion-go-client/fusion/fault"
	"net/http"

	"github.com/jvgrootveld/fusion-go-client/fusion/query/stage"

	"github.com/jvgrootveld/fusion-go-client/fusion/connection"
	"github.com/jvgrootveld/fusion-go-client/fusion/except"
	"github.com/jvgrootveld/fusion-go-client/fusion/pathbuilder"
)

// PipelineCreator builder to create new Query Pipeline's
type PipelineCreator struct {
	connection  *connection.Connection
	application string
	id          string
	stages      []stage.Stage
}

// NewPipelineCreator with Connection
func NewPipelineCreator(connection *connection.Connection) *PipelineCreator {
	return &PipelineCreator{
		connection: connection,
	}
}

// ForApplication specifies the application the pipeline should be created.
// Note: This overrides the default in the Connection
func (creator *PipelineCreator) ForApplication(application string) *PipelineCreator {
	creator.application = application
	return creator
}

func (creator *PipelineCreator) WithID(id string) *PipelineCreator {
	creator.id = id
	return creator
}

func (creator *PipelineCreator) WithStages(stages ...stage.Stage) *PipelineCreator {
	creator.stages = stages
	return creator
}

// Do create the specified Query Pipeline in Fusion
func (creator *PipelineCreator) Do(ctx context.Context) error {
	body, err := creator.CreateRequestObject()
	if err != nil {
		return err
	}

	path := pathbuilder.AppsPath(pathbuilder.Components{
		Application: creator.application,
		ApiName:     PipelineApiName,
	})
	responseData, err := creator.connection.RunREST(ctx, path, http.MethodPost, body)
	return except.CheckResponseDataErrorAndStatusCode(responseData, err, 200, 201)
}

type PipelineCreatorRequestData struct {
	Id     string        `json:"id"`
	Stages []stage.Stage `json:"stages"`
}

func (creator *PipelineCreator) CreateRequestObject() (*PipelineCreatorRequestData, error) {
	err := creator.checkRequired()
	if err != nil {
		return nil, err
	}

	return &PipelineCreatorRequestData{
		Id:     creator.id,
		Stages: creator.stages,
	}, err
}

func (creator *PipelineCreator) checkRequired() error {
	typeName := fmt.Sprint(PipelineApiName, "Creator")

	if creator.id == "" {
		return fault.NewRequiredError(typeName, "id")
	}

	return nil
}
