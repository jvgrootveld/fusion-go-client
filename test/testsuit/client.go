package testsuit

import (
	"net/http"

	"github.com/jvgrootveld/fusion-go-client/fusion"
	"github.com/jvgrootveld/fusion-go-client/fusion/pathbuilder"
)

const (
	host        = "fusion.unittest"
	scheme      = "http"
	application = "test-app"
)

// CreateApplicationUrl creates an url for tests with an application in the path
func CreateApplicationUrl(apiName, id string) string {
	path := pathbuilder.AppsPath(pathbuilder.Components{
		Application: application,
		ApiName:     apiName,
		ObjectId:    id,
	})

	return scheme + "://" + host + path
}

// CreateFusionTestClient creates a Fusion Client with given http client
func CreateFusionTestClient(client *http.Client) *fusion.Client {
	fusionClient, err := fusion.NewClient(fusion.Config{
		Host:             host,
		Scheme:           scheme,
		Application:      application,
		ConnectionClient: client,
	})
	if err != nil {
		panic(err)
	}
	return fusionClient
}
