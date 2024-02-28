package job

import (
	"context"
	"fmt"
	"github.com/jvgrootveld/fusion-go-client/fusion/job"
	"testing"

	"github.com/jvgrootveld/fusion-go-client/test/testsuit"
	"github.com/stretchr/testify/assert"
)

func TestSparkCreator(t *testing.T) {
	t.Run("Spark - Create", func(t *testing.T) {
		expectStatusCode := 201

		client := testsuit.CreateStatusCodeUrlValidatorHttpClient(t, expectStatusCode, buildSparkJobUrl(""))

		err := testsuit.CreateFusionTestClient(client).Job().Spark().Creator().
			WithModel(&job.SparkParallelBulkLoader{
				SparkGeneric: job.SparkGeneric{
					Id:   "job-id",
					Type: job.SparkTypeParallelBulkLoader,
				},
				Config:                       nil,
				Format:                       "",
				ReadOptions:                  nil,
				OutputCollection:             "",
				OutputIndexPipeline:          "",
				ClearDatasource:              false,
				DefineFieldsUsingInputSchema: false,
				AtomicUpdates:                false,
				TransformScala:               "",
				CacheAfterRead:               false,
				ContinueAfterFailure:         false,
			}).
			Do(context.Background())

		assert.NoError(t, err)
	})

	t.Run("Spark - Create - Unsuported Type", func(t *testing.T) {
		expectStatusCode := 201

		client := testsuit.CreateStatusCodeUrlValidatorHttpClient(t, expectStatusCode, buildSparkJobUrl(""))

		model := &job.SparkGeneric{
			Id:   "job-id",
			Type: "unsupported-type",
		}
		err := testsuit.CreateFusionTestClient(client).Job().Spark().Creator().
			WithModel(model).
			Do(context.Background())

		assert.EqualError(t, err, fmt.Sprintf("unsupported Spark type '%T' for creation. TODO: Add model and validator", model))
	})
}
