package query

import (
	"context"
	"testing"

	"github.com/jvgrootveld/fusion-go-client/fusion/query"

	"github.com/jvgrootveld/fusion-go-client/test/testsuit"
	"github.com/stretchr/testify/assert"
)

func TestCreator(t *testing.T) {
	t.Run("Query Profile - Create", func(t *testing.T) {
		expectStatusCode := 201

		client := testsuit.CreateStatusCodeUrlValidatorHttpClient(t, expectStatusCode, buildProfileUrl(""))

		err := testsuit.CreateFusionTestClient(client).QueryProfile().Creator().
			WithID("profile-id").
			WithQueryPipeline("pipeline-id").
			WithCollection("collection-id").
			WithParam(query.Param{Key: "Key", Value: "Value", Policy: query.Append}).
			WithAdditionalProperty("property-key", "property value").
			Do(context.Background())

		assert.NoError(t, err)
	})
}
