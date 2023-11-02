package ocr

type BankStatement = OCRResult[BankStatementData]

type BankStatementData struct {
	BankName              string                     `json:"bank_name,omitempty"`
	AccountNumber         string                     `json:"account_number,omitempty"`
	AccountName           string                     `json:"account_name,omitempty"`
	Currency              string                     `json:"currency,omitempty"`
	BeginningBalance      string                     `json:"beginning_balance,omitempty"`
	TotalDebitTransaction string                     `json:"total_debit_transaction,omitempty"`
	ProductName           string                     `json:"product_name,omitempty"`
	StartPeriod           string                     `json:"start_period,omitempty"`
	EndPeriod             string                     `json:"end_period,omitempty"`
	Transactions          []BankStatementTransaction `json:"transactions,omitempty"`
}

type BankStatementTransaction struct {
	PostingDate       string `json:"posting_date,omitempty"`
	PostingTime       string `json:"posting_time,omitempty"`
	EffectiveDate     string `json:"effective_date,omitempty"`
	EffectiveTime     string `json:"effective_time,omitempty"`
	Description       string `json:"description,omitempty"`
	DebitTransaction  string `json:"debig_transaction,omitempty"`
	CreditTransaction string `json:"credit_transaction,omitempty"`
	SignedAmount      string `json:"signed_amount,omitempty"`
}
