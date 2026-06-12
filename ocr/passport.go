package ocr

// Passport stores OCR result of Passport model from the given input
type Passport = OCRResult[PassportData]

type PassportData struct {
	BirthDate      OCRField[string] `json:"birth_date,omitempty"`
	BirthDateHash  OCRField[string] `json:"birth_date_hash,omitempty"`
	Country        OCRField[string] `json:"country,omitempty"`
	DocNumber      OCRField[string] `json:"doc_number,omitempty"`
	DocNumberHash  OCRField[string] `json:"doc_number_hash,omitempty"`
	DocumentType   OCRField[string] `json:"document_type,omitempty"`
	ExpiryDate     OCRField[string] `json:"expiry_date,omitempty"`
	ExpiryDateHash OCRField[string] `json:"expiry_date_hash,omitempty"`
	FinalHash      OCRField[string] `json:"final_hash,omitempty"`
	Name           OCRField[string] `json:"name,omitempty"`
	Nationality    OCRField[string] `json:"nationality,omitempty"`
	Sex            OCRField[string] `json:"sex,omitempty"`
	Surname        OCRField[string] `json:"surname,omitempty"`
}
