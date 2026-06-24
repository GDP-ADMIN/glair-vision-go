package internal

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/glair-ai/glair-vision-go"
	"github.com/stretchr/testify/assert"
)

type mockStruct struct {
	Name string `json:"name,omitempty"`
}

type failingClient struct{}

func (c failingClient) Do(req *http.Request) (*http.Response, error) {
	return nil, errors.New("failed to send request")
}

// failingBodyReader is a Reader that always returns an error after the first read.
type failingBodyReader struct{}

func (r failingBodyReader) Read(p []byte) (int, error) {
	return 0, errors.New("simulated body read error")
}

func (r failingBodyReader) Close() error {
	return nil
}

// bodyFailClient returns a 200 response with a body reader that fails immediately.
type bodyFailClient struct{}

func (c bodyFailClient) Do(req *http.Request) (*http.Response, error) {
	return &http.Response{
		StatusCode: http.StatusOK,
		Body:       failingBodyReader{},
	}, nil
}

func TestMakeMultipartRequest(t *testing.T) {
	file, _ := os.Open("../examples/ocr/images/ktp.jpeg")

	baseCtx := context.Background()
	// set unrealistic timeout
	timeoutCtx, cancel := context.WithTimeout(baseCtx, 0*time.Nanosecond)
	defer cancel()

	tests := []struct {
		name       string
		ctx        context.Context
		config     *glair.Config
		mockServer *httptest.Server
		want       mockStruct
		wantErr    *glair.Error
	}{
		{
			name: "failed to send request due to bad url",
			ctx:  context.Background(),
			config: glair.NewConfig("username", "password", "api-key").
				WithBaseURL("%+0"),
			want: mockStruct{},
			wantErr: &glair.Error{
				Code:    glair.ErrorCodeNetworkError,
				Message: "Invalid base URL %+0 is provided in configuration.",
			},
		},
		{
			name: "failed to send request due to network error",
			ctx:  context.Background(),
			config: glair.NewConfig("username", "password", "api-key").
				WithClient(failingClient{}),
			want: mockStruct{},
			wantErr: &glair.Error{
				Code:    glair.ErrorCodeNetworkError,
				Message: "Failed to send HTTP request due to network error.",
			},
		},
		{
			name:   "failed to send request due to timeout",
			ctx:    timeoutCtx,
			config: glair.NewConfig("username", "password", "api-key"),
			want:   mockStruct{},
			wantErr: &glair.Error{
				Code:    glair.ErrorCodeTimeout,
				Message: "Request to https://api.vision.glair.ai timed out",
			},
		},
		{
			name:   "response is not OK, forbidden",
			ctx:    context.Background(),
			config: glair.NewConfig("username", "password", "api-key"),
			mockServer: httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusForbidden)
				w.Write([]byte(`{"message": "Access to this API has been disallowed."}`))
			})),
			want: mockStruct{},
			wantErr: &glair.Error{
				Code:    glair.ErrorCodeForbidden,
				Message: "Insufficient access to the API endpoint",
			},
		},
		{
			name:   "response is not OK, handled error, OCR",
			ctx:    context.Background(),
			config: glair.NewConfig("username", "password", "api-key"),
			mockServer: httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(`{"status": "NO_FILE", "reason": "No file in request body"}`))
			})),
			want: mockStruct{},
			wantErr: &glair.Error{
				Code:    glair.ErrorCodeAPIError,
				Message: "GLAIR API returned non-OK response. Please check the Response property for more detailed explanation.",
				Response: glair.Response{
					Status: 400,
					Body: map[string]interface{}{
						"status": "NO_FILE",
						"reason": "No file in request body",
					},
				},
			},
		},
		{
			name:   "response is not OK, gateway and miscellanous errors",
			ctx:    context.Background(),
			config: glair.NewConfig("username", "password", "api-key"),
			mockServer: httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusBadGateway)
				w.Write([]byte(`<html><body><p>502 - Bad Gateway</p></body></html>`))
			})),
			want: mockStruct{},
			wantErr: &glair.Error{
				Code:    glair.ErrorCodeInvalidResponse,
				Message: "Failed to parse API response. Please contact us about this error.",
				Response: glair.Response{
					Status: 502,
				},
			},
		},
		{
			name:   "success",
			ctx:    context.Background(),
			config: glair.NewConfig("username", "password", "api-key"),
			mockServer: httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(`{"name":"foo"}`))
			})),
			want: mockStruct{
				Name: "foo",
			},
			wantErr: nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			url := tc.config.BaseUrl
			if tc.mockServer != nil {
				url = tc.mockServer.URL
			}

			params := RequestParameters{
				Url:       url,
				RequestID: "samples",
				Body: map[string]interface{}{
					"image":          file,
					"category":       "do-not-panic",
					"pointer-string": glair.String("do-not-panic"),
					"page":           glair.Int(3),
				},
			}

			var res mockStruct
			err := MakeMultipartRequest(
				tc.ctx,
				params,
				tc.config,
				&res,
			)

			assert.Equal(t, tc.want, res)

			if tc.wantErr == nil {
				assert.Equal(t, nil, err)
			} else {
				glairError := err.(*glair.Error)

				assert.Equal(t, tc.wantErr.Code, glairError.Code)
				assert.Equal(t, tc.wantErr.Message, glairError.Message)
				assert.Equal(t, tc.wantErr.Response, glairError.Response)
			}
		})
	}
}

