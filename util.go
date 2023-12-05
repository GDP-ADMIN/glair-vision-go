package glair

// String is an utility function to convert strings
// to its pointer variants
//
// Useful for filling out optional fields
func String(str string) *string {
	return &str
}

// Int is an utility function to convert integers
// to its pointer variants
//
// Useful for filling out optional fields
func Int(num int) *int {
	return &num
}
