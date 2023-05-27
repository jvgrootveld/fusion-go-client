package datamodel

import (
	"context"
	"github.com/jvgrootveld/fusion-go-client/fusion/datamodel"
	"testing"

	"github.com/jvgrootveld/fusion-go-client/test/testsuit"
	"github.com/stretchr/testify/assert"
)

func TestProfileCreator(t *testing.T) {
	t.Run("Data Model - Create", func(t *testing.T) {
		expectStatusCode := 201

		client := testsuit.CreateStatusCodeUrlValidatorHttpClient(t, expectStatusCode, buildUrl(""))

		err := testsuit.CreateFusionTestClient(client).DataModel().Creator().
			WithID("data-model-id").
			WithName("data-model").
			WithDescription("Data Model").
			WithIndexPipeline("pipeline-id").
			WithField(datamodel.DataModelField{
				Name:        "Name",
				Description: "Description",
			}).
			Do(context.Background())

		assert.NoError(t, err)
	})
}
