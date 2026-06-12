package ocr

// SKPR stores OCR result of SKPR model from the given input
type SKPR = OCRResult[SKPRData]

type SKPRData struct {
	DebitAccountRequestType    OCRField[string]   `json:"debit_account_request_type,omitempty"`
	CompleteName               OCRField[string]   `json:"complete_name,omitempty"`
	IDCardNumber               OCRField[string]   `json:"id_card_number,omitempty"`
	ValidityPeriod             OCRField[string]   `json:"validity_period,omitempty"`
	ActFor                     OCRField[string]   `json:"act_for,omitempty"`
	Position                   OCRField[string]   `json:"position,omitempty"`
	CompanyName                OCRField[string]   `json:"company_name,omitempty"`
	Address                    OCRField[string]   `json:"address,omitempty"`
	RtRw                       OCRField[string]   `json:"rt_rw,omitempty"`
	District                   OCRField[string]   `json:"district,omitempty"`
	City                       OCRField[string]   `json:"city,omitempty"`
	ZipCode                    OCRField[string]   `json:"zip_code,omitempty"`
	Province                   OCRField[string]   `json:"province,omitempty"`
	AttorneyInFactName         OCRField[string]   `json:"attorney_in_fact_name,omitempty"`
	AttorneyInFactAddress      OCRField[string]   `json:"attorney_in_fact_address,omitempty"`
	AttorneyInFactIDCardNumber OCRField[string]   `json:"attorney_in_fact_id_card_number,omitempty"`
	AttorneyInFactCompanyCode  OCRField[string]   `json:"attorney_in_fact_company_code,omitempty"`
	OfficePhone                OCRField[string]   `json:"office_phone,omitempty"`
	HomePhone                  OCRField[string]   `json:"home_phone,omitempty"`
	MobilePhoneNumber1         OCRField[string]   `json:"mobile_phone_number_1,omitempty"`
	MobilePhoneNumber2         OCRField[string]   `json:"mobile_phone_number_2,omitempty"`
	AccountNumber              OCRField[string]   `json:"account_number,omitempty"`
	EmailAddress               OCRField[string]   `json:"email_address,omitempty"`
	BankName                   OCRField[string]   `json:"bank_name,omitempty"`
	AccountCurrency            OCRField[string]   `json:"account_currency,omitempty"`
	PolicyCurrency             OCRField[string]   `json:"policy_currency,omitempty"`
	PolicyAmount               OCRField[string]   `json:"policy_amount,omitempty"`
	SigningDate                OCRField[string]   `json:"signing_date,omitempty"`
	SigningLocation            OCRField[string]   `json:"signing_location,omitempty"`
	BankAccountDetails         BankAccountDetails `json:"bank_account_details,omitempty"`
	PolicyDetails              PolicyDetails      `json:"policy_details,omitempty"`
}

type BankAccountDetails struct {
	BankAccountNumber    OCRField[string] `json:"bank_account_number,omitempty"`
	BankAccountOwnerName OCRField[string] `json:"bank_account_owner_name,omitempty"`
}

type PolicyDetails struct {
	PolicyNumber                 OCRField[string] `json:"policy_number,omitempty"`
	PolicyHolderName             OCRField[string] `json:"policyholder_name,omitempty"`
	RelationshipWithPolicyHolder OCRField[string] `json:"relationship_with_policyholder,omitempty"`
}
