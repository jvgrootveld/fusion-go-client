package auth

import "net/http"

// JwtConfig . Config for JWT Authentication
type JwtConfig struct {
	token string
}

func NewJwtConfig(token string) Config {
	return &JwtConfig{
		token: token,
	}
}

func (config *JwtConfig) AddAuth(request *http.Request) {
	request.Header.Set("Authorization", "Basic "+config.token)
}
