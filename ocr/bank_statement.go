package ocr

// BankStatement stores OCR result of Bank Stattement model
// from the given input
type BankStatement = OCRResult[BankStatementData]

type BankStatementData struct {
	BankName               OCRField[string]           `json:"bank_name,omitempty"`
	AccountNumber          OCRField[string]           `json:"account_number,omitempty"`
	AccountName            OCRField[string]           `json:"account_name,omitempty"`
	Currency               OCRField[string]           `json:"currency,omitempty"`
	BeginningBalance       OCRField[int64]            `json:"beginning_balance,omitempty"`
	TotalDebitTransaction  OCRField[int64]            `json:"total_debit_transaction,omitempty"`
	TotalCreditTransaction OCRField[int64]            `json:"total_credit_transaction,omitempty"`
	TotalTransaction       OCRField[int64]            `json:"total_transaction,omitempty"`
	ProductName            OCRField[string]           `json:"product_name,omitempty"`
	StartPeriod            OCRField[string]           `json:"start_period,omitempty"`
	EndPeriod              OCRField[string]           `json:"end_period,omitempty"`
	Transactions           []BankStatementTransaction `json:"transactions,omitempty"`
}

type BankStatementTransaction struct {
	PostingDate       OCRField[string] `json:"posting_date,omitempty"`
	PostingTime       OCRField[string] `json:"posting_time,omitempty"`
	EffectiveDate     OCRField[string] `json:"effective_date,omitempty"`
	EffectiveTime     OCRField[string] `json:"effective_time,omitempty"`
	Description       OCRField[string] `json:"description,omitempty"`
	DebitTransaction  OCRField[int64]  `json:"debit_transaction,omitempty"`
	CreditTransaction OCRField[int64]  `json:"credit_transaction,omitempty"`
	SignedAmount      OCRField[int64]  `json:"signed_amount,omitempty"`
}
