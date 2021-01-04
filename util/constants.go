package util

import "time"

// Credentials username/password retrieved from yaml config file
type Credentials struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

// ErrorInfo generic failed http response
type ErrorInfo struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// AuthInfo response to get access token
type AuthInfo struct {
	AccessToken    string `json:"access_token"`
	ExpiresIn      int    `json:"expires_in"`
	RefreshToken   string `json:"refresh_token"`
	ExpirationTime time.Time
}

// RefreshInfo response to refresh access token
type RefreshInfo struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

const (
	// GetAccessTokenURL url to get access token
	GetAccessTokenURL = "https://rest.sandbox.karhoo.com/v1/auth/token"
	// RefreshAccessTokenURL url to refresh access token
	RefreshAccessTokenURL = "https://rest.sandbox.karhoo.com/v1/auth/refresh"
)
