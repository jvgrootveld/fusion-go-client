package collection

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/jvgrootveld/fusion-go-client/fusion/collection"

	"github.com/jvgrootveld/fusion-go-client/test/testsuit"
	"github.com/stretchr/testify/assert"
)

func TestCollectionGetter(t *testing.T) {
	t.Run("Collection - Getter - one", func(t *testing.T) {
		expectStatusCode := 200
		id := "collection-id"

		body := createCollectionModel("1")

		jsonBody, err := json.Marshal(body)
		assert.NoError(t, err)

		client := testsuit.CreateStatusCodeUrlValidatorWithBodyHttpClient(t, expectStatusCode, buildUrl(id), jsonBody)

		response, err := testsuit.CreateFusionTestClient(client).Collection().Getter().
			WithID(id).
			Do(context.Background())

		assert.NoError(t, err)
		assert.Equal(t, 1, len(response))
		assert.Equal(t, body.Id, response[0].Id)
	})

	t.Run("Collection - Getter - multiple", func(t *testing.T) {
		expectStatusCode := 200

		body := []collection.Collection{
			createCollectionModel("1"),
			createCollectionModel("2"),
		}

		jsonBody, err := json.Marshal(body)
		assert.NoError(t, err)

		client := testsuit.CreateStatusCodeUrlValidatorWithBodyHttpClient(t, expectStatusCode, buildUrl(""), jsonBody)

		response, err := testsuit.CreateFusionTestClient(client).Collection().Getter().
			Do(context.Background())

		assert.NoError(t, err)
		assert.Equal(t, 2, len(response))
		assert.Equal(t, body[0].Id, response[0].Id)
	})
}

func createCollectionModel(id string) collection.Collection {
	return collection.Collection{
		Id: id,
	}
}
