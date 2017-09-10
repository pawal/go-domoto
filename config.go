// Package domoto adds a Golang interface for the Domoticz JSON HTTP API.
// All API calls are done through the handle returned by New.
// The Go API implements parts of the Domoticz API documentation:
// https://www.domoticz.com/wiki/Domoticz_API/JSON_URL%27s
package domoto

import (
	"encoding/base64"
)

const (
	mediaType = "application/json"
)

// Config is the configuration for executing calls to the Domoticz
// JSON HTTP API.
type Config struct {
	BaseURL string
	secret  string // encoded as user:pass by New()
}

// New creates a HTTP handler for calling the Domoticz JSON API.
// Add the base URL, username and password, these will be used for
// any calls using the handler.
func New(baseURL, user, pass string) *Config {
	if baseURL == "" {
		baseURL = "http://localhost:8080" // default base URL
	}
	var secret string
	if user != "" && pass != "" {
		secret = base64.StdEncoding.EncodeToString([]byte(user + ":" + pass))
	}
	return &Config{BaseURL: baseURL, secret: secret}
}
