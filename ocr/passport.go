package ocr

// Passport stores OCR result of Passport model from the given input
type Passport = OCRResult[PassportData]

type PassportData struct {
	BirthDate      string `json:"birth_date,omitempty"`
	BirthDateHash  string `json:"birth_date_hash,omitempty"`
	Country        string `json:"country,omitempty"`
	DocNumber      string `json:"doc_number,omitempty"`
	DocNumberHash  string `json:"doc_number_hash,omitempty"`
	DocumentType   string `json:"document_type,omitempty"`
	ExpiryDate     string `json:"expiry_date,omitempty"`
	ExpiryDateHash string `json:"expiry_date_hash,omitempty"`
	FinalHash      string `json:"final_hash,omitempty"`
	Name           string `json:"name,omitempty"`
	Nationality    string `json:"nationality,omitempty"`
	Sex            string `json:"sex,omitempty"`
	Surname        string `json:"surname,omitempty"`
}
