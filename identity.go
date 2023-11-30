package glair

// BasicVerificationInput stores parameters for basic identity
// verification request
type BasicVerificationInput struct {
	// RequestID represents request identifier for debugging purposes
	RequestID string
	// Nik represents single identification number of the
	// person to be verified
	Nik string
	// Name represents the name of the person to be verified
	//
	// Name is optional
	Name *string
	// DateOfBirth represents date of birth of the person to
	// be verified
	//
	// DateOfBirth must be provided as a time.Time instance
	// or as a string in dd-mm-yyyy format
	//
	// DateOfBirth is optional
	DateOfBirth interface{}
	// NoKk represents family registration number of the person
	//
	// NoKk is optional
	NoKk *string
	// MotherMaidenName represents identity holder's mother
	// maiden name
	//
	// MotherMaidenName is optional
	MotherMaidenName *string
	// PlaceOfBirth represents place of birth of the identity holder
	// in their KTP
	//
	// PlaceOfBirth is optional
	PlaceOfBirth *string
	// Address represents address of the identity holder
	// in their KTP
	//
	// Address is optional
	Address *string
	// Gender represents gender of the identity holder
	// in their KTP
	//
	// Gender is optional
	Gender *string
	// MaritalStatus represents marital status of the identity holder
	// in their KTP
	//
	// MaritalStatus is optional
	MaritalStatus *string
	// JobType represents job of the identity holder
	// in their KTP
	//
	// JobType is optional
	JobType *string
	// Province represents the province of residence of the identity holder
	// in their KTP
	//
	// Province is optional
	Province *string
	// City represents the city of residence of the identity holder
	// in their KTP
	//
	// City is optional
	City *string
	// District represents the district of residence of the identity holder
	// in their KTP
	//
	// District is optional
	District *string
	// Subdistrict represents the subdistrict of residence of the identity holder
	// in their KTP
	//
	// Subdistrict is optional
	Subdistrict *string
	// Rt represents the RT of residence of the identity holder
	// in their KTP
	//
	// Rt is optional
	Rt *string
	// Rw represents the RW of residence of the identity holder
	// in their KTP
	//
	// Rw is optional
	Rw *string
}

// FaceVerificationInput stores parameters for identity
// verification request with face
type FaceVerificationInput struct {
	// RequestID represents request identifier for debugging purposes
	RequestID string
	// Nik represents single identification number of the
	// person to be verified
	Nik string
	// FaceImage represents the input image that will be used
	// to verify the identity of the identity card
	//
	// FaceImage must be provided as a string that represents a path to the
	// image or an object that implements *os.File
	FaceImage interface{}
	// Name represents the name of the person to be verified
	Name string
	// DateOfBirth represents date of birth of the person to
	// be verified
	//
	// DateOfBirth must be provided as a string with yyyy-mm-dd
	// format or as an instance of *time.Time
	DateOfBirth interface{}
}
