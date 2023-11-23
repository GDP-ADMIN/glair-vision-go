package glair

// OCRInput stores parameters for OCR requests
type OCRInput struct {
	// RequestID represents request identifier for debugging purposes
	RequestID string
	// File represents the input image that will be processed by GLAIR
	// Vision OCR API.
	//
	// File must be provided as a string that represents a path to the
	// image or an object that implements *os.File
	File interface{}
}
