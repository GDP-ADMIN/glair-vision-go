package identity

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/glair-ai/glair-vision-go"
	"github.com/stretchr/testify/assert"
)

func successServer() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"success"}`))
	}))
}

func newConfig(url string) *glair.Config {
	return glair.NewConfig("u", "p", "k").WithBaseURL(url)
}

func TestNew(t *testing.T) {
	config := glair.NewConfig("", "", "")
	id := New(config)

	assert.Equal(t, config, id.config)
}

func TestBasicVerification(t *testing.T) {
	srv := successServer()
	defer srv.Close()

	id := New(newConfig(srv.URL))

	_, err := id.BasicVerification(context.Background(), glair.BasicVerificationInput{
		Nik: "1234567890123456",
	})
	assert.NoError(t, err)
}

func TestBasicVerificationRaw(t *testing.T) {
	srv := successServer()
	defer srv.Close()

	id := New(newConfig(srv.URL))

	_, err := id.BasicVerificationRaw(context.Background(), glair.BasicVerificationInput{
		Nik: "1234567890123456",
	})
	assert.NoError(t, err)
}

func TestFaceVerification(t *testing.T) {
	srv := successServer()
	defer srv.Close()

	id := New(newConfig(srv.URL))

	// face image read error
	_, err := id.FaceVerification(context.Background(), glair.FaceVerificationInput{
		Nik:       "1234567890123456",
		FaceImage: "does-not-exist.jpg",
	})
	assert.Error(t, err)

	// success
	_, err = id.FaceVerification(context.Background(), glair.FaceVerificationInput{
		Nik:       "1234567890123456",
		FaceImage: "../examples/ocr/images/ktp.jpeg",
	})
	assert.NoError(t, err)
}

func TestFaceVerificationRaw(t *testing.T) {
	srv := successServer()
	defer srv.Close()

	id := New(newConfig(srv.URL))

	_, err := id.FaceVerificationRaw(context.Background(), glair.FaceVerificationInput{
		Nik:       "1234567890123456",
		FaceImage: "../examples/ocr/images/ktp.jpeg",
	})
	assert.NoError(t, err)
}

func TestParseDateOfBirth(t *testing.T) {
	strDob := "01-01-1990"
	ts := time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC)

	tests := []struct {
		name  string
		input interface{}
		want  *string
	}{
		{
			name:  "string input",
			input: "01-01-1990",
			want:  glair.String("01-01-1990"),
		},
		{
			name:  "*string input",
			input: &strDob,
			want:  &strDob,
		},
		{
			name:  "time.Time input",
			input: ts,
			want:  glair.String("01-01-1990"),
		},
		{
			name:  "nil/default returns nil",
			input: nil,
			want:  nil,
		},
		{
			name:  "int (unknown type) returns nil",
			input: 12345,
			want:  nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := parseDateOfBirth(tc.input)
			assert.Equal(t, tc.want, result)
		})
	}
}
