package glair

import (
	"net/http"
	"strings"
)

const (
	defaultUrl     = "https://api.vision.glair.ai"
	defaultVersion = "v1"
)

// HTTPClient is an interface that users can implement
// to customize HTTP calls behavior when interacting
// with GLAIR Vision API
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// Config provides configuration for a client instance,
// including credentials and API configuration such as
// base URL and API version.
type Config struct {
	// Username represents username to be used for basic authentication
	// with GLAIR Vision API
	Username string
	// Password represents password to be used for basic authentication
	// with GLAIR Vision API
	Password string
	// API key represents API key to be used for authentication
	// with GLAIR Vision API
	ApiKey string

	// BaseUrl represents base URL path for GLAIR Vision API
	// endpoints. Defaults to "https://api.vision.glair.ai"
	BaseUrl string
	// ApiVersion represents GLAIR Vision API version to
	// be called. Defaults to "v1"
	ApiVersion string

	// Client represents the HTTP client that is used
	// to call the HTTP endpoints of GLAIR Vision API.
	// Defaults to the default HTTP client of Go
	Client HTTPClient
}

// NewConfig creates a new configuration object with default values for
// base URL and API Version.
func NewConfig(username string, password string, apiKey string) *Config {
	defaultClient := http.DefaultClient

	return &Config{
		Username:   username,
		Password:   password,
		ApiKey:     apiKey,
		BaseUrl:    defaultUrl,
		ApiVersion: defaultVersion,
		Client:     defaultClient,
	}
}

// GetEndpointURL creates service URL with base URL and API version
func (c *Config) GetEndpointURL(
	service string,
	endpoint string,
) string {
	parts := []string{c.BaseUrl, service, c.ApiVersion, endpoint}
	return strings.Join(parts, "/")
}

// WithCredentials sets user credentials for the configuration object
func (c *Config) WithCredentials(
	username string,
	password string,
	apiKey string,
) *Config {
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
