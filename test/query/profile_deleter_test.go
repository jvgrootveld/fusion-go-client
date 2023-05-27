package query

import (
	"context"
	"testing"

	"github.com/jvgrootveld/fusion-go-client/test/testsuit"
	"github.com/stretchr/testify/assert"
)

func TestDeleter(t *testing.T) {
	t.Run("Query Profile - Delete", func(t *testing.T) {
		expectStatusCode := 204
		id := "profile-id"

		client := testsuit.CreateStatusCodeUrlValidatorHttpClient(t, expectStatusCode, buildProfileUrl(id))

		err := testsuit.CreateFusionTestClient(client).QueryProfile().Deleter().
			WithID(id).
			Do(context.Background())

		assert.NoError(t, err)
	})
}
