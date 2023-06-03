package testsuit

import (
	"net/http"

	"github.com/jvgrootveld/fusion-go-client/fusion"
	"github.com/jvgrootveld/fusion-go-client/fusion/pathbuilder"
)

const (
	Host        = "fusion.unittest"
	Scheme      = "http"
	Application = "test-app"
)

// CreateGenericApiUrl creates an url for tests without an Application in the path
func CreateGenericApiUrl(apiName, id string) string {
	path := pathbuilder.ApiPath(pathbuilder.Components{
		ApiName:  apiName,
		ObjectId: id,
	})

	return Scheme + "://" + Host + path
}

// CreateApplicationUrl creates an url for tests with an Application in the path
func CreateApplicationUrl(apiName, id string) string {
	path := pathbuilder.ApiPath(pathbuilder.Components{
		Application: Application,
		ApiName:     apiName,
		ObjectId:    id,
	})

	return Scheme + "://" + Host + path
}

// CreateFusionTestClient creates a Fusion Client with given http client
func CreateFusionTestClient(client *http.Client) *fusion.Client {
	fusionClient, err := fusion.NewClient(fusion.Config{
		Host:             Host,
		Scheme:           Scheme,
		Application:      Application,
		ConnectionClient: client,
	})
	if err != nil {
		panic(err)
	}
	return fusionClient
}
