package query

import (
	"fmt"
	"github.com/jvgrootveld/fusion-go-client/fusion/fault"
)

func createPipelineValidator(model Pipeline) error {
	typeName := fmt.Sprint(PipelineApiName, "Creator")

	if model.Id == "" {
		return fault.NewRequiredError(typeName, "id")
	}

	return nil
}

func createProfileValidator(model Profile) error {
	typeName := fmt.Sprint(ProfileApiName, "Creator")

	if model.Id == "" {
		return fault.NewRequiredError(typeName, "id")
	}
	if model.QueryPipeline == "" {
		return fault.NewRequiredError(typeName, "queryPipeline")
	}
	if model.Collection == "" {
		return fault.NewRequiredError(typeName, "collection")
	}

	return nil
}
