package face

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

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
	f := New(config)

	assert.Equal(t, config, f.config)
}

func TestFaceMatching(t *testing.T) {
	srv := successServer()
	defer srv.Close()

	f := New(newConfig(srv.URL))

	// stored image read error
	err := f.FaceMatching(context.Background(), glair.FaceMatchingInput{
		StoredImage:   "does-not-exist.jpg",
		CapturedImage: "../examples/ocr/images/ktp.jpeg",
	}, nil)
	assert.Error(t, err)

	// captured image read error
	err = f.FaceMatching(context.Background(), glair.FaceMatchingInput{
		StoredImage:   "../examples/ocr/images/ktp.jpeg",
		CapturedImage: "does-not-exist.jpg",
	}, nil)
	assert.Error(t, err)

	// success
	err = f.FaceMatching(context.Background(), glair.FaceMatchingInput{
		StoredImage:   "../examples/ocr/images/ktp.jpeg",
		CapturedImage: "../examples/ocr/images/ktp.jpeg",
	}, nil)
	assert.NoError(t, err)
}

func TestPassiveLiveness(t *testing.T) {
	srv := successServer()
	defer srv.Close()

	f := New(newConfig(srv.URL))

	err := f.PassiveLiveness(context.Background(), glair.PassiveLivenessInput{
		Image: "does-not-exist.jpg",
	}, nil)
	assert.Error(t, err)

	err = f.PassiveLiveness(context.Background(), glair.PassiveLivenessInput{
		Image: "../examples/ocr/images/ktp.jpeg",
	}, nil)
	assert.NoError(t, err)
}

func TestActiveLiveness(t *testing.T) {
	srv := successServer()
	defer srv.Close()

	f := New(newConfig(srv.URL))

	err := f.ActiveLiveness(context.Background(), glair.ActiveLivenessInput{
		Image:       "does-not-exist.jpg",
		GestureCode: "HAND_00000",
	}, nil)
	assert.Error(t, err)

	err = f.ActiveLiveness(context.Background(), glair.ActiveLivenessInput{
		Image:       "../examples/ocr/images/ktp.jpeg",
		GestureCode: "HAND_00000",
	}, nil)
	assert.NoError(t, err)
}

func TestPassiveLivenessSessions(t *testing.T) {
	srv := successServer()
	defer srv.Close()

	f := New(newConfig(srv.URL))

	// without cancel URL
	err := f.PassiveLivenessSessions(context.Background(), glair.SessionsInput{
		SuccessURL: "https://example.com/success",
	}, nil)
	assert.NoError(t, err)

	// with cancel URL
	err = f.PassiveLivenessSessions(context.Background(), glair.SessionsInput{
		SuccessURL: "https://example.com/success",
		CancelURL:  glair.String("https://example.com/cancel"),
	}, nil)
	assert.NoError(t, err)
}

func TestActiveLivenessSessions(t *testing.T) {
	srv := successServer()
	defer srv.Close()

	f := New(newConfig(srv.URL))

	// without cancel URL
	err := f.ActiveLivenessSessions(context.Background(), glair.SessionsInput{
		SuccessURL: "https://example.com/success",
	}, nil)
	assert.NoError(t, err)

	// with cancel URL
	err = f.ActiveLivenessSessions(context.Background(), glair.SessionsInput{
		SuccessURL: "https://example.com/success",
		CancelURL:  glair.String("https://example.com/cancel"),
	}, nil)
	assert.NoError(t, err)
}
