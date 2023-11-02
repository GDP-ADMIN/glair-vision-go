package ocr

// Receipt stores OCR result of Receipt model from the given input
type Receipt = OCRResult[ReceiptData]

type ReceiptData struct {
	MerchantName    string        `json:"merchant_name,omitempty"`
	MerchantAddress string        `json:"merchant_address,omitempty"`
	ReceiptDate     string        `json:"receipt_date,omitempty"`
	ReceiptTime     string        `json:"receipt_time.omitempty"`
	Items           []ReceiptItem `json:"items,omitempty"`
	SubTotalAmount  string        `json:"sub_total_amount,omitempty"`
	TaxAmount       string        `json:"tax_amount,omitempty"`
	TipAmount       string        `json:"tip_amount,omitempty"`
	TotalAmount     string        `json:"total_amount,omitempty"`
}

type ReceiptItem struct {
	ItemName        string `json:"item_name,omitempty"`
	ItemQuantity    string `json:"item_quantity,omitempty"`
	ItemTotalPrice  string `json:"item_total_price,omitempty"`
	ItemProductCode string `json:"item_product_code"`
}
