package identity

// FaceIdentityVerificationResult represents face
// identity verification result from GLAIR Vision API
// of the given face data and NIK
type FaceIdentityVerificationResult = IdentityVerificationResult[FaceIdentityVerification]

type FaceIdentityVerification struct {
	Nik                 bool    `json:"nik,omitempty"`
	Name                bool    `json:"name,omitempty"`
	DateOfBirth         bool    `json:"date_of_birth,omitempty"`
	FaceImagePercentage float64 `json:"face_image_percentage,omitempty"`
}
