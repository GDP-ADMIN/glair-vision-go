package ocr

// Invoice stores OCR result of Invoice model from the given input
type Invoice = OCRResult[InvoiceData]

type InvoiceData struct {
	InvoiceNumber              OCRStringField `json:"invoice_number,omitempty"`
	InvoiceDate                OCRStringField `json:"invoice_date,omitempty"`
	InvoiceDueDate             OCRStringField `json:"invoice_due_date,omitempty"`
	PurchaseOrder              OCRStringField `json:"purchase_order,omitempty"`
	VendorName                 OCRStringField `json:"vendor_name,omitempty"`
	VendorTaxId                OCRStringField `json:"vendor_tax_id,omitempty"`
	VendorAddress              OCRStringField `json:"vendor_address,omitempty"`
	CustomerName               OCRStringField `json:"customer_name,omitempty"`
	CustomerTaxId              OCRStringField `json:"customer_tax_id,omitempty"`
	CustomerAddress            OCRStringField `json:"customer_address,omitempty"`
	ShippingAddress            OCRStringField `json:"shipping_address,omitempty"`
	ShippingAddressRecipient   OCRStringField `json:"shipping_address_recipient,omitempty"`
	PaymentTerm                OCRStringField `json:"payment_term,omitempty"`
	SubTotalAmount             OCRIntField    `json:"sub_total_amount,omitempty"`
	InvoiceTotalTax            OCRIntField    `json:"invoice_total_tax,omitempty"`
	InvoiceTotal               OCRIntField    `json:"invoice_total,omitempty"`
	RemittanceAddress          OCRStringField `json:"remittance_address,omitempty"`
	RemittanceAddressRecipient OCRStringField `json:"remittance_address_recipient,omitempty"`
	Currency                   OCRStringField `json:"currency,omitempty"`
	Items                      []InvoiceItem  `json:"items,omitempty"`
}

type InvoiceItem struct {
	ItemName        OCRStringField `json:"item_name,omitempty"`
	ItemQuantity    OCRStringField `json:"item_quantity,omitempty"`
	ItemPrice       OCRIntField    `json:"item_price,omitempty"`
	ItemUnitPrice   OCRStringField `json:"item_unit_price,omitempty"`
	ItemProductCode OCRStringField `json:"item_product_code,omitempty"`
	ItemUnit        OCRStringField `json:"item_unit,omitempty"`
}
