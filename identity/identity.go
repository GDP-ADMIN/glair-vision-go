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

func parseDateOfBirth(dateOfBirth interface{}) *string {
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
	url := identity.config.GetEndpointURL("identity", "verification")

	params := internal.RequestParameters{
		Url:       url,
		RequestID: input.RequestID,
		Body: map[string]interface{}{
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

	return internal.MakeMultipartRequest[BasicIdentityVerificationResult](
		ctx,
		params,
		identity.config,
	)
}

// FaceVerification performs KTP data verification
// to Dukcapil database from the given NIK
//
// API Docs: https://docs.glair.ai/vision/identity-face-verification
func (identity *Identity) FaceVerification(
	ctx context.Context,
	input glair.FaceVerificationInput,
) (FaceIdentityVerificationResult, error) {
	face, err := internal.ReadFile(input.FaceImage)
	if err != nil {
		return FaceIdentityVerificationResult{}, err
	}

	url := identity.config.GetEndpointURL("identity", "face-verification")

	params := internal.RequestParameters{
		Url:       url,
		RequestID: input.RequestID,
		Body: map[string]interface{}{
			"nik":           input.Nik,
			"name":          input.Name,
			"date_of_birth": parseDateOfBirth(input.DateOfBirth),
			"face_image":    face,
		},
	}

	return internal.MakeMultipartRequest[FaceIdentityVerificationResult](
		ctx,
		params,
		identity.config,
	)
}
