package index

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/jvgrootveld/fusion-go-client/fusion/index"

	"github.com/jvgrootveld/fusion-go-client/test/testsuit"
	"github.com/stretchr/testify/assert"
)

func TestProfileGetter(t *testing.T) {
	t.Run("Index Profile - Getter - one", func(t *testing.T) {
		expectStatusCode := 200
		pipelineId := "pipeline-id"

		body := createDataIndexProfile("1")

		jsonBody, err := json.Marshal(body)
		assert.NoError(t, err)

		client := testsuit.CreateStatusCodeUrlValidatorWithBodyHttpClient(t, expectStatusCode, buildProfileUrl(pipelineId), jsonBody)

		response, err := testsuit.CreateFusionTestClient(client).IndexProfile().Getter().
			WithID(pipelineId).
			Do(context.Background())

		assert.NoError(t, err)
		assert.Equal(t, 1, len(response))
		assert.Equal(t, body.Id, response[0].Id)
		assert.Equal(t, body.IndexPipeline, response[0].IndexPipeline)
	})

	t.Run("Index Profile - Getter - multiple", func(t *testing.T) {
		expectStatusCode := 200

		body := []index.Profile{
			createDataIndexProfile("1"),
			createDataIndexProfile("2"),
		}

		jsonBody, err := json.Marshal(body)
		assert.NoError(t, err)

		client := testsuit.CreateStatusCodeUrlValidatorWithBodyHttpClient(t, expectStatusCode, buildProfileUrl(""), jsonBody)

		response, err := testsuit.CreateFusionTestClient(client).IndexProfile().Getter().
			Do(context.Background())

		assert.NoError(t, err)
		assert.Equal(t, 2, len(response))
		assert.Equal(t, body[0].Id, response[0].Id)
		assert.Equal(t, body[0].IndexPipeline, response[0].IndexPipeline)
	})
}

func createDataIndexProfile(id string) index.Profile {
	return index.Profile{
		Id:            id,
		IndexPipeline: "test-pipeline",
	}
}
