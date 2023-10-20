package internal

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

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

func TestMakeRequest(t *testing.T) {
	file, _ := os.Open("../examples/images/ktp.jpeg")

	tests := []struct {
		name       string
		config     *glair.Config
		mockServer *httptest.Server
		want       mockStruct
		wantErr    error
	}{
		{
			name: "failed to send request",
			config: glair.NewConfig("username", "password", "api-key").
				WithClient(failingClient{}),
			want:    mockStruct{},
			wantErr: glair.ErrBadClient,
		},
		{
			name:   "response is not OK",
			config: glair.NewConfig("username", "password", "api-key"),
			mockServer: httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(`{"status": "NO_FILE", "reason": "No file in request body"}`))
			})),
			want: mockStruct{},
			wantErr: glair.RequestError{
				StatusCode:   http.StatusBadRequest,
				ResponseBody: []byte(`{"status": "NO_FILE", "reason": "No file in request body"}`),
			},
		},
		{
			name:   "response is not valid JSON",
			config: glair.NewConfig("username", "password", "api-key"),
			mockServer: httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(`im not a valid JSON`))
			})),
			want:    mockStruct{},
			wantErr: glair.ErrInvalidResponseBody,
		},
		{
			name:   "success",
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
			cfg := tc.config
			if tc.mockServer != nil {
				cfg = tc.config.WithBaseURL(tc.mockServer.URL)
			}

			url, _ := cfg.GetEndpointURL("ocr", "ktp")

			res, err := MakeRequest[mockStruct](cfg, url, file)

			assert.Equal(t, tc.want, res)
			assert.Equal(t, tc.wantErr, err)
		})
	}
}
