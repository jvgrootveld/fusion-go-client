package index

import (
	"context"
	"encoding/json"
	"github.com/jvgrootveld/fusion-go-client/fusion/index"
	"github.com/jvgrootveld/fusion-go-client/fusion/index/stage"
	"testing"

	"github.com/jvgrootveld/fusion-go-client/test/testsuit"
	"github.com/stretchr/testify/assert"
)

func TestPipelineGetter(t *testing.T) {
	t.Run("Index Pipeline - Getter - one", func(t *testing.T) {
		expectStatusCode := 200
		pipelineId := "pipeline-id"

		body := createDataIndexPipeline("1")

		jsonBody, err := json.Marshal(body)
		assert.NoError(t, err)

		client := testsuit.CreateStatusCodeUrlValidatorWithBodyHttpClient(t, expectStatusCode, buildPipelineUrl(pipelineId), jsonBody)

		response, err := testsuit.CreateFusionTestClient(client).IndexPipeline().Getter().
			WithID(pipelineId).
			Do(context.Background())

		assert.NoError(t, err)
		assert.Equal(t, 1, len(response))
		assert.Equal(t, body.Id, response[0].Id)
		assert.Equal(t, len(body.Stages), len(response[0].Stages))
	})

	t.Run("Index Pipeline - Getter - multiple", func(t *testing.T) {
		expectStatusCode := 200

		body := []index.Pipeline{
			createDataIndexPipeline("1"),
			createDataIndexPipeline("2"),
		}

		jsonBody, err := json.Marshal(body)
		assert.NoError(t, err)

		client := testsuit.CreateStatusCodeUrlValidatorWithBodyHttpClient(t, expectStatusCode, buildPipelineUrl(""), jsonBody)

		response, err := testsuit.CreateFusionTestClient(client).IndexPipeline().Getter().
			Do(context.Background())

		assert.NoError(t, err)
		assert.Equal(t, 2, len(response))
		assert.Equal(t, body[0].Id, response[0].Id)
		assert.Equal(t, len(body[0].Stages), len(response[0].Stages))
	})
}

func createDataIndexPipeline(id string) index.Pipeline {
	return index.Pipeline{
		Id: id,
		Stages: []stage.Stage{
			stage.NewBaseStage("test-stage", "stage-1-"+id),
			stage.NewBaseStage("test-stage", "stage-2-"+id),
		},
	}
}
