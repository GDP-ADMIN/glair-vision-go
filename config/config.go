// Package config is collection of objects that provides configuration
// for GLAIR Vision SDK client
package config

import (
	"encoding/base64"
	"net/http"
	"net/url"
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

	BaseUrl    *url.URL
	ApiVersion string

	Client HTTPClient
}

// New creates a new configuration object with default values for
// base URL and API Version.
func New(username string, password string, apiKey string) *Config {
	defaultUrl, _ := url.Parse("https://api.vision.glair.ai")
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

// GetAuth creates authentication headers to interact with
// GLAIR Vision products
func (c *Config) GetAuth() string {
	basicCredentials := base64.StdEncoding.EncodeToString(
		[]byte(c.Username + ":" + c.Password),
	)

	return "Basic " + basicCredentials
}

// GetEndpointURL creates service URL with base URL and API version
func (c *Config) GetEndpointURL(service string, endpoint string) *url.URL {
	return c.BaseUrl.JoinPath(service, c.ApiVersion, endpoint)
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
func (c *Config) WithBaseURL(baseUrl *url.URL) *Config {
	c.BaseUrl = baseUrl

	return c
}

// WithVersion sets API version for the configuration object
func (c *Config) WithVersion(version string) *Config {
	c.ApiVersion = version

	return c
}
