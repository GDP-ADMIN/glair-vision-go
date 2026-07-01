// Package identity is a collection of functions and objects that interacts
// with GLAIR Vision identity verification API and its results
package identity

import (
	"context"
	"time"

	"github.com/glair-ai/glair-vision-go"
	"github.com/glair-ai/glair-vision-go/internal"
)

// Identity provides functions to interact with
// GLAIR Vision identity verification API
type Identity struct {
	config *glair.Config
}

// IdentityVerificationResult is a wrapper object for identity
// verification result
type IdentityVerificationResult[T any] struct {
	VerificationStatus string `json:"verification_result,omitempty"`
	Reason             string `json:"reason,omitempty"`
	Result             T      `json:"result"`
}

func parseDateOfBirth(dateOfBirth any) *string {
	switch dob := dateOfBirth.(type) {
	case string:
		return glair.String(dob)
	case *string:
		return dob
	case time.Time:
		return glair.String(dob.Format("02-01-2006"))
	default:
		return nil
	}
}

func basicVerificationParams(identity *Identity, input glair.BasicVerificationInput) internal.RequestParameters {
	url := identity.config.GetEndpointURL("identity", "verification")
	return internal.RequestParameters{
		Url:       url,
		RequestID: input.RequestID,
		Body: map[string]any{
			"nik":                input.Nik,
			"name":               input.Name,
			"date_of_birth":      parseDateOfBirth(input.DateOfBirth),
			"no_kk":              input.NoKk,
			"mother_maiden_name": input.MotherMaidenName,
			"place_of_birth":     input.PlaceOfBirth,
			"address":            input.Address,
			"gender":             input.Gender,
			"marital_status":     input.MaritalStatus,
			"job_type":           input.JobType,
			"province":           input.Province,
			"city":               input.City,
			"district":           input.District,
			"subdistrict":        input.Subdistrict,
			"rt":                 input.Rt,
			"rw":                 input.Rw,
		},
	}
}

func faceVerificationParams(identity *Identity, input glair.FaceVerificationInput) (internal.RequestParameters, error) {
	face, err := internal.ReadFile(input.FaceImage)
	if err != nil {
		return internal.RequestParameters{}, err
	}

	url := identity.config.GetEndpointURL("identity", "face-verification")
	return internal.RequestParameters{
		Url:       url,
		RequestID: input.RequestID,
		Body: map[string]any{
			"nik":           input.Nik,
			"name":          input.Name,
			"date_of_birth": parseDateOfBirth(input.DateOfBirth),
			"face_image":    face,
		},
	}, nil
}

// New creates a GLAIR Vision identity verification API client
// with the given config
func New(config *glair.Config) *Identity {
	return &Identity{
		config: config,
	}
}

// BasicVerification performs KTP data verification
// to Dukcapil database from the given NIK
//
// API Docs: https://docs.glair.ai/vision/identity-verification
func (identity *Identity) BasicVerification(
	ctx context.Context,
	input glair.BasicVerificationInput,
) (BasicIdentityVerificationResult, error) {
	params := basicVerificationParams(identity, input)

	return internal.MakeMultipartRequest[BasicIdentityVerificationResult](ctx, params, identity.config)
}

// BasicVerificationRaw performs KTP data verification
// to Dukcapil database and returns the raw API response as bytes
//
// WARNING: The raw response contains unredacted PII (NIK, name, date of birth, etc.).
// Do not log or persist the raw bytes in plaintext.
//
// API Docs: https://docs.glair.ai/vision/identity-verification
func (identity *Identity) BasicVerificationRaw(
	ctx context.Context,
	input glair.BasicVerificationInput,
) ([]byte, error) {
	params := basicVerificationParams(identity, input)

	return internal.MakeMultipartRequestRaw(ctx, params, identity.config)
}

// FaceVerification performs face image verification
// against the Dukcapil database from the given NIK
//
// API Docs: https://docs.glair.ai/vision/identity-face-verification
func (identity *Identity) FaceVerification(
	ctx context.Context,
	input glair.FaceVerificationInput,
) (FaceIdentityVerificationResult, error) {
	params, err := faceVerificationParams(identity, input)
	if err != nil {
		return FaceIdentityVerificationResult{}, err
	}

	return internal.MakeMultipartRequest[FaceIdentityVerificationResult](ctx, params, identity.config)
}

// FaceVerificationRaw performs face image verification
// against the Dukcapil database and returns the raw API response as bytes
//
// WARNING: The raw response contains unredacted PII (NIK, name, date of birth, face image, etc.).
// Do not log or persist the raw bytes in plaintext.
//
// API Docs: https://docs.glair.ai/vision/identity-face-verification
func (identity *Identity) FaceVerificationRaw(
	ctx context.Context,
	input glair.FaceVerificationInput,
) ([]byte, error) {
	params, err := faceVerificationParams(identity, input)
	if err != nil {
		return nil, err
	}

	return internal.MakeMultipartRequestRaw(ctx, params, identity.config)
}
