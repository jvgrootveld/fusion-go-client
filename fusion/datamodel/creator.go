package datamodel

import (
	"fmt"
	"github.com/jvgrootveld/fusion-go-client/fusion/fault"
)

func createValidator(model DataModel) error {
	typeName := fmt.Sprint(ApiName, "Creator")

	if model.Id == "" {
		return fault.NewRequiredError(typeName, "object.Id")
	}
	if model.Name == "" {
		return fault.NewRequiredError(typeName, "object.Name")
	}
	if model.Description == "" {
		return fault.NewRequiredError(typeName, "object.Description")
	}
	if model.IndexPipeline == "" {
		return fault.NewRequiredError(typeName, "object.IndexPipeline")
	}

	return nil
}
