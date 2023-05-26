package except

import (
	"fmt"

	"github.com/jvgrootveld/fusion-go-client/fusion/fault"

	"github.com/jvgrootveld/fusion-go-client/fusion/connection"
)

// NewFusionClientError from status code and error message
func NewFusionClientError(statusCode int, format string, args ...interface{}) *fault.FusionClientError {
	return &fault.FusionClientError{
		IsUnexpectedStatusCode: true,
		StatusCode:             statusCode,
		Msg:                    fmt.Sprintf(format, args...),
	}
}

// NewDerivedFusionClientError wraps an error into a FusionClientError as derived error
func NewDerivedFusionClientError(err error) *fault.FusionClientError {
	return &fault.FusionClientError{
		IsUnexpectedStatusCode: false,
		StatusCode:             -1,
		Msg:                    "check the DerivedFromError field for more information",
		DerivedFromError:       err,
	}
}

// NewUnexpectedStatusCodeErrorFromRESTResponse creates the error based on a response data object
func NewUnexpectedStatusCodeErrorFromRESTResponse(responseData *connection.ResponseData) *fault.FusionClientError {
	return NewFusionClientError(responseData.StatusCode, string(responseData.Body))
}

// CheckResponseDataErrorAndStatusCode returns the response error if it is not nil,
//
//	and an FusionClientError if the status code is not matching
func CheckResponseDataErrorAndStatusCode(responseData *connection.ResponseData, responseErr error, expectedStatusCodes ...int) error {
	if responseErr != nil {
		return NewDerivedFusionClientError(responseErr)
	}
	for _, expectedStatusCode := range expectedStatusCodes {
		if responseData.StatusCode == expectedStatusCode {
			return nil
		}
	}
	return NewUnexpectedStatusCodeErrorFromRESTResponse(responseData)
}
