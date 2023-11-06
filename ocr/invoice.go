package ocr

// Invoice stores OCR result of Invoice model from the given input
type Invoice = OCRResult[InvoiceData]

type InvoiceData struct {
	InvoiceNumber              OCRField      `json:"invoice_number,omitempty"`
	InvoiceDate                OCRField      `json:"invoice_date,omitempty"`
	InvoiceDueDate             OCRField      `json:"invoice_due_date,omitempty"`
	VendorName                 OCRField      `json:"vendor_name,omitempty"`
	VendorTaxId                OCRField      `json:"vendor_tax_id,omitempty"`
	VendorAddress              OCRField      `json:"vendor_address,omitempty"`
	CustomerName               OCRField      `json:"customer_name,omitempty"`
	CustomerTaxId              OCRField      `json:"customer_tax_id,omitempty"`
	CustomerAddress            OCRField      `json:"customer_address,omitempty"`
	ShippingAddress            OCRField      `json:"shipping_address,omitempty"`
	ShippingAddressRecipient   OCRField      `json:"shipping_address_recipient,omitempty"`
	PaymentTerm                OCRField      `json:"payment_term,omitempty"`
	SubTotalAmount             OCRField      `json:"sub_total_amount,omitempty"`
	InvoiceTotalTax            OCRField      `json:"invoice_total_tax,omitempty"`
	RemittanceAddress          OCRField      `json:"remittance_address,omitempty"`
	RemittanceAddressRecipient OCRField      `json:"remittance_address_recipient,omitempty"`
	Currency                   OCRField      `json:"currency,omitempty"`
	Item                       []InvoiceItem `json:"items,omitempty"`
}

type InvoiceItem struct {
	ItemName        OCRField `json:"item_name,omitempty"`
	ItemQuantity    OCRField `json:"item_quantity,omitempty"`
	ItemPrice       OCRField `json:"item_price,omitempty"`
	ItemUnitPrice   OCRField `json:"item_unit_price,omitempty"`
	ItemProductCode OCRField `json:"item_product_code,omitempty"`
	ItemUnit        OCRField `json:"item_unit,omitempty"`
}
