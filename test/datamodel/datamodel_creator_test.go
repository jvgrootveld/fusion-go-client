package datamodel

import (
	"context"
	"testing"

	"github.com/jvgrootveld/fusion-go-client/fusion/datamodel"

	"github.com/jvgrootveld/fusion-go-client/test/testsuit"
	"github.com/stretchr/testify/assert"
)

func TestDataModelCreator(t *testing.T) {
	t.Run("Data Model - Create", func(t *testing.T) {
		expectStatusCode := 201

		client := testsuit.CreateStatusCodeUrlValidatorHttpClient(t, expectStatusCode, buildUrl(""))

		err := testsuit.CreateFusionTestClient(client).DataModel().Creator().
			WithModel(datamodel.DataModel{
				Id:            "data-model-id",
				Name:          "data-model",
				Description:   "Data Model",
				IndexPipeline: "pipeline-id",
				Fields: []datamodel.Field{
					{
						Name:        "Field 1",
						Description: "Field 1 Description",
						Required:    true,
						Mappings: []datamodel.FieldMapping{
							{
								SolrField:   "SolrField",
								QueryField:  true,
								PhraseMatch: false,
								BoostValue:  0.1,
								PhraseBoost: 0.2,
							},
						},
					},
				},
			}).
			Do(context.Background())

		assert.NoError(t, err)
	})
}
