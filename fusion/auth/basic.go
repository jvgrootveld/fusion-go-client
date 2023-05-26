package auth

import "net/http"

// BasicAuthConfig . Config for Basic Authentication
type BasicAuthConfig struct {
	username string
	password string
}

func NewBasicAuthConfig(username, password string) Config {
	return &BasicAuthConfig{
		username: username,
		password: password,
	}
}

func (config *BasicAuthConfig) AddAuth(request *http.Request) {
	request.SetBasicAuth(config.username, config.password)
}
