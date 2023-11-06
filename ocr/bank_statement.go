package ocr

// BankStatement stores OCR result of Bank Stattement model
// from the given input
type BankStatement = OCRResult[BankStatementData]

type BankStatementData struct {
	BankName              OCRStringField             `json:"bank_name,omitempty"`
	AccountNumber         OCRStringField             `json:"account_number,omitempty"`
	AccountName           OCRStringField             `json:"account_name,omitempty"`
	Currency              OCRStringField             `json:"currency,omitempty"`
	BeginningBalance      OCRStringField             `json:"beginning_balance,omitempty"`
	TotalDebitTransaction OCRStringField             `json:"total_debit_transaction,omitempty"`
	ProductName           OCRStringField             `json:"product_name,omitempty"`
	StartPeriod           OCRStringField             `json:"start_period,omitempty"`
	EndPeriod             OCRStringField             `json:"end_period,omitempty"`
	Transactions          []BankStatementTransaction `json:"transactions,omitempty"`
}

type BankStatementTransaction struct {
	PostingDate       OCRStringField `json:"posting_date,omitempty"`
	PostingTime       OCRStringField `json:"posting_time,omitempty"`
	EffectiveDate     OCRStringField `json:"effective_date,omitempty"`
	EffectiveTime     OCRStringField `json:"effective_time,omitempty"`
	Description       OCRStringField `json:"description,omitempty"`
	DebitTransaction  OCRStringField `json:"debig_transaction,omitempty"`
	CreditTransaction OCRStringField `json:"credit_transaction,omitempty"`
	SignedAmount      OCRStringField `json:"signed_amount,omitempty"`
}
