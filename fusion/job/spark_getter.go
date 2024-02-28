package job

import (
	"context"
	"encoding/json"
	"github.com/jvgrootveld/fusion-go-client/fusion/connection"
	"github.com/jvgrootveld/fusion-go-client/fusion/except"
	"github.com/jvgrootveld/fusion-go-client/fusion/pathbuilder"
	"net/http"
)

// SparkGetter builder to get one or all configured objects
type SparkGetter struct {
	connection  *connection.Connection
	apiName     string
	application string
	id          string
}

// NewSparkGetter with connection and apiName of the configured object to get like index-pipeline or query-profile
func NewSparkGetter(connection *connection.Connection, apiName string) *SparkGetter {
	return &SparkGetter{
		connection: connection,
		apiName:    apiName,
	}
}

// ForApplication specifies the application the configured object is.
// When set it's interpreted as an application api e.g. `/api/apps/acme/spark/configurations`
// When empty it's interpreted as a generic api e.g. `/api/spark/configurations'
func (getter *SparkGetter) ForApplication(application string) *SparkGetter {
	getter.application = application
	return getter
}

// WithID specifies a specific configured object to get.
// When empty, all are returned.
func (getter *SparkGetter) WithID(id string) *SparkGetter {
	getter.id = id
	return getter
}

// Do get one or all configured objects
func (getter *SparkGetter) Do(ctx context.Context) ([]Spark, error) {
	path := pathbuilder.ApiPath(pathbuilder.Components{
		ApiName:     getter.apiName,
		Application: getter.application,
		ObjectId:    getter.id,
	})

	responseData, err := getter.connection.RunREST(ctx, path, http.MethodGet, nil)
	err = except.CheckResponseDataErrorAndStatusCode(responseData, err, 200)
	if err != nil {
		return nil, err
	}

	// If id is set, unmarshal to a single instance
	if getter.id != "" {
		var rawResult json.RawMessage
		err = responseData.DecodeBodyIntoTarget(&rawResult)
		if err != nil {
			return nil, err
		}

		spark, err := parseSpark(rawResult)

		return []Spark{spark}, err
	}

	var rawResults []json.RawMessage
	err = responseData.DecodeBodyIntoTarget(&rawResults)
	if err != nil {
		return nil, err
	}

	sparks := make([]Spark, len(rawResults))
	for i, raw := range rawResults {
		spark, err := parseSpark(raw)
		if err != nil {
			return nil, err
		}
		sparks[i] = spark
	}
	return sparks, err
}

// parseSpark unmarshal the raw json into the correct Spark type
func parseSpark(raw json.RawMessage) (Spark, error) {
	var data map[string]interface{}
	err := json.Unmarshal(raw, &data)
	if err != nil {
		return nil, err
	}

	sparkType := ""
	if t, ok := data["type"].(string); ok {
		sparkType = t
	}

	// Unmarshal again into the correct type
	var actual Spark
	switch sparkType {
	case SparkTypeParallelBulkLoader:
		actual = &SparkParallelBulkLoader{}
	default:
		actual = &SparkGeneric{}
	}

	err = json.Unmarshal(raw, actual)
	return actual, err
}
