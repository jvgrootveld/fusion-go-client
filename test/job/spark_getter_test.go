package job

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/jvgrootveld/fusion-go-client/fusion/job"

	"github.com/jvgrootveld/fusion-go-client/test/testsuit"
	"github.com/stretchr/testify/assert"
)

func TestSparkGetter(t *testing.T) {
	t.Run("Spark - SparkGetter - one", func(t *testing.T) {
		expectStatusCode := 200
		id := "spark-id"

		body := createDataSpark("1")

		jsonBody, err := json.Marshal(body)
		assert.NoError(t, err)

		client := testsuit.CreateStatusCodeUrlValidatorWithBodyHttpClient(t, expectStatusCode, buildSparkJobUrl(id), jsonBody)

		response, err := testsuit.CreateFusionTestClient(client).Job().Spark().Getter().
			WithID(id).
			Do(context.Background())

		responseItem := response[0].(*job.SparkParallelBulkLoader)

		assert.NoError(t, err)
		assert.Equal(t, 1, len(response))
		assert.Equal(t, body.Id, responseItem.Id)
	})

	t.Run("Spark - SparkGetter - multiple", func(t *testing.T) {
		expectStatusCode := 200

		firstBodyItem := createDataSpark("1")
		body := []job.Spark{
			firstBodyItem,
			createDataSpark("2"),
		}

		jsonBody, err := json.Marshal(body)
		assert.NoError(t, err)

		client := testsuit.CreateStatusCodeUrlValidatorWithBodyHttpClient(t, expectStatusCode, buildSparkJobUrl(""), jsonBody)

		response, err := testsuit.CreateFusionTestClient(client).Job().Spark().Getter().
			Do(context.Background())

		firstResponseItem := response[0].(*job.SparkParallelBulkLoader)

		assert.NoError(t, err)
		assert.Equal(t, 2, len(response))
		assert.Equal(t, firstBodyItem.Id, firstResponseItem.Id)
		assert.Equal(t, firstBodyItem.Type, firstResponseItem.Type)
	})
}

func createDataSpark(id string) job.SparkParallelBulkLoader {
	return job.SparkParallelBulkLoader{
		SparkGeneric: job.SparkGeneric{
			Id:   id,
			Type: job.SparkTypeParallelBulkLoader,
		},
	}
}
