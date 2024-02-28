package job

import (
	"context"
	"testing"

	"github.com/jvgrootveld/fusion-go-client/test/testsuit"
	"github.com/stretchr/testify/assert"
)

func TestSparkDeleter(t *testing.T) {
	t.Run("Spark - Delete", func(t *testing.T) {
		expectStatusCode := 200
		id := "spark-id"

		client := testsuit.CreateStatusCodeUrlValidatorHttpClient(t, expectStatusCode, buildSparkJobUrl(id))

		err := testsuit.CreateFusionTestClient(client).Job().Spark().Deleter().
			WithID(id).
			Do(context.Background())

		assert.NoError(t, err)
	})
}
