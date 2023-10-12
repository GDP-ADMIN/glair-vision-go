package config

import (
	"encoding/base64"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	username := "username"
	password := "password"
	apiKey := "api-key"

	url, _ := url.Parse("https://api.vision.glair.ai")

	config := New(username, password, apiKey)

	assert.Equal(t, username, config.Username)
	assert.Equal(t, password, config.Password)
	assert.Equal(t, apiKey, config.ApiKey)

	assert.Equal(t, url, config.BaseUrl)
	assert.Equal(t, "v1", config.ApiVersion)
}

func TestConfig_GetAuth(t *testing.T) {
	config := New("a", "b", "c")
	authHeader := config.GetAuth()

	expected := base64.RawStdEncoding.EncodeToString([]byte("a:b"))

	assert.Equal(t, "Basic "+expected, authHeader)
}

func TestConfig_WithCredentials(t *testing.T) {
	config := New("a", "b", "c").WithCredentials("d", "e", "f")

	assert.Equal(t, "d", config.Username)
	assert.Equal(t, "e", config.Password)
	assert.Equal(t, "f", config.ApiKey)
}

func TestConfig_WithBaseURL(t *testing.T) {
	url, _ := url.Parse("https://api.vision.glare.ai")

	config := New("a", "b", "c").WithBaseURL(url)

	assert.Equal(t, url, config.BaseUrl)
}

func TestConfig_WithVersion(t *testing.T) {
	config := New("a", "b", "c").WithVersion("v2")

	assert.Equal(t, "v2", config.ApiVersion)
}

func TestConfig_GetEndpointURL(t *testing.T) {
	config := New("a", "b", "c")
	url := config.GetEndpointURL("ocr", "ktp")

	expected, _ := url.Parse("https://api.vision.glair.ai/ocr/v1/ktp")

	assert.Equal(t, expected.String(), url.String())
}
