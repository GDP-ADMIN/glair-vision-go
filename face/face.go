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
    target any,
) error {
    storedImage, err := internal.ReadFile(input.StoredImage)
    if err != nil {
        return err
    }

    capturedImage, err := internal.ReadFile(input.CapturedImage)
    if err != nil {
        return err
    }

    params := internal.RequestParameters{
        Url:       face.config.GetEndpointURL("face", "match"),
        RequestID: input.RequestID,
        Body: map[string]any{
            "stored_image":   storedImage,
            "captured_image": capturedImage,
        },
    }

    return internal.MakeMultipartRequest(
        ctx,
        params,
        face.config,
        target,
    )
}

// PassiveLiveness performs liveness detection in passive environment
//
// API Docs: https://docs.glair.ai/vision/passive-liveness
func (face *FaceBio) PassiveLiveness(
    ctx context.Context,
    input glair.PassiveLivenessInput,
    target any,
) error {
    image, err := internal.ReadFile(input.Image)
    if err != nil {
        return err
    }

    params := internal.RequestParameters{
        Url:       face.config.GetEndpointURL("face", "passive-liveness"),
        RequestID: input.RequestID,
        Body: map[string]any{
            "image": image,
        },
    }

    return internal.MakeMultipartRequest(
        ctx,
        params,
        face.config,
        target,
    )
}

// ActiveLiveness performs liveness detection using predefined
// gestures and poses
//
// API Docs: https://docs.glair.ai/vision/active-liveness
func (face *FaceBio) ActiveLiveness(
    ctx context.Context,
    input glair.ActiveLivenessInput,
    target any,
) error {
    image, err := internal.ReadFile(input.Image)
    if err != nil {
        return err
    }

    params := internal.RequestParameters{
        Url:       face.config.GetEndpointURL("face", "active-liveness"),
        RequestID: input.RequestID,
        Body: map[string]any{
            "image":        image,
            "gesture-code": input.GestureCode,
        },
    }

    return internal.MakeMultipartRequest(
        ctx,
        params,
        face.config,
        target,
    )
}

// PassiveLivenessSessions sends session request for passive liveness
// using the prebuilt web page
//
// API Docs: https://docs.glair.ai/vision/passive-liveness-sessions
func (face *FaceBio) PassiveLivenessSessions(
    ctx context.Context,
    input glair.SessionsInput,
    target any,
) error {
    payload := map[string]any{
        "success_url": input.SuccessURL,
    }

    if input.CancelURL != nil {
        payload["cancel_url"] = input.CancelURL
    }

    params := internal.RequestParameters{
        Url: face.config.GetEndpointURL("face", "passive-liveness-sessions"),
        Body: payload,
    }

    return internal.MakeJSONRequest(
        ctx,
        params,
        face.config,
        target,
    )
}

// ActiveLivenessSessions sends session request for active liveness
// using the prebuilt web page
//
// API Docs: https://docs.glair.ai/vision/active-liveness-sessions
func (face *FaceBio) ActiveLivenessSessions(
    ctx context.Context,
    input glair.SessionsInput,
    target any,
) error {
    payload := map[string]any{
        "success_url": input.SuccessURL,
    }

    if input.CancelURL != nil {
        payload["cancel_url"] = input.CancelURL
    }

    params := internal.RequestParameters{
        Url: face.config.GetEndpointURL("face", "active-liveness-sessions"),
        Body: payload,
    }

    return internal.MakeJSONRequest(
        ctx,
        params,
        face.config,
        target,
    )
}