func TestMakeJSONRequest(t *testing.T) {
	tests := []struct {
		name       string
		config     *glair.Config
		mockServer *httptest.Server
		want       mockStruct
		wantErr    *glair.Error
	}{
		{
			name: "failed to send request due to bad url",
			config: glair.NewConfig("username", "password", "api-key").
				WithBaseURL("%+0"),
			want: mockStruct{},
			wantErr: &glair.Error{
				Code:    glair.ErrorCodeNetworkError,
				Message: "Invalid base URL %+0 is provided in configuration.",
			},
		},
		{
			name:   "success",
			config: glair.NewConfig("username", "password", "api-key"),
			mockServer: httptest.NewServer(
				http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusOK)
					w.Write([]byte(`{"name":"foo"}`))
				}),
			),
			want: mockStruct{
				Name: "foo",
			},
			wantErr: nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			url := tc.config.BaseUrl
			if tc.mockServer != nil {
				url = tc.mockServer.URL
			}

			params := RequestParameters{
				Url:       url,
				RequestID: "samples",
				Body: map[string]interface{}{
					"success_url": "https://www.google.com",
					"cancel_url":  "https://www.bing.com",
				},
			}

			var res mockStruct
			err := MakeJSONRequest(
				context.TODO(),
				params,
				tc.config,
				&res,
			)

			assert.Equal(t, tc.want, res)
			if tc.wantErr == nil {
				assert.Equal(t, nil, err)
			} else {
				glairError := err.(*glair.Error)

				assert.Equal(t, tc.wantErr.Code, glairError.Code)
				assert.Equal(t, tc.wantErr.Message, glairError.Message)
				assert.Equal(t, tc.wantErr.Response, glairError.Response)
			}
		})
	}
}

func TestSendRequest_NilContext(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"name":"foo"}`))
	}))
	defer srv.Close()

	config := glair.NewConfig("u", "p", "k").WithBaseURL(srv.URL)

	params := RequestParameters{
		Url:  srv.URL,
		Body: map[string]interface{}{},
	}

	// Intentionally passing nil context to exercise the nil-ctx branch (line 89-91)
	_, _, err := sendRequest(nil, RequestPayload{
		RequestParameters: params,
		Header:            map[string]string{"Content-Type": "application/json"},
		Body:              strings.NewReader(`{}`),
	}, config)

	assert.NoError(t, err)
}

func TestSendRequest_BodyReadError(t *testing.T) {
	config := glair.NewConfig("u", "p", "k").
		WithBaseURL("https://api.vision.glair.ai").
		WithClient(bodyFailClient{})

	params := RequestParameters{
		Url:  "https://api.vision.glair.ai",
		Body: map[string]interface{}{},
	}

	_, _, err := sendRequest(context.Background(), RequestPayload{
		RequestParameters: params,
		Header:            map[string]string{"Content-Type": "application/json"},
		Body:              strings.NewReader(`{}`),
	}, config)

	assert.Error(t, err)
	glairErr, ok := err.(*glair.Error)
	assert.True(t, ok)
	assert.Equal(t, glair.ErrorCodeInvalidResponse, glairErr.Code)
}

func TestCreateMultipartPayload_ClosedFileCopyError(t *testing.T) {
	// Open and immediately close a file so io.Copy will fail when reading it.
	tmpFile, err := os.CreateTemp("", "test-*.jpg")
	assert.NoError(t, err)
	tmpFile.WriteString("dummy content")
	tmpFile.Close() // close so subsequent reads fail

	logger := &glair.LeveledLogger{Level: glair.LevelDebug}

	_, _, err = createMultipartPayload(map[string]interface{}{
		"image": tmpFile,
	}, logger)

	assert.Error(t, err)
	glairErr, ok := err.(*glair.Error)
	assert.True(t, ok)
	assert.Equal(t, glair.ErrorCodeFileError, glairErr.Code)

	os.Remove(tmpFile.Name())
}

func TestMakeMultipartRequest_CreateMultipartError(t *testing.T) {
	// Open and immediately close a file so createMultipartPayload fails.
	tmpFile, err := os.CreateTemp("", "test-*.jpg")
	assert.NoError(t, err)
	tmpFile.WriteString("dummy content")
	tmpFile.Close()

	config := glair.NewConfig("u", "p", "k")

	params := RequestParameters{
		Url: config.GetEndpointURL("ocr", "ktp"),
		Body: map[string]interface{}{
			"image": tmpFile,
		},
	}

	_, err = MakeMultipartRequest[mockStruct](context.Background(), params, config)
	assert.Error(t, err)
	glairErr, ok := err.(*glair.Error)
	assert.True(t, ok)
	assert.Equal(t, glair.ErrorCodeFileError, glairErr.Code)

	os.Remove(tmpFile.Name())
}
