package ocr

// BankStatement stores OCR result of Bank Stattement model
// from the given input
type BankStatement = OCRResult[BankStatementData]

type BankStatementData struct {
	BankName               PaperlessOCRStringField    `json:"bank_name,omitempty"`
	AccountNumber          PaperlessOCRStringField    `json:"account_number,omitempty"`
	AccountName            PaperlessOCRStringField    `json:"account_name,omitempty"`
	Currency               PaperlessOCRStringField    `json:"currency,omitempty"`
	BeginningBalance       PaperlessOCRIntField       `json:"beginning_balance,omitempty"`
	TotalDebitTransaction  PaperlessOCRIntField       `json:"total_debit_transaction,omitempty"`
	TotalCreditTransaction PaperlessOCRIntField       `json:"total_credit_transaction,omitempty"`
	TotalTransaction       PaperlessOCRIntField       `json:"total_transaction,omitempty"`
	ProductName            PaperlessOCRStringField    `json:"product_name,omitempty"`
	StartPeriod            PaperlessOCRStringField    `json:"start_period,omitempty"`
	EndPeriod              PaperlessOCRStringField    `json:"end_period,omitempty"`
	Transactions           []BankStatementTransaction `json:"transactions,omitempty"`
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

type BankStatementOCRTransactionField struct {
	Confidence     float32 `json:"confidence"`
	ConfidenceText float32 `json:"confidence_text"`
	PageIndex      int     `json:"page_index"`
	PolygonText    [][]int `json:"polygon_text,omitempty"`
	ValueOriginal  string  `json:"value_original,omitempty"`
}

type BankStatementOCRStringTransactionField struct {
	BankStatementOCRTransactionField

	Value string `json:"value"`
}

type BankStatementOCRIntTransactionField struct {
	BankStatementOCRTransactionField

	Value int64 `json:"value"`
}
