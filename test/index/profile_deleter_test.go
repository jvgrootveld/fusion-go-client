package index

import (
	"context"
	"testing"

	"github.com/jvgrootveld/fusion-go-client/test/testsuit"
	"github.com/stretchr/testify/assert"
)

func TestProfileDeleter(t *testing.T) {
	t.Run("Index Profile - Delete", func(t *testing.T) {
		expectStatusCode := 204
		id := "profile-id"

		client := testsuit.CreateStatusCodeUrlValidatorHttpClient(t, expectStatusCode, buildProfileUrl(id))

		err := testsuit.CreateFusionTestClient(client).IndexProfile().Deleter().
			WithID(id).
			Do(context.Background())

		assert.NoError(t, err)
	})
}
