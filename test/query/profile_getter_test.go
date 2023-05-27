package query

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/jvgrootveld/fusion-go-client/fusion/query"

	"github.com/jvgrootveld/fusion-go-client/test/testsuit"
	"github.com/stretchr/testify/assert"
)

func TestProfileGetter(t *testing.T) {
	t.Run("Query Profile - Getter - one", func(t *testing.T) {
		expectStatusCode := 200
		id := "pipeline-id"

		body := createDataQueryProfile("1")

		jsonBody, err := json.Marshal(body)
		assert.NoError(t, err)

		client := testsuit.CreateStatusCodeUrlValidatorWithBodyHttpClient(t, expectStatusCode, buildProfileUrl(id), jsonBody)

		response, err := testsuit.CreateFusionTestClient(client).QueryProfile().Getter().
			WithID(id).
			Do(context.Background())

		assert.NoError(t, err)
		assert.Equal(t, 1, len(response))
		assert.Equal(t, body.Id, response[0].Id)
		assert.Equal(t, body.QueryPipeline, response[0].QueryPipeline)
	})

	t.Run("Query Profile - Getter - multiple", func(t *testing.T) {
		expectStatusCode := 200

		body := []query.Profile{
			createDataQueryProfile("1"),
			createDataQueryProfile("2"),
		}

		jsonBody, err := json.Marshal(body)
		assert.NoError(t, err)

		client := testsuit.CreateStatusCodeUrlValidatorWithBodyHttpClient(t, expectStatusCode, buildProfileUrl(""), jsonBody)

		response, err := testsuit.CreateFusionTestClient(client).QueryProfile().Getter().
			Do(context.Background())

		assert.NoError(t, err)
		assert.Equal(t, 2, len(response))
		assert.Equal(t, body[0].Id, response[0].Id)
		assert.Equal(t, body[0].QueryPipeline, response[0].QueryPipeline)
	})
}

func createDataQueryProfile(id string) query.Profile {
	return query.Profile{
		Id:            id,
		QueryPipeline: "test-pipeline",
	}
}
