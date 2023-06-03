package query

import (
	"context"
	"github.com/jvgrootveld/fusion-go-client/fusion/query"
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
			WithModel(query.Pipeline{
				Id: "pipeline-id",
				Stages: []stage.Stage{
					stage.NewTextTagger("Text Tagger"),
					stage.NewBoostWithSignals("Boost with Signals"),
					stage.NewQueryFields("Query Fields"),
					stage.NewFacets("Facets"),
					stage.NewApplyRules("Apply Rules"),
					stage.NewSolrQuery("SolrQuery"),
					stage.NewModifyResponseWithRules("Modify Response with Rules"),
				},
			}).
			Do(context.Background())

		assert.NoError(t, err)
	})
}
