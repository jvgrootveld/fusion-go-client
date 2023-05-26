package index

import (
	"context"
	"testing"

	"github.com/jvgrootveld/fusion-go-client/test/testsuit"
	"github.com/stretchr/testify/assert"
)

func TestProfileCreator(t *testing.T) {
	t.Run("Index Profile - Create", func(t *testing.T) {
		expectStatusCode := 201

		client := testsuit.CreateStatusCodeUrlValidatorHttpClient(t, expectStatusCode, buildProfileUrl(""))

		err := testsuit.CreateFusionTestClient(client).IndexProfile().Creator().
			WithID("profile-id").
			WithIndexPipeline("pipeline-id").
			WithCollection("collection-id").
			WithParser("parser-id").
			WithAdditionalProperty("property-key", "property value").
			Do(context.Background())

		assert.NoError(t, err)
	})
}
