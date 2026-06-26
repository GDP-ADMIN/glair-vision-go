package ocr

// Passport stores OCR result of Passport model from the given input.
type Passport = OCRResult[PassportData]

type PassportData struct {
	BirthDate       OCRPageField[string] `json:"birth_date,omitempty"`
	BirthDateHash   OCRPageField[string] `json:"birth_date_hash,omitempty"`
	Country         OCRPageField[string] `json:"country,omitempty"`
	DocNumber       OCRPageField[string] `json:"doc_number,omitempty"`
	DocNumberHash   OCRPageField[string] `json:"doc_number_hash,omitempty"`
	DocumentType    OCRPageField[string] `json:"document_type,omitempty"`
	ExpiryDate      OCRPageField[string] `json:"expiry_date,omitempty"`
	ExpiryDateHash  OCRPageField[string] `json:"expiry_date_hash,omitempty"`
	FinalHash       OCRPageField[string] `json:"final_hash,omitempty"`
	Name            OCRPageField[string] `json:"name,omitempty"`
	Nationality     OCRPageField[string] `json:"nationality,omitempty"`
	OptionalData    OCRPageField[string] `json:"optional_data,omitempty"`
	OptionalDataHash OCRPageField[string] `json:"optional_data_hash,omitempty"`
	Sex             OCRPageField[string] `json:"sex,omitempty"`
	Surname         OCRPageField[string] `json:"surname,omitempty"`
}
