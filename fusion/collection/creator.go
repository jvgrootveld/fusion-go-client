package collection

import (
	"fmt"
	"github.com/jvgrootveld/fusion-go-client/fusion/fault"
)

func createValidator(model CreateCollection) error {
	typeName := fmt.Sprint(ApiName, "Creator")

	if model.Id == "" {
		return fault.NewRequiredError(typeName, "object.Id")
	}

	return nil
}
