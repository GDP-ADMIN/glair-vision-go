package identity

// BasicIdentityVerificationResult represents basic
// identity verification result from GLAIR Vision API
// of the given NIK
type BasicIdentityVerificationResult = IdentityVerificationResult[BasicIdentityVerification]

type BasicIdentityVerification struct {
	Nik              bool  `json:"nik"`
	Name             *bool `json:"name,omitempty"`
	DateOfBirth      *bool `json:"date_of_birth,omitempty"`
	NoKk             *bool `json:"no_kk,omitempty"`
	MotherMaidenName *bool `json:"mother_maiden_name,omitempty"`
	PlaceOfBirth     *bool `json:"place_of_birth,omitempty"`
	Address          *bool `json:"address,omitempty"`
	Gender           *bool `json:"gender,omitempty"`
	MaritalStatus    *bool `json:"marital_status,omitempty"`
	JobType          *bool `json:"job_type,omitempty"`
	Province         *bool `json:"province,omitempty"`
	City             *bool `json:"city,omitempty"`
	District         *bool `json:"district,omitempty"`
	Subdistrict      *bool `json:"subdistrict,omitempty"`
	Rt               *bool `json:"rt,omitempty"`
	Rw               *bool `json:"rw,omitempty"`
}
