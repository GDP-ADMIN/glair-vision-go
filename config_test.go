package glair

import (
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewConfig(t *testing.T) {
	username := "username"
	password := "password"
	apiKey := "api-key"

	config := NewConfig(username, password, apiKey)

	assert.Equal(t, username, config.Username)
	assert.Equal(t, password, config.Password)
	assert.Equal(t, apiKey, config.ApiKey)

	assert.Equal(t, "https://api.vision.glair.ai", config.BaseUrl)
	assert.Equal(t, "v1", config.ApiVersion)
}

func TestConfig_WithCredentials(t *testing.T) {
	config := NewConfig("a", "b", "c").WithCredentials("d", "e", "f")

	assert.Equal(t, "d", config.Username)
	assert.Equal(t, "e", config.Password)
	assert.Equal(t, "f", config.ApiKey)
}

func TestConfig_WithBaseURL(t *testing.T) {
	url := "https://api.vision.glare.ai"

	config := NewConfig("a", "b", "c").WithBaseURL(url)

	assert.Equal(t, url, config.BaseUrl)
}

func TestConfig_WithVersion(t *testing.T) {
	config := NewConfig("a", "b", "c").WithVersion("v2")

	assert.Equal(t, "v2", config.ApiVersion)
}

type fakeClient struct{}

func (c fakeClient) Do(_ *http.Request) (*http.Response, error) {
	return nil, errors.New("failed to send request")
}

func TestConfig_WithClient(t *testing.T) {
	httpClient := fakeClient{}
	config := NewConfig("a", "b", "c").WithClient(httpClient)

	assert.Equal(t, httpClient, config.Client)
}

func TestConfig_GetEndpointURL(t *testing.T) {
	url := NewConfig("a", "b", "c").GetEndpointURL("ocr", "ktp")

	assert.Equal(t, "https://api.vision.glair.ai/ocr/v1/ktp", url)
}

type SampleLogger struct{}

func (l SampleLogger) Debugf(format string, val ...interface{}) {}
func (l SampleLogger) Infof(format string, val ...interface{})  {}
func (l SampleLogger) Warnf(format string, val ...interface{})  {}
func (l SampleLogger) Errorf(format string, val ...interface{}) {}

func TestConfig_WithLogger(t *testing.T) {
	logger := SampleLogger{}

	config := NewConfig("a", "b", "c").WithLogger(logger)

	assert.Equal(t, logger, config.Logger)
}
