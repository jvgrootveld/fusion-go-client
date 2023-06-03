package collection

import (
	"context"
	"github.com/jvgrootveld/fusion-go-client/fusion/collection"
	"testing"

	"github.com/jvgrootveld/fusion-go-client/test/testsuit"
	"github.com/stretchr/testify/assert"
)

func TestCollectionCreator(t *testing.T) {
	t.Run("Collection - Create", func(t *testing.T) {
		expectStatusCode := 201

		client := testsuit.CreateStatusCodeUrlValidatorHttpClient(t, expectStatusCode, buildUrl(""))

		err := testsuit.CreateFusionTestClient(client).Collection().Creator().
			WithModel(collection.CreateCollection{
				Id:              "collection-id",
				SearchClusterId: "default",
				SolrParams: collection.SolrParam{
					NumShards:         1,
					ReplicationFactor: 1,
				}},
			).
			Do(context.Background())

		assert.NoError(t, err)
	})
}
