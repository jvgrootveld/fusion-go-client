package job

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/jvgrootveld/fusion-go-client/fusion/job"

	"github.com/jvgrootveld/fusion-go-client/test/testsuit"
	"github.com/stretchr/testify/assert"
)

func TestJobGetter(t *testing.T) {
	t.Run("Job - Getter - one", func(t *testing.T) {
		expectStatusCode := 200
		id := "job-id"

		body := createDataJob("1")

		jsonBody, err := json.Marshal(body)
		assert.NoError(t, err)

		client := testsuit.CreateStatusCodeUrlValidatorWithBodyHttpClient(t, expectStatusCode, buildJobUrl(id), jsonBody)

		response, err := testsuit.CreateFusionTestClient(client).Job().Getter().
			WithID(id).
			Do(context.Background())

		assert.NoError(t, err)
		assert.Equal(t, 1, len(response))
		assert.Equal(t, body.Id, response[0].Id)
	})

	t.Run("Job - Getter - multiple", func(t *testing.T) {
		expectStatusCode := 200

		body := []job.Job{
			createDataJob("1"),
			createDataJob("2"),
		}

		jsonBody, err := json.Marshal(body)
		assert.NoError(t, err)

		client := testsuit.CreateStatusCodeUrlValidatorWithBodyHttpClient(t, expectStatusCode, buildJobUrl(""), jsonBody)

		response, err := testsuit.CreateFusionTestClient(client).Job().Getter().
			Do(context.Background())

		assert.NoError(t, err)
		assert.Equal(t, 2, len(response))
		assert.Equal(t, body[0].Id, response[0].Id)
		assert.Equal(t, body[0].Status, response[0].Status)
		assert.Equal(t, body[0].Enabled, response[0].Enabled)
	})
}

func createDataJob(id string) job.Job {
	return job.Job{
		Id: id,
	}
}
