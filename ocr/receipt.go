package ocr

// Receipt stores OCR result of Receipt model from the given input
type Receipt = OCRResult[ReceiptData]

type ReceiptData struct {
	MerchantName    OCRStringField `json:"merchant_name,omitempty"`
	MerchantAddress OCRStringField `json:"merchant_address,omitempty"`
	ReceiptDate     OCRStringField `json:"receipt_date,omitempty"`
	ReceiptTime     OCRStringField `json:"receipt_time,omitempty"`
	Items           []ReceiptItem  `json:"items,omitempty"`
	SubTotalAmount  OCRStringField `json:"sub_total_amount,omitempty"`
	TaxAmount       OCRStringField `json:"tax_amount,omitempty"`
	TipAmount       OCRStringField `json:"tip_amount,omitempty"`
	TotalAmount     OCRStringField `json:"total_amount,omitempty"`
}

type ReceiptItem struct {
	ItemName        OCRStringField `json:"item_name,omitempty"`
	ItemQuantity    OCRStringField `json:"item_quantity,omitempty"`
	ItemTotalPrice  OCRStringField `json:"item_total_price,omitempty"`
	ItemProductCode OCRStringField `json:"item_product_code"`
}
