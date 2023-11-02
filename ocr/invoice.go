package ocr

type Invoice = OCRResult[InvoiceData]

type InvoiceData struct {
	InvoiceNumber              string        `json:"invoice_number,omitempty"`
	InvoiceDate                string        `json:"invoice_date,omitempty"`
	InvoiceDueDate             string        `json:"invoice_due_date,omitempty"`
	VendorName                 string        `json:"vendor_name,omitempty"`
	VendorTaxId                string        `json:"vendor_tax_id,omitempty"`
	VendorAddress              string        `json:"vendor_address,omitempty"`
	CustomerName               string        `json:"customer_name,omitempty"`
	CustomerTaxId              string        `json:"customer_tax_id,omitempty"`
	CustomerAddress            string        `json:"customer_address,omitempty"`
	ShippingAddress            string        `json:"shipping_address,omitempty"`
	ShippingAddressRecipient   string        `json:"shipping_address_recipient,omitempty"`
	PaymentTerm                string        `json:"payment_term,omitempty"`
	SubTotalAmount             string        `json:"sub_total_amount,omitempty"`
	InvoiceTotalTax            string        `json:"invoice_total_tax,omitempty"`
	RemittanceAddress          string        `json:"remittance_address,omitempty"`
	RemittanceAddressRecipient string        `json:"remittance_address_recipient,omitempty"`
	Currency                   string        `json:"currency,omitempty"`
	Item                       []InvoiceItem `json:"items,omitempty"`
}

type InvoiceItem struct {
	ItemName        string `json:"item_name,omitempty"`
	ItemQuantity    string `json:"item_quantity,omitempty"`
	ItemPrice       string `json:"item_price,omitempty"`
	ItemUnitPrice   string `json:"item_unit_price,omitempty"`
	ItemProductCode string `json:"item_product_code,omitempty"`
	ItemUnit        string `json:"item_unit,omitempty"`
}
