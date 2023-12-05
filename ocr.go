package glair

// OCRInput stores parameters for OCR requests
type OCRInput struct {
	// RequestID represents request identifier for debugging purposes
	RequestID string
	// Image represents the input image that will be processed by GLAIR
	// Vision OCR API.
	//
	// Image must be provided as a string that represents a path to the
	// image or an object that implements *os.Image
	Image interface{}
}

// BPKBInput stores parameters for BPKB requests
type BPKBInput struct {
	// RequestID represents request identifier for debugging purposes
	RequestID string
	// Image represents the input image that will be processed by GLAIR
	// Vision OCR API.
	//
	// Image must be provided as a string that represents a path to the
	// image or an object that implements *os.Image
	Image interface{}
	// Page represents specific page number to be read from the BPKB
	// image file. If this argument is omitted, the API will read
	// all pages.
	//
	// Page is optional
	Page *int
}
