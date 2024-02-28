package job

import (
	"errors"
	"fmt"
	"github.com/jvgrootveld/fusion-go-client/fusion/fault"
)

func createJobValidator(model Job) error {
	typeName := fmt.Sprint(ApiName, "JobCreator")

	if model.Id == "" {
		return fault.NewRequiredError(typeName, "id")
	}

	return nil
}

func createSparkJobValidator(model Spark) error {
	typeName := fmt.Sprint(SparkApiName, "SparkCreator")

	switch t := model.(type) {
	case *SparkParallelBulkLoader:
		return createSparkParallelBulkLoaderValidator(typeName, model.(*SparkParallelBulkLoader))
	default:
		return errors.New(fmt.Sprintf("unsupported Spark type '%T' for creation. TODO: Add model and validator", t))
	}
}

func createSparkParallelBulkLoaderValidator(typeName string, model *SparkParallelBulkLoader) error {
	if model.Id == "" {
		return fault.NewRequiredError(typeName, "id")
	}

	if model.Updates != nil && len(model.Updates) > 0 {
		return fault.NewSetFieldError(typeName, "updates")
	}

	return nil
}
