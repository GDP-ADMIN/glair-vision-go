// Package config is collection of objects that provides configuration
// for GLAIR Vision SDK client
package config

import (
	"encoding/base64"
	"net/url"
)

// Credentials provides user credentials
type Credentials struct {
	username string
	password string
	apiKey   string
}

// Config provides configuration for a client instance,
// including credentials and API configuration such as
// base URL and API version.
type Config struct {
	username string
	password string
	apiKey   string

	baseUrl    *url.URL
	apiVersion string
}

// NewConfig creates a new configuration object with default values for
// base URL and API Version.
func NewConfig(credentials Credentials) *Config {
	defaultUrl, _ := url.Parse("https://api.vision.glair.ai")
	defaultApiVersion := "v1"

	return &Config{
		username:   credentials.username,
		password:   credentials.password,
		apiKey:     credentials.apiKey,
		baseUrl:    defaultUrl,
		apiVersion: defaultApiVersion,
	}
}

// GetAuth creates authentication headers to interact with
// GLAIR Vision products
func (c *Config) GetAuth() string {
	basicCredentials := base64.StdEncoding.EncodeToString(
		[]byte(c.username + ":" + c.password),
	)

	return "Basic " + basicCredentials
}

// WithCredentials sets user credentials for a configuration instance
func (c *Config) WithCredentials(credentials Credentials) *Config {
	c.username = credentials.username
	c.password = credentials.password
	c.apiKey = credentials.apiKey

	return c
}

// WithBaseURL sets base URL for a configuration instance
func (c *Config) WithBaseURL(baseUrl *url.URL) *Config {
	c.baseUrl = baseUrl

	return c
}

// WithVersion
func (c *Config) WithVersion(version string) *Config {
	c.apiVersion = version

	return c
}
