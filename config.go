// Package glair provides common objects that can be used
// across all GLAIR Vision Go SDK packages
package glair

import (
	"net/http"
	"net/url"
	"strings"
)

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// Config provides configuration for a client instance,
// including credentials and API configuration such as
// base URL and API version.
type Config struct {
	Username string
	Password string
	ApiKey   string

	BaseUrl    string
	ApiVersion string

	Client HTTPClient
}

// New creates a new configuration object with default values for
// base URL and API Version.
func NewConfig(username string, password string, apiKey string) *Config {
	defaultUrl := "https://api.vision.glair.ai"
	defaultApiVersion := "v1"
	defaultClient := http.DefaultClient

	return &Config{
		Username:   username,
		Password:   password,
		ApiKey:     apiKey,
		BaseUrl:    defaultUrl,
		ApiVersion: defaultApiVersion,
		Client:     defaultClient,
	}
}

// GetEndpointURL creates service URL with base URL and API version
func (c *Config) GetEndpointURL(service string, endpoint string) (*url.URL, error) {
	parts := []string{c.BaseUrl, service, c.ApiVersion, endpoint}
	return url.Parse(strings.Join(parts, "/"))
}

// WithCredentials sets user credentials for the configuration object
func (c *Config) WithCredentials(username string, password string, apiKey string) *Config {
	c.Username = username
	c.Password = password
	c.ApiKey = apiKey

	return c
}

// WithClient sets HTTP client for the configuration object
// that will be used for calling GLAIR Vision API endpoints
func (c *Config) WithClient(client HTTPClient) *Config {
	c.Client = client

	return c
}

// WithBaseURL sets base API URL for the configuration object
func (c *Config) WithBaseURL(url string) *Config {
	c.BaseUrl = url

	return c
}

// WithVersion sets API version for the configuration object
func (c *Config) WithVersion(version string) *Config {
	c.ApiVersion = version

	return c
}
