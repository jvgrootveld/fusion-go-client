package index

import (
	"context"
	"testing"

	"github.com/jvgrootveld/fusion-go-client/fusion/index/stage"

	"github.com/jvgrootveld/fusion-go-client/test/testsuit"
	"github.com/stretchr/testify/assert"
)

func TestPipelineCreator(t *testing.T) {
	t.Run("Index Pipeline - Create", func(t *testing.T) {
		expectStatusCode := 201

		client := testsuit.CreateStatusCodeUrlValidatorHttpClient(t, expectStatusCode, buildPipelineUrl(""))

		err := testsuit.CreateFusionTestClient(client).IndexPipeline().Creator().
			WithID("pipeline-id").
			WithStages(
				stage.NewFieldMapping("Field Mapping"),
				stage.NewSolrDynamicFieldNameMapping("Solr Dynamic Field Mapping"),
				stage.NewSolrIndex("Solr Index"),
			).
			Do(context.Background())

		assert.NoError(t, err)
	})
}
