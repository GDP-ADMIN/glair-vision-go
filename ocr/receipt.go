package ocr

// Receipt stores OCR result of Receipt model from the given input
type Receipt = OCRResult[ReceiptData]

type ReceiptData struct {
	MerchantName        OCRField[string] `json:"merchant_name,omitempty"`
	MerchantAddress     OCRField[string] `json:"merchant_address,omitempty"`
	MerchantPlace       OCRField[string] `json:"merchant_place,omitempty"`
	ReceiptDate         OCRField[string] `json:"receipt_date,omitempty"`
	ReceiptNumber       OCRField[string] `json:"receipt_number,omitempty"`
	ReceiptOrder        OCRField[string] `json:"receipt_order,omitempty"`
	ReceiptTime         OCRField[string] `json:"receipt_time,omitempty"`
	SubTotalAmount      OCRField[int64]  `json:"sub_total_amount,omitempty"`
	TaxAmount           OCRField[int64]  `json:"tax_amount,omitempty"`
	DiscountAmount      OCRField[int64]  `json:"discount_amount,omitempty"`
	ServiceChargeAmount OCRField[int64]  `json:"service_charge_amount,omitempty"`
	TipAmount           OCRField[int64]  `json:"tip_amount,omitempty"`
	TotalAmount         OCRField[int64]  `json:"total_amount,omitempty"`
	CCLast4Digit        OCRField[string] `json:"cc_last_4_digit,omitempty"`
	Currency            OCRField[string] `json:"currency,omitempty"`
	MerchantPhoneNumber OCRField[string] `json:"merchant_phone_number,omitempty"`
	PaymentMethod       OCRField[string] `json:"payment_method,omitempty"`
	PaymentProduct      OCRField[string] `json:"payment_product,omitempty"`

	Items []ReceiptItem `json:"items,omitempty"`
}

type ReceiptItem struct {
	ItemName         OCRField[string] `json:"item_name,omitempty"`
	ItemQuantity     OCRField[int64]  `json:"item_quantity,omitempty"`
	ItemTotalPrice   OCRField[int64]  `json:"item_total_price,omitempty"`
	ItemProductCode  OCRField[string] `json:"item_product_code,omitempty"`
	ItemEachPrice    OCRField[int64]  `json:"item_each_price,omitempty"`
	ItemQuantityUnit OCRField[string] `json:"item_quantity_unit,omitempty"`
}
