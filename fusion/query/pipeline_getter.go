package query

import (
	"context"
	"net/http"

	"github.com/jvgrootveld/fusion-go-client/fusion/connection"
	"github.com/jvgrootveld/fusion-go-client/fusion/except"
	"github.com/jvgrootveld/fusion-go-client/fusion/pathbuilder"
)

// PipelineGetter builder to get an Query pipeline
type PipelineGetter struct {
	connection  *connection.Connection
	application string
	id          string
}

// NewPipelineGetter with connection
func NewPipelineGetter(connection *connection.Connection) *PipelineGetter {
	return &PipelineGetter{
		connection: connection,
	}
}

// ForApplication specifies the application the pipeline is configured.
// Note: This overrides the default in the Connection
func (getter *PipelineGetter) ForApplication(application string) *PipelineGetter {
	getter.application = application
	return getter
}

// WithID specifies a specific Query Pipeline to get.
// When empty, all are returned.
func (getter *PipelineGetter) WithID(id string) *PipelineGetter {
	getter.id = id
	return getter
}

// Do get one or all Query pipeline's
func (getter *PipelineGetter) Do(ctx context.Context) ([]Pipeline, error) {
	path := pathbuilder.AppsPath(pathbuilder.Components{
		Application: getter.application,
		ApiName:     PipelineApiName,
		ObjectId:    getter.id,
	})
	responseData, err := getter.connection.RunREST(ctx, path, http.MethodGet, nil)
	err = except.CheckResponseDataErrorAndStatusCode(responseData, err, 200)
	if err != nil {
		return nil, err
	}

	// If id is set, unmarshal to a single instance
	if getter.id != "" {
		var resultPipeline Pipeline
		parseErr := responseData.DecodeBodyIntoTarget(&resultPipeline)
		return []Pipeline{resultPipeline}, parseErr
	}

	var resultPipelines []Pipeline
	parseErr := responseData.DecodeBodyIntoTarget(&resultPipelines)
	return resultPipelines, parseErr
}
