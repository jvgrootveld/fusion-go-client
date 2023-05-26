package query

import (
	"context"
	"testing"

	"github.com/jvgrootveld/fusion-go-client/fusion/query/stage"

	"github.com/jvgrootveld/fusion-go-client/test/testsuit"
	"github.com/stretchr/testify/assert"
)

func TestPipelineCreator(t *testing.T) {
	t.Run("Query Pipeline - Create", func(t *testing.T) {
		expectStatusCode := 201

		client := testsuit.CreateStatusCodeUrlValidatorHttpClient(t, expectStatusCode, buildPipelineUrl(""))

		err := testsuit.CreateFusionTestClient(client).QueryPipeline().Creator().
			WithID("pipeline-id").
			WithStages(
				stage.NewTextTagger("Text Tagger"),
				stage.NewBoostWithSignals("Boost with Signals"),
				stage.NewQueryFields("Query Fields"),
				stage.NewFacets("Facets"),
				stage.NewApplyRules("Apply Rules"),
				stage.NewSolrQuery("SolrQuery"),
				stage.NewModifyResponseWithRules("Modify Response with Rules"),
			).
			Do(context.Background())

		assert.NoError(t, err)
	})
}
