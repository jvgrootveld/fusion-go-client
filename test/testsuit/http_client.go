package testsuit

import (
	"bytes"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

// RoundTripFunc .
type RoundTripFunc func(req *http.Request) *http.Response

// RoundTrip .
func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

// NewTestClient returns *http.Client with Transport replaced to avoid making real calls
func NewTestClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: fn,
	}
}

type TestFunc func(req *http.Request)

// CreateStatusCodeHttpClient creates a http.Client which returns given statusCode and
// calls the testFunc callback to validate sent request in the test.
func CreateStatusCodeHttpClient(statusCode int, body []byte, testFunc TestFunc) *http.Client {
	return NewTestClient(func(req *http.Request) *http.Response {
		// Call test callback
		testFunc(req)

		return &http.Response{
			StatusCode: statusCode,
			Body:       io.NopCloser(bytes.NewBuffer(body)),
			// Must be set to non-nil value or it panics
			Header: make(http.Header),
		}
	})
}

func CreateStatusCodeUrlValidatorHttpClient(t *testing.T, expectStatusCode int, expectedUrl string) *http.Client {
	return CreateStatusCodeHttpClient(expectStatusCode, []byte(""), func(req *http.Request) {
		assert.Equal(t, expectedUrl, req.URL.String())
	})
}

func CreateStatusCodeUrlValidatorWithBodyHttpClient(t *testing.T, expectStatusCode int, expectedUrl string, body []byte) *http.Client {
	return CreateStatusCodeHttpClient(expectStatusCode, body, func(req *http.Request) {
		assert.Equal(t, expectedUrl, req.URL.String())
	})
}
