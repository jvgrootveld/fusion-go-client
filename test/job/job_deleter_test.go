package job

import (
	"context"
	"testing"

	"github.com/jvgrootveld/fusion-go-client/test/testsuit"
	"github.com/stretchr/testify/assert"
)

func TestJobDeleter(t *testing.T) {
	t.Run("Job - Delete", func(t *testing.T) {
		expectStatusCode := 204
		id := "job-id"

		client := testsuit.CreateStatusCodeUrlValidatorHttpClient(t, expectStatusCode, buildJobUrl(id))

		err := testsuit.CreateFusionTestClient(client).Job().Deleter().
			WithID(id).
			Do(context.Background())

		assert.NoError(t, err)
	})
}
