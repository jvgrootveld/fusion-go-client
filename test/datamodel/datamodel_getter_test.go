package datamodel

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/jvgrootveld/fusion-go-client/fusion/datamodel"

	"github.com/jvgrootveld/fusion-go-client/test/testsuit"
	"github.com/stretchr/testify/assert"
)

func TestProfileGetter(t *testing.T) {
	t.Run("Data Model - Getter - one", func(t *testing.T) {
		expectStatusCode := 200
		id := "data-model-id"

		body := createDataDataModel("1")

		jsonBody, err := json.Marshal(body)
		assert.NoError(t, err)

		client := testsuit.CreateStatusCodeUrlValidatorWithBodyHttpClient(t, expectStatusCode, buildUrl(id), jsonBody)

		response, err := testsuit.CreateFusionTestClient(client).DataModel().Getter().
			WithID(id).
			Do(context.Background())

		assert.NoError(t, err)
		assert.Equal(t, 1, len(response))
		assert.Equal(t, body.Id, response[0].Id)
		assert.Equal(t, body.Name, response[0].Name)
	})

	t.Run("Data Model - Getter - multiple", func(t *testing.T) {
		expectStatusCode := 200

		body := []datamodel.DataModel{
			createDataDataModel("1"),
			createDataDataModel("2"),
		}

		jsonBody, err := json.Marshal(body)
		assert.NoError(t, err)

		client := testsuit.CreateStatusCodeUrlValidatorWithBodyHttpClient(t, expectStatusCode, buildUrl(""), jsonBody)

		response, err := testsuit.CreateFusionTestClient(client).DataModel().Getter().
			Do(context.Background())

		assert.NoError(t, err)
		assert.Equal(t, 2, len(response))
		assert.Equal(t, body[0].Id, response[0].Id)
		assert.Equal(t, body[0].Name, response[0].Name)
	})
}

func createDataDataModel(id string) datamodel.DataModel {
	return datamodel.DataModel{
		Id:            id,
		Name:          "test data model",
		Description:   "Test data model",
		IndexPipeline: "test-pipeline",
	}
}
