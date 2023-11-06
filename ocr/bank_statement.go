package ocr

// BankStatement stores OCR result of Bank Stattement model
// from the given input
type BankStatement = OCRResult[BankStatementData]

type BankStatementData struct {
	BankName               BankStatementOCRStringField `json:"bank_name,omitempty"`
	AccountNumber          BankStatementOCRStringField `json:"account_number,omitempty"`
	AccountName            BankStatementOCRStringField `json:"account_name,omitempty"`
	Currency               BankStatementOCRStringField `json:"currency,omitempty"`
	BeginningBalance       BankStatementOCRIntField    `json:"beginning_balance,omitempty"`
	TotalDebitTransaction  BankStatementOCRIntField    `json:"total_debit_transaction,omitempty"`
	TotalCreditTransaction BankStatementOCRIntField    `json:"total_credit_transaction,omitempty"`
	TotalTransaction       BankStatementOCRIntField    `json:"total_transaction,omitempty"`
	ProductName            BankStatementOCRStringField `json:"product_name,omitempty"`
	StartPeriod            BankStatementOCRStringField `json:"start_period,omitempty"`
	EndPeriod              BankStatementOCRStringField `json:"end_period,omitempty"`
	Transactions           []BankStatementTransaction  `json:"transactions,omitempty"`
}

type BankStatementTransaction struct {
	PostingDate       BankStatementOCRStringTransactionField `json:"posting_date,omitempty"`
	PostingTime       BankStatementOCRStringTransactionField `json:"posting_time,omitempty"`
	EffectiveDate     BankStatementOCRStringTransactionField `json:"effective_date,omitempty"`
	EffectiveTime     BankStatementOCRStringTransactionField `json:"effective_time,omitempty"`
	Description       BankStatementOCRStringTransactionField `json:"description,omitempty"`
	DebitTransaction  BankStatementOCRIntTransactionField    `json:"debit_transaction,omitempty"`
	CreditTransaction BankStatementOCRIntTransactionField    `json:"credit_transaction,omitempty"`
	SignedAmount      BankStatementOCRIntTransactionField    `json:"signed_amount,omitempty"`
}

type BankStatementOCRField struct {
	OCRField

	ConfidenceText float32 `json:"confidence_text,omitempty"`
	PageIndex      int     `json:"page_index"`
	ValueOriginal  string  `json:"value_original,omitempty"`
}

type BankStatementOCRTransactionField struct {
	Confidence     float32 `json:"confidence"`
	ConfidenceText float32 `json:"confidence_text"`
	PageIndex      int     `json:"page_index"`
	PolygonText    [][]int `json:"polygon_text,omitempty"`
	ValueOriginal  string  `json:"value_original,omitempty"`
}

type BankStatementOCRStringField struct {
	BankStatementOCRField

	Value string `json:"value"`
}

type BankStatementOCRIntField struct {
	BankStatementOCRField
	Value int64 `json:"value"`
}

type BankStatementOCRStringTransactionField struct {
	BankStatementOCRTransactionField

	Value string `json:"value"`
}

type BankStatementOCRIntTransactionField struct {
	BankStatementOCRTransactionField

	Value int64 `json:"value"`
}
