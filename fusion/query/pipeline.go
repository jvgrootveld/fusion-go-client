package query

import (
	"github.com/jvgrootveld/fusion-go-client/fusion/connection"
	"github.com/jvgrootveld/fusion-go-client/fusion/generic"
	"github.com/jvgrootveld/fusion-go-client/fusion/query/stage"
)

const PipelineApiName = "query-pipelines"

// PipelineAPI Contains all the builders required to access the Fusion Query Pipeline
// Also see Fusion docs: https://doc.lucidworks.com/fusion/5.8/360/query-pipelines-api
type PipelineAPI struct {
	connection  *connection.Connection
	application string
}

// NewPipelineApi api group with connection
func NewPipelineApi(con *connection.Connection, application string) *PipelineAPI {
	return &PipelineAPI{connection: con, application: application}
}

// Deleter new builder to delete Query Pipelines
func (api *PipelineAPI) Deleter() *generic.Deleter {
	return generic.NewDeleter(api.connection, PipelineApiName).
		ForApplication(api.application)
}

// Getter new builder to retrieve one or all Query Pipelines
func (api *PipelineAPI) Getter() *generic.Getter[Pipeline] {
	return generic.NewGetter[Pipeline](api.connection, PipelineApiName).
		ForApplication(api.application)
}

// Creator builder to create new Query Pipelines
func (api *PipelineAPI) Creator() *generic.Creator[Pipeline] {
	return generic.NewCreator[Pipeline](api.connection, PipelineApiName).
		ForApplication(api.application).
		WithModelValidator(createPipelineValidator)
}

type Pipeline struct {
	Id     string        `json:"id"`
	Stages []stage.Stage `json:"stages"`
}
