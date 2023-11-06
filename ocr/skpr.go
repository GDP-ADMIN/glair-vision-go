package ocr

// SKPR stores OCR result of SKPR model from the given input
type SKPR = OCRResult[SKPRData]

type SKPRData struct {
	DebitAccountRequestType    OCRField           `json:"debit_account_request_type,omitempty"`
	CompleteName               OCRField           `json:"complete_name,omitempty"`
	IDCardNumber               OCRField           `json:"id_card_number,omitempty"`
	ValidityPeriod             OCRField           `json:"validity_period,omitempty"`
	ActFor                     OCRField           `json:"act_for,omitempty"`
	Position                   OCRField           `json:"position,omitempty"`
	CompanyName                OCRField           `json:"company_name,omitempty"`
	Address                    OCRField           `json:"address,omitempty"`
	RtRw                       OCRField           `json:"rt_rw,omitempty"`
	District                   OCRField           `json:"district,omitempty"`
	City                       OCRField           `json:"city,omitempty"`
	ZipCode                    OCRField           `json:"zip_code,omitempty"`
	Province                   OCRField           `json:"province,omitempty"`
	AttorneyInFactName         OCRField           `json:"attorney_in_fact_name,omitempty"`
	AttorneyInFactAddress      OCRField           `json:"attorney_in_fact_address,omitempty"`
	AttorneyInFactIDCardNumber OCRField           `json:"attorney_in_fact_id_card_number,omitempty"`
	AttorneyInFactCompanyCode  OCRField           `json:"attorney_in_fact_company_code,omitempty"`
	OfficePhone                OCRField           `json:"office_phone,omitempty"`
	HomePhone                  OCRField           `json:"home_phone,omitempty"`
	MobilePhoneNumber1         OCRField           `json:"mobile_phone_number_1,omitempty"`
	MobilePhoneNumber2         OCRField           `json:"mobile_phone_number_2,omitempty"`
	AccountNumber              OCRField           `json:"account_number,omitempty"`
	EmailAddress               OCRField           `json:"email_address,omitempty"`
	BankName                   OCRField           `json:"bank_name,omitempty"`
	AccountCurrency            OCRField           `json:"account_currency,omitempty"`
	PolicyCurrency             OCRField           `json:"policy_currency,omitempty"`
	PolicyAmount               OCRField           `json:"policy_amount,omitempty"`
	SigningDate                OCRField           `json:"signing_date,omitempty"`
	SigningLocation            OCRField           `json:"signing_location,omitempty"`
	BankAccountDetails         BankAccountDetails `json:"bank_account_details,omitempty"`
	PolicyDetails              PolicyDetails      `json:"policy_details,omitempty"`
}

type BankAccountDetails struct {
	BankAccountNumber    OCRField `json:"bank_account_number,omitempty"`
	BankAccountOwnerName OCRField `json:"bank_account_owner_name,omitempty"`
}

type PolicyDetails struct {
	PolicyNumber                 OCRField `json:"policy_number,omitempty"`
	PolicyHolderName             OCRField `json:"policyholder_name,omitempty"`
	RelationshipWithPolicyHolder OCRField `json:"relationship_with_policyholder,omitempty"`
}
