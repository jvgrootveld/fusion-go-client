package collection

import (
	"context"
	"testing"

	"github.com/jvgrootveld/fusion-go-client/test/testsuit"
	"github.com/stretchr/testify/assert"
)

func TestCollectionClearer(t *testing.T) {
	t.Run("Collection - Clear", func(t *testing.T) {
		expectStatusCode := 200
		id := "collection-id"

		url := testsuit.Scheme + "://" + testsuit.Host + "/api/solr/" + id + "/update?commit=true"

		client := testsuit.CreateStatusCodeUrlValidatorHttpClient(t, expectStatusCode, url)

		err := testsuit.CreateFusionTestClient(client).Collection().Clearer().
			WithID(id).
			Do(context.Background())

		assert.NoError(t, err)
	})
}
