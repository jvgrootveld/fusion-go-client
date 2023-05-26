package auth

import "net/http"

// Config for authentication
type Config interface {
	// AddAuth to the request
	AddAuth(request *http.Request)
}
