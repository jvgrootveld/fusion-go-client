package index

import (
	"context"
	"github.com/jvgrootveld/fusion-go-client/fusion/index"
	"testing"

	"github.com/jvgrootveld/fusion-go-client/test/testsuit"
	"github.com/stretchr/testify/assert"
)

func TestProfileCreator(t *testing.T) {
	t.Run("Index Profile - Create", func(t *testing.T) {
		expectStatusCode := 201

		client := testsuit.CreateStatusCodeUrlValidatorHttpClient(t, expectStatusCode, buildProfileUrl(""))

		err := testsuit.CreateFusionTestClient(client).IndexProfile().Creator().
			WithModel(index.Profile{
				Id:            "profile-id",
				IndexPipeline: "pipeline-id",
				Collection:    "collection-id",
				Parser:        "parser-id",
				AdditionalProperties: map[string]string{
					"property-key": "property value",
				},
			}).
			Do(context.Background())

		assert.NoError(t, err)
	})
}
