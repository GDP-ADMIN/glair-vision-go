package glair

// FaceMatchingInput stores parameters for Face Matching request
type FaceMatchingInput struct {
	// RequestID represents request identifier for debugging purposes
	RequestID string
	// StoredImage represents the stored image that will be used
	// as a base for face matching
	//
	// StoredImage must be provided as a string that represents a path to the
	// image or an object that implements *os.File
	StoredImage interface{}
	// CapturedImage represents the captured image that will be compared
	// to the stored image.
	//
	// CapturedImage must be provided as a string that represents a path to the
	// image or an object that implements *os.File
	CapturedImage interface{}
}

// PassiveLivenessInput stores parameters for passive liveness requests
type PassiveLivenessInput struct {
	// RequestID represents request identifier for debugging purposes
	RequestID string
	// Image represents the input image that will be processed by GLAIR
	// Vision Face Matching API.
	//
	// Image must be provided as a string that represents a path to the
	// image or an object that implements *os.Image
	Image interface{}
}

// ActiveLivenessInput stores parameters for active liveness requests
type ActiveLivenessInput struct {
	// RequestID represents request identifier for debugging purposes
	RequestID string
	// Image represents the input image that will be processed by GLAIR
	// Vision Face Matching API.
	//
	// Image must be provided as a string that represents a path to the
	// image or an object that implements *os.Image
	Image interface{}
	// GestureCode represents gesture idenfification code that will
	// be used to determine which gesture should be used to detect
	// liveness from the image.
	//
	// Please refer to https://docs.glair.ai/vision/active-liveness
	// for the list of all supported gesture codes
	GestureCode string
}
