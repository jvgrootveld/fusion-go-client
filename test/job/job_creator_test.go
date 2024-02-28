package job

import (
	"context"
	"github.com/jvgrootveld/fusion-go-client/fusion/job"
	"testing"

	"github.com/jvgrootveld/fusion-go-client/test/testsuit"
	"github.com/stretchr/testify/assert"
)

func TestJobCreator(t *testing.T) {
	t.Run("Job - Create", func(t *testing.T) {
		expectStatusCode := 201

		client := testsuit.CreateStatusCodeUrlValidatorHttpClient(t, expectStatusCode, buildJobUrl(""))

		err := testsuit.CreateFusionTestClient(client).Job().Creator().
			WithModel(job.Job{
				Id:      "job-id",
				Enabled: true,
				Status:  "ready",
				Extra:   map[string]interface{}{},
			}).
			Do(context.Background())

		assert.NoError(t, err)
	})
}
