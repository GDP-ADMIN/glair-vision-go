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
func (f *FaceBio) FaceMatching(
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

	url := f.config.GetEndpointURL("face", "match")
	params := internal.RequestParameters{
		Url:       url,
		RequestID: input.RequestID,
		Payload: map[string]interface{}{
			"stored_image":   storedImage,
			"captured_image": capturedImage,
		},
	}

	return internal.MakeRequest[FaceMatching](ctx, params, f.config)
}

// PassiveLiveness performs liveness detection in passive environment
//
// API Docs: https://docs.glair.ai/vision/passive-liveness
func (f *FaceBio) PassiveLiveness(
	ctx context.Context,
	input glair.PassiveLivenessInput,
) (PassiveLiveness, error) {
	image, err := internal.ReadFile(input.Image)
	if err != nil {
		return PassiveLiveness{}, err
	}

	url := f.config.GetEndpointURL("face", "passive-liveness")
	params := internal.RequestParameters{
		Url:       url,
		RequestID: input.RequestID,
		Payload: map[string]interface{}{
			"image": image,
		},
	}

	return internal.MakeRequest[PassiveLiveness](ctx, params, f.config)
}

// ActiveLiveness performs liveness detection using predefined
// gestures and poses
//
// API Docs: https://docs.glair.ai/vision/active-liveness
func (f *FaceBio) ActiveLiveness(
	ctx context.Context,
	input glair.ActiveLivenessInput,
) (ActiveLiveness, error) {
	image, err := internal.ReadFile(input.Image)
	if err != nil {
		return ActiveLiveness{}, err
	}

	if input.GestureCode == "" {
		return ActiveLiveness{}, &glair.Error{
			Code:    glair.ErrorCodeInvalidArgs,
			Message: "Gesture code is required.",
		}
	}

	url := f.config.GetEndpointURL("face", "active-liveness")
	params := internal.RequestParameters{
		Url:       url,
		RequestID: input.RequestID,
		Payload: map[string]interface{}{
			"image":        image,
			"gesture-code": input.GestureCode,
		},
	}

	return internal.MakeRequest[ActiveLiveness](ctx, params, f.config)
}
