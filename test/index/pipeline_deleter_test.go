package index

import (
	"context"
	"testing"

	"github.com/jvgrootveld/fusion-go-client/test/testsuit"
	"github.com/stretchr/testify/assert"
)

func TestPipelineDeleter(t *testing.T) {
	t.Run("Index Pipeline - Delete", func(t *testing.T) {
		expectStatusCode := 204
		id := "pipeline-id"

		client := testsuit.CreateStatusCodeUrlValidatorHttpClient(t, expectStatusCode, buildPipelineUrl(id))

		err := testsuit.CreateFusionTestClient(client).IndexPipeline().Deleter().
			WithID(id).
			Do(context.Background())

		assert.NoError(t, err)
	})
}
