// Package face is a collection of functions and objects that interacts
// with GLAIR Vision Face Biometrics API and its results
package face

import (
	"context"

	"github.com/glair-ai/glair-vision-go"
	"github.com/glair-ai/glair-vision-go/internal"
)

// FaceBIO provides functions to interact with GLAIR Vision
// Face Biometric products
type FaceBio struct {
	config *glair.Config
}

// New creates a GLAIR Vision Face Biometric API Client with
// the provided config
func New(config *glair.Config) *FaceBio {
	return &FaceBio{
		config: config,
	}
}

// FaceMatching performs face matching between stored
// and captured image
//
// API Docs: https://docs.glair.ai/vision/face-matching
func (face *FaceBio) FaceMatching(
	ctx context.Context,
	input glair.FaceMatchingInput,
) (FaceMatching, error) {
	storedImage, err := internal.ReadFile(input.StoredImage)
	if err != nil {
		return FaceMatching{}, err
	}

	capturedImage, err := internal.ReadFile(input.CapturedImage)
	if err != nil {
		return FaceMatching{}, err
	}

	url := face.config.GetEndpointURL("face", "match")
	params := internal.RequestParameters{
		Url:       url,
		RequestID: input.RequestID,
		Body: map[string]interface{}{
			"stored_image":   storedImage,
			"captured_image": capturedImage,
		},
	}

	return internal.MakeMultipartRequest[FaceMatching](ctx, params, face.config)
}

// PassiveLiveness performs liveness detection in passive environment
//
// API Docs: https://docs.glair.ai/vision/passive-liveness
func (face *FaceBio) PassiveLiveness(
	ctx context.Context,
	input glair.PassiveLivenessInput,
) (PassiveLiveness, error) {
	image, err := internal.ReadFile(input.Image)
	if err != nil {
		return PassiveLiveness{}, err
	}

	url := face.config.GetEndpointURL("face", "passive-liveness")
	params := internal.RequestParameters{
		Url:       url,
		RequestID: input.RequestID,
		Body: map[string]interface{}{
			"image": image,
		},
	}

	return internal.MakeMultipartRequest[PassiveLiveness](ctx, params, face.config)
}

// ActiveLiveness performs liveness detection using predefined
// gestures and poses
//
// API Docs: https://docs.glair.ai/vision/active-liveness
func (face *FaceBio) ActiveLiveness(
	ctx context.Context,
	input glair.ActiveLivenessInput,
) (ActiveLiveness, error) {
	image, err := internal.ReadFile(input.Image)
	if err != nil {
		return ActiveLiveness{}, err
	}

	url := face.config.GetEndpointURL("face", "active-liveness")
	params := internal.RequestParameters{
		Url:       url,
		RequestID: input.RequestID,
		Body: map[string]interface{}{
			"image":        image,
			"gesture-code": input.GestureCode,
		},
	}

	return internal.MakeMultipartRequest[ActiveLiveness](ctx, params, face.config)
}

// PassiveLivenessSessions sends session request for passive liveness
// using the prebuilt web page
//
// API Docs: https://docs.glair.ai/vision/passive-liveness-sessions
func (face *FaceBio) PassiveLivenessSessions(
	ctx context.Context,
	input glair.SessionsInput,
) (glair.Session, error) {
	payload := map[string]interface{}{
		"success_url": input.SuccessURL,
	}

	if input.CancelURL != nil {
		payload["cancel_url"] = input.CancelURL
	}
	url := face.config.GetEndpointURL("face", "passive-liveness-sessions")
	params := internal.RequestParameters{
		Url:  url,
		Body: payload,
	}

	return internal.MakeJSONRequest[glair.Session](ctx, params, face.config)
}

// ActiveLivenessSessions sends session request for passive liveness
// using the prebuilt web page
//
// API Docs: https://docs.glair.ai/vision/active-liveness-sessions
func (face *FaceBio) ActiveLivenessSessions(
	ctx context.Context,
	input glair.SessionsInput,
) (glair.Session, error) {
	payload := map[string]interface{}{
		"success_url": input.SuccessURL,
	}

	if input.CancelURL != nil {
		payload["cancel_url"] = input.CancelURL
	}
	url := face.config.GetEndpointURL("face", "active-liveness-sessions")
	params := internal.RequestParameters{
		Url:  url,
		Body: payload,
	}

	return internal.MakeJSONRequest[glair.Session](ctx, params, face.config)
}
