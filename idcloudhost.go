package idcloudhost

import (
	"net/http"
)

// Make a new client
func New(config *Config) (*Client, error) {
	return &Client{
		apikey: config.ApiKey,
		http:   &http.Client{
			// Timeout: time.Duration(30 * time.Second),
		},
	}, nil
}
