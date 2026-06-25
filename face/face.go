package face

import (
	"context"

	"github.com/glair-ai/glair-vision-go"
	"github.com/glair-ai/glair-vision-go/internal"
)

type FaceBio struct {
	config *glair.Config
}

func New(config *glair.Config) *FaceBio {
	return &FaceBio{
		config: config,
	}
}

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

	params := internal.RequestParameters{
		Url:       face.config.GetEndpointURL("face", "match"),
		RequestID: input.RequestID,
		Body: map[string]any{
			"stored_image":   storedImage,
			"captured_image": capturedImage,
		},
	}

	return internal.MakeMultipartRequest[FaceMatching](ctx, params, face.config)
}

func (face *FaceBio) FaceMatchingRaw(
	ctx context.Context,
	input glair.FaceMatchingInput,
) ([]byte, error) {
	storedImage, err := internal.ReadFile(input.StoredImage)
	if err != nil {
		return nil, err
	}

	capturedImage, err := internal.ReadFile(input.CapturedImage)
	if err != nil {
		return nil, err
	}

	params := internal.RequestParameters{
		Url:       face.config.GetEndpointURL("face", "match"),
		RequestID: input.RequestID,
		Body: map[string]any{
			"stored_image":   storedImage,
			"captured_image": capturedImage,
		},
	}

	return internal.MakeMultipartRequestRaw(ctx, params, face.config)
}

func (face *FaceBio) PassiveLiveness(
	ctx context.Context,
	input glair.PassiveLivenessInput,
) (PassiveLiveness, error) {
	image, err := internal.ReadFile(input.Image)
	if err != nil {
		return PassiveLiveness{}, err
	}

	params := internal.RequestParameters{
		Url:       face.config.GetEndpointURL("face", "passive-liveness"),
		RequestID: input.RequestID,
		Body: map[string]any{
			"image": image,
		},
	}

	return internal.MakeMultipartRequest[PassiveLiveness](ctx, params, face.config)
}

func (face *FaceBio) PassiveLivenessRaw(
	ctx context.Context,
	input glair.PassiveLivenessInput,
) ([]byte, error) {
	image, err := internal.ReadFile(input.Image)
	if err != nil {
		return nil, err
	}

	params := internal.RequestParameters{
		Url:       face.config.GetEndpointURL("face", "passive-liveness"),
		RequestID: input.RequestID,
		Body: map[string]any{
			"image": image,
		},
	}

	return internal.MakeMultipartRequestRaw(ctx, params, face.config)
}

func (face *FaceBio) ActiveLiveness(
	ctx context.Context,
	input glair.ActiveLivenessInput,
) (ActiveLiveness, error) {
	image, err := internal.ReadFile(input.Image)
	if err != nil {
		return ActiveLiveness{}, err
	}

	params := internal.RequestParameters{
		Url:       face.config.GetEndpointURL("face", "active-liveness"),
		RequestID: input.RequestID,
		Body: map[string]any{
			"image":        image,
			"gesture-code": input.GestureCode,
		},
	}

	return internal.MakeMultipartRequest[ActiveLiveness](ctx, params, face.config)
}

func (face *FaceBio) ActiveLivenessRaw(
	ctx context.Context,
	input glair.ActiveLivenessInput,
) ([]byte, error) {
	image, err := internal.ReadFile(input.Image)
	if err != nil {
		return nil, err
	}

	params := internal.RequestParameters{
		Url:       face.config.GetEndpointURL("face", "active-liveness"),
		RequestID: input.RequestID,
		Body: map[string]any{
			"image":        image,
			"gesture-code": input.GestureCode,
		},
	}

	return internal.MakeMultipartRequestRaw(ctx, params, face.config)
}

func (face *FaceBio) PassiveLivenessSessions(
	ctx context.Context,
	input glair.SessionsInput,
) (glair.Session, error) {
	payload := map[string]any{
		"success_url": input.SuccessURL,
	}

	if input.CancelURL != nil {
		payload["cancel_url"] = input.CancelURL
	}

	params := internal.RequestParameters{
		Url:  face.config.GetEndpointURL("face", "passive-liveness-sessions"),
		Body: payload,
	}

	return internal.MakeJSONRequest[glair.Session](ctx, params, face.config)
}

func (face *FaceBio) ActiveLivenessSessions(
	ctx context.Context,
	input glair.SessionsInput,
) (glair.Session, error) {
	payload := map[string]any{
		"success_url": input.SuccessURL,
	}

	if input.CancelURL != nil {
		payload["cancel_url"] = input.CancelURL
	}

	params := internal.RequestParameters{
		Url:  face.config.GetEndpointURL("face", "active-liveness-sessions"),
		Body: payload,
	}

	return internal.MakeJSONRequest[glair.Session](ctx, params, face.config)
}
