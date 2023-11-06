package ocr

// SKPR stores OCR result of SKPR model from the given input
type SKPR = OCRResult[SKPRData]

type SKPRData struct {
	DebitAccountRequestType    PaperlessOCRStringField `json:"debit_account_request_type,omitempty"`
	CompleteName               PaperlessOCRStringField `json:"complete_name,omitempty"`
	IDCardNumber               PaperlessOCRStringField `json:"id_card_number,omitempty"`
	ValidityPeriod             PaperlessOCRStringField `json:"validity_period,omitempty"`
	ActFor                     PaperlessOCRStringField `json:"act_for,omitempty"`
	Position                   PaperlessOCRStringField `json:"position,omitempty"`
	CompanyName                PaperlessOCRStringField `json:"company_name,omitempty"`
	Address                    PaperlessOCRStringField `json:"address,omitempty"`
	RtRw                       PaperlessOCRStringField `json:"rt_rw,omitempty"`
	District                   PaperlessOCRStringField `json:"district,omitempty"`
	City                       PaperlessOCRStringField `json:"city,omitempty"`
	ZipCode                    PaperlessOCRStringField `json:"zip_code,omitempty"`
	Province                   PaperlessOCRStringField `json:"province,omitempty"`
	AttorneyInFactName         PaperlessOCRStringField `json:"attorney_in_fact_name,omitempty"`
	AttorneyInFactAddress      PaperlessOCRStringField `json:"attorney_in_fact_address,omitempty"`
	AttorneyInFactIDCardNumber PaperlessOCRStringField `json:"attorney_in_fact_id_card_number,omitempty"`
	AttorneyInFactCompanyCode  PaperlessOCRStringField `json:"attorney_in_fact_company_code,omitempty"`
	OfficePhone                PaperlessOCRStringField `json:"office_phone,omitempty"`
	HomePhone                  PaperlessOCRStringField `json:"home_phone,omitempty"`
	MobilePhoneNumber1         PaperlessOCRStringField `json:"mobile_phone_number_1,omitempty"`
	MobilePhoneNumber2         PaperlessOCRStringField `json:"mobile_phone_number_2,omitempty"`
	AccountNumber              PaperlessOCRStringField `json:"account_number,omitempty"`
	EmailAddress               PaperlessOCRStringField `json:"email_address,omitempty"`
	BankName                   PaperlessOCRStringField `json:"bank_name,omitempty"`
	AccountCurrency            PaperlessOCRStringField `json:"account_currency,omitempty"`
	PolicyCurrency             PaperlessOCRStringField `json:"policy_currency,omitempty"`
	PolicyAmount               PaperlessOCRStringField `json:"policy_amount,omitempty"`
	SigningDate                PaperlessOCRStringField `json:"signing_date,omitempty"`
	SigningLocation            PaperlessOCRStringField `json:"signing_location,omitempty"`
	BankAccountDetails         BankAccountDetails      `json:"bank_account_details,omitempty"`
	PolicyDetails              PolicyDetails           `json:"policy_details,omitempty"`
}

type BankAccountDetails struct {
	BankAccountNumber    PaperlessOCRStringField `json:"bank_account_number,omitempty"`
	BankAccountOwnerName PaperlessOCRStringField `json:"bank_account_owner_name,omitempty"`
}

type PolicyDetails struct {
	PolicyNumber                 PaperlessOCRStringField `json:"policy_number,omitempty"`
	PolicyHolderName             PaperlessOCRStringField `json:"policyholder_name,omitempty"`
	RelationshipWithPolicyHolder PaperlessOCRStringField `json:"relationship_with_policyholder,omitempty"`
}
