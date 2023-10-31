package internal

import (
	"context"
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
		wantErr    *glair.Error
	}{
		{
			name: "failed to send request due to bad url",
			config: glair.NewConfig("username", "password", "api-key").
				WithBaseURL("this-is-invalid"),
			want: mockStruct{},
			wantErr: &glair.Error{
				Code:    glair.ErrorCodeInvalidURL,
				Message: "invalid base url provided in configuration",
			},
		},
		{
			name: "failed to send request due to bad client",
			config: glair.NewConfig("username", "password", "api-key").
				WithClient(failingClient{}),
			want: mockStruct{},
			wantErr: &glair.Error{
				Code:    glair.ErrorCodeBadClient,
				Message: "bad http client provided in configuration",
			},
		},
		{
			name:   "response is not OK",
			config: glair.NewConfig("username", "password", "api-key"),
			mockServer: httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(`{"status": "NO_FILE", "reason": "No file in request body"}`))
			})),
			want: mockStruct{},
			wantErr: &glair.Error{
				Code:    "NO_FILE",
				Message: "No file in request body",
				Body: glair.ResponseBody{
					Status: "NO_FILE",
					Reason: "No file in request body",
				},
			},
		},
		{
			name:   "response is not valid JSON",
			config: glair.NewConfig("username", "password", "api-key"),
			mockServer: httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(`im not a valid JSON`))
			})),
			want: mockStruct{},
			wantErr: &glair.Error{
				Code:    glair.ErrorCodeInvalidResponse,
				Message: "failed to parse http response",
			},
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
			res, err := MakeRequest[mockStruct](
				context.TODO(),
				tc.mockServer.URL,
				tc.config,
				file,
			)

			assert.Equal(t, tc.want, res)
			assert.Equal(t, tc.wantErr, err)
		})
	}
}
