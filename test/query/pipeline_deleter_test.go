package query

import (
	"context"
	"testing"

	"github.com/jvgrootveld/fusion-go-client/test/testsuit"
	"github.com/stretchr/testify/assert"
)

func TestPipelineDeleter(t *testing.T) {
	t.Run("Query Pipeline - Delete", func(t *testing.T) {
		expectStatusCode := 204
		pipelineId := "pipeline-id"

		client := testsuit.CreateStatusCodeUrlValidatorHttpClient(t, expectStatusCode, buildPipelineUrl(pipelineId))

		err := testsuit.CreateFusionTestClient(client).QueryPipeline().Deleter().
			WithID(pipelineId).
			Do(context.Background())

		assert.NoError(t, err)
	})
}