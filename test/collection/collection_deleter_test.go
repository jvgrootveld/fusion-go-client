package collection

import (
	"context"
	"testing"

	"github.com/jvgrootveld/fusion-go-client/test/testsuit"
	"github.com/stretchr/testify/assert"
)

func TestCollectionDeleter(t *testing.T) {
	t.Run("Collection - Delete", func(t *testing.T) {
		expectStatusCode := 204
		id := "collection-id"

		client := testsuit.CreateStatusCodeUrlValidatorHttpClient(t, expectStatusCode, buildUrl(id))

		err := testsuit.CreateFusionTestClient(client).Collection().Deleter().
			WithID(id).
			Do(context.Background())

		assert.NoError(t, err)
	})
}
