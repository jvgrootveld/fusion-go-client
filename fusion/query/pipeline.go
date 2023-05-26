package query

import (
	"github.com/jvgrootveld/fusion-go-client/fusion/connection"
	"github.com/jvgrootveld/fusion-go-client/fusion/generic"
)

const PipelineApiName = "query-pipelines"

// PipelineAPI Contains all the builders required to access the Fusion Query Pipeline PipelineAPI
type PipelineAPI struct {
	connection  *connection.Connection
	application string
}

// NewPipeline api group with connection
func NewPipeline(con *connection.Connection, application string) *PipelineAPI {
	return &PipelineAPI{connection: con, application: application}
}

// Deleter new builder to delete Query Pipeline's
func (api *PipelineAPI) Deleter() *generic.Deleter {
	return generic.NewDeleter(api.connection, PipelineApiName).
		ForApplication(api.application)
}

// Creator builder to create new Query Pipeline's
func (api *PipelineAPI) Creator() *PipelineCreator {
	return NewPipelineCreator(api.connection).
		ForApplication(api.application)
}
