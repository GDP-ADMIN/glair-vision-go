package glair

// String is an utility function to convert strings
// to its pointer variants
//
// Useful for filling out optional fields
func String(str string) *string {
	return &str
}
