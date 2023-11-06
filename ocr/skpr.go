package ocr

// SKPR stores OCR result of SKPR model from the given input
type SKPR = OCRResult[SKPRData]

type SKPRData struct {
	DebitAccountRequestType    OCRStringField     `json:"debit_account_request_type,omitempty"`
	CompleteName               OCRStringField     `json:"complete_name,omitempty"`
	IDCardNumber               OCRStringField     `json:"id_card_number,omitempty"`
	ValidityPeriod             OCRStringField     `json:"validity_period,omitempty"`
	ActFor                     OCRStringField     `json:"act_for,omitempty"`
	Position                   OCRStringField     `json:"position,omitempty"`
	CompanyName                OCRStringField     `json:"company_name,omitempty"`
	Address                    OCRStringField     `json:"address,omitempty"`
	RtRw                       OCRStringField     `json:"rt_rw,omitempty"`
	District                   OCRStringField     `json:"district,omitempty"`
	City                       OCRStringField     `json:"city,omitempty"`
	ZipCode                    OCRStringField     `json:"zip_code,omitempty"`
	Province                   OCRStringField     `json:"province,omitempty"`
	AttorneyInFactName         OCRStringField     `json:"attorney_in_fact_name,omitempty"`
	AttorneyInFactAddress      OCRStringField     `json:"attorney_in_fact_address,omitempty"`
	AttorneyInFactIDCardNumber OCRStringField     `json:"attorney_in_fact_id_card_number,omitempty"`
	AttorneyInFactCompanyCode  OCRStringField     `json:"attorney_in_fact_company_code,omitempty"`
	OfficePhone                OCRStringField     `json:"office_phone,omitempty"`
	HomePhone                  OCRStringField     `json:"home_phone,omitempty"`
	MobilePhoneNumber1         OCRStringField     `json:"mobile_phone_number_1,omitempty"`
	MobilePhoneNumber2         OCRStringField     `json:"mobile_phone_number_2,omitempty"`
	AccountNumber              OCRStringField     `json:"account_number,omitempty"`
	EmailAddress               OCRStringField     `json:"email_address,omitempty"`
	BankName                   OCRStringField     `json:"bank_name,omitempty"`
	AccountCurrency            OCRStringField     `json:"account_currency,omitempty"`
	PolicyCurrency             OCRStringField     `json:"policy_currency,omitempty"`
	PolicyAmount               OCRStringField     `json:"policy_amount,omitempty"`
	SigningDate                OCRStringField     `json:"signing_date,omitempty"`
	SigningLocation            OCRStringField     `json:"signing_location,omitempty"`
	BankAccountDetails         BankAccountDetails `json:"bank_account_details,omitempty"`
	PolicyDetails              PolicyDetails      `json:"policy_details,omitempty"`
}

type BankAccountDetails struct {
	BankAccountNumber    OCRStringField `json:"bank_account_number,omitempty"`
	BankAccountOwnerName OCRStringField `json:"bank_account_owner_name,omitempty"`
}

type PolicyDetails struct {
	PolicyNumber                 OCRStringField `json:"policy_number,omitempty"`
	PolicyHolderName             OCRStringField `json:"policyholder_name,omitempty"`
	RelationshipWithPolicyHolder OCRStringField `json:"relationship_with_policyholder,omitempty"`
}
