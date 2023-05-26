package connection

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/jvgrootveld/fusion-go-client/fusion/auth"
)

// Connection networking layer accessing Fusion using http requests
type Connection struct {
	basePath   string
	httpClient *http.Client
	auth       auth.Config
	headers    map[string]string
	doneCh     chan bool
}

// NewConnection based on scheme://host
// If httpClient is nil a default client will be used
func NewConnection(scheme string, host string, httpClient *http.Client, auth auth.Config, headers map[string]string) *Connection {
	client := httpClient
	if client == nil {
		client = &http.Client{}
	}

	connection := &Connection{
		basePath:   scheme + "://" + host,
		httpClient: client,
		auth:       auth,
		headers:    headers,
	}

	return connection
}

func (conn *Connection) marshalBody(body interface{}) (io.Reader, error) {
	if body == nil {
		return nil, nil
	}
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	return bytes.NewBuffer(jsonBody), nil
}

func (conn *Connection) addAuthToRequest(request *http.Request) {
	if conn.auth != nil {
		conn.auth.AddAuth(request)
	}
}

func (conn *Connection) addHeaderToRequest(request *http.Request) {
	for k, v := range conn.headers {
		request.Header.Add(k, v)
	}
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")
}

func (conn *Connection) createRequest(
	ctx context.Context,
	path string,
	restMethod string,
	body interface{},
) (*http.Request, error) {
	url := conn.basePath + path // Create the URL

	jsonBody, err := conn.marshalBody(body)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest(restMethod, url, jsonBody)
	if err != nil {
		return nil, err
	}

	conn.addAuthToRequest(request)
	conn.addHeaderToRequest(request)

	return request.WithContext(ctx), nil
}

// RunREST executes a http request
// path: expects a resource path e.g. `/api/apps/acme/index-pipelines/test-pipeline`
// restMethod: as they are defined in constants in the *http* package
// Returns:
//
//	a response that may be parsed into a struct after the fact
//	error if there was a network issue
func (conn *Connection) RunREST(
	ctx context.Context,
	path string,
	restMethod string,
	requestBody interface{},
) (*ResponseData, error) {
	request, requestErr := conn.createRequest(ctx, path, restMethod, requestBody)
	if requestErr != nil {
		return nil, requestErr
	}
	response, responseErr := conn.httpClient.Do(request)
	if responseErr != nil {
		return nil, responseErr
	}

	defer response.Body.Close()
	body, bodyErr := io.ReadAll(response.Body)
	if bodyErr != nil {
		return nil, bodyErr
	}

	return &ResponseData{
		Body:       body,
		StatusCode: response.StatusCode,
	}, nil
}

// ResponseData encapsulation of the http request body and status
type ResponseData struct {
	Body       []byte
	StatusCode int
}
