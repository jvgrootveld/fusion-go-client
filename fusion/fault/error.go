package fault

import "fmt"

// FusionClientError is returned if the client experienced an error.
//
//	If the error is due to Fusion returning an unexpected status code the IsUnexpectedStatusCode field will be true
//	 and the StatusCode field will be set
//	If the error occurred for another reason the DerivedFromError will be set and IsUnexpectedStatusCode will be false
type FusionClientError struct {
	IsUnexpectedStatusCode bool
	StatusCode             int
	Msg                    string
	DerivedFromError       error
}

// Error message of the FusionClientError
func (uce *FusionClientError) Error() string {
	msg := uce.Msg
	if uce.DerivedFromError != nil {
		msg = fmt.Sprintf("%s: %s", uce.Msg, uce.DerivedFromError.Error())
	}
	return fmt.Sprintf("status code: %v, error: %v", uce.StatusCode, msg)
}
