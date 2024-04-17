package ocr

// Invoice stores OCR result of Invoice model from the given input
type Invoice = OCRResult[InvoiceData]

type InvoiceData struct {
	InvoiceNumber       OCRStringField `json:"invoice_number,omitempty"`
	InvoiceDate         OCRStringField `json:"invoice_date,omitempty"`
	VATTax              OCRIntField    `json:"vat_tax,omitempty"`
	IncomeTax           OCRIntField    `json:"income_tax,omitempty"`
	InvoiceTotalTax     OCRIntField    `json:"invoice_total_tax,omitempty"`
	InvoiceTotal        OCRIntField    `json:"invoice_total,omitempty"`
	InvoiceSubTotal     OCRIntField    `json:"invoice_sub_total,omitempty"`
	InvoiceDiscount     OCRIntField    `json:"invoice_discount,omitempty"`
	DownPayment         OCRIntField    `json:"down_payment,omitempty"`
	Voucher             OCRIntField    `json:"voucher,omitempty"`
	VendorName          OCRStringField `json:"vendor_name,omitempty"`
	VendorAddress       OCRStringField `json:"vendor_address,omitempty"`
	VendorTaxId         OCRStringField `json:"vendor_tax_id,omitempty"`
	CustomerName        OCRStringField `json:"customer_name,omitempty"`
	CustomerAddress     OCRStringField `json:"customer_address,omitempty"`
	CustomerTaxId       OCRStringField `json:"customer_tax_id,omitempty"`
	SenderName          OCRStringField `json:"sender_name,omitempty"`
	Currency            OCRStringField `json:"currency,omitempty"`
	PurchaseOrderNumber OCRStringField `json:"purchase_order_number,omitempty"`

	PaymentBanks []PaymentBank `json:"payment_banks,omitempty"`
	Items        []InvoiceItem `json:"items,omitempty"`
}

type PaymentBank struct {
	PaymentBankName      OCRStringField `json:"payment_bank_name,omitempty"`
	PaymentAccountNumber OCRStringField `json:"payment_account_number,omitempty"`
	PaymentAccountName   OCRStringField `json:"payment_account_name,omitempty"`
}

type InvoiceItem struct {
	ItemName         OCRStringField `json:"item_name,omitempty"`
	ItemQuantity     OCRStringField `json:"item_quantity,omitempty"`
	ItemTotalPrice   OCRIntField    `json:"item_total_price,omitempty"`
	ItemEachPrice    OCRIntField    `json:"item_each_price,omitempty"`
	ItemProductCode  OCRStringField `json:"item_product_code,omitempty"`
	ItemQuantityUnit OCRStringField `json:"item_quantity_unit,omitempty"`
}
