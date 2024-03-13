package ocr

// Receipt stores OCR result of Receipt model from the given input
type Receipt = OCRResult[ReceiptData]

type ReceiptData struct {
	MerchantName        OCRStringField `json:"merchant_name,omitempty"`
	MerchantAddress     OCRStringField `json:"merchant_address,omitempty"`
	MerchantPlace       OCRStringField `json:"merchant_place,omitempty"`
	ReceiptDate         OCRStringField `json:"receipt_date,omitempty"`
	ReceiptNumber       OCRStringField `json:"receipt_number,omitempty"`
	ReceiptOrder        OCRStringField `json:"receipt_order,omitempty"`
	ReceiptTime         OCRStringField `json:"receipt_time,omitempty"`
	SubTotalAmount      OCRStringField `json:"sub_total_amount,omitempty"`
	TaxAmount           OCRIntField    `json:"tax_amount,omitempty"`
	DiscountAmount      OCRIntField    `json:"discount_amount,omitempty"`
	ServiceChargeAmount OCRIntField    `json:"service_charge_amount,omitempty"`
	TipAmount           OCRIntField    `json:"tip_amount,omitempty"`
	TotalAmount         OCRStringField `json:"total_amount,omitempty"`
	CCLast4Digit        OCRStringField `json:"cc_last_4_digit,omitempty"`
	Currency            OCRStringField `json:"currency,omitempty"`
	MerchantPhoneNumber OCRStringField `json:"merchant_phone_number,omitempty"`
	PaymentMethod       OCRStringField `json:"payment_method,omitempty"`
	PaymentProduct      OCRStringField `json:"payment_product,omitempty"`

	Items []ReceiptItem `json:"items,omitempty"`
}

type ReceiptItem struct {
	ItemName         OCRStringField `json:"item_name,omitempty"`
	ItemQuantity     OCRStringField `json:"item_quantity,omitempty"`
	ItemTotalPrice   OCRIntField    `json:"item_total_price,omitempty"`
	ItemProductCode  OCRStringField `json:"item_product_code,omitempty"`
	ItemEachPrice    OCRIntField    `json:"item_each_price,omitempty"`
	ItemQuantityUnit OCRStringField `json:"item_quantity_unit,omitempty"`
}
