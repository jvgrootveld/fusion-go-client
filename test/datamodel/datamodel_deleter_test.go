package datamodel

import (
	"context"
	"testing"

	"github.com/jvgrootveld/fusion-go-client/test/testsuit"
	"github.com/stretchr/testify/assert"
)

func TestDataModelDeleter(t *testing.T) {
	t.Run("Data Model - Delete", func(t *testing.T) {
		expectStatusCode := 204
		id := "data-model-id"

		client := testsuit.CreateStatusCodeUrlValidatorHttpClient(t, expectStatusCode, buildUrl(id))

		err := testsuit.CreateFusionTestClient(client).DataModel().Deleter().
			WithID(id).
			Do(context.Background())

		assert.NoError(t, err)
	})
}
