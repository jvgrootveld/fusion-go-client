package generic

import (
	"context"
	"fmt"
	"net/http"

	"github.com/jvgrootveld/fusion-go-client/fusion/fault"

	"github.com/jvgrootveld/fusion-go-client/fusion/connection"
	"github.com/jvgrootveld/fusion-go-client/fusion/except"
	"github.com/jvgrootveld/fusion-go-client/fusion/pathbuilder"
)

// Deleter builder to delete a configured object like an Index or query pipeline / profile
type Deleter struct {
	connection  *connection.Connection
	apiName     string
	application string
	id          string
}

// NewDeleter with connection and apiName of the configured object to delete like index-pipeline or query-profile
func NewDeleter(connection *connection.Connection, apiName string) *Deleter {
	return &Deleter{
		connection: connection,
		apiName:    apiName,
	}
}

// ForApplication specifies the application the configured object is in.
// When set it's interpreted as an application api e.g. `/api/apps/acme/index-pipelines`
// When empty it's interpreted as a generic api e.g. `/api/collections'
func (deleter *Deleter) ForApplication(application string) *Deleter {
	deleter.application = application
	return deleter
}

func (deleter *Deleter) WithID(id string) *Deleter {
	deleter.id = id
	return deleter
}

// Do delete the specified configured object from Fusion
func (deleter *Deleter) Do(ctx context.Context) error {
	err := deleter.checkRequired()
	if err != nil {
		return err
	}

	path := pathbuilder.ApiPath(pathbuilder.Components{
		Application: deleter.application,
		ApiName:     deleter.apiName,
		ObjectId:    deleter.id,
	})
	responseData, err := deleter.connection.RunREST(ctx, path, http.MethodDelete, nil)
	return except.CheckResponseDataErrorAndStatusCode(responseData, err, 204, 200)
}

func (deleter *Deleter) checkRequired() error {
	typeName := fmt.Sprint(deleter.apiName, "Deleter")

	if deleter.id == "" {
		return fault.NewRequiredError(typeName, "id")
	}

	return nil
}
