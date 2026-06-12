package ocr

// Invoice stores OCR result of Invoice model from the given input
type Invoice = OCRResult[InvoiceData]

type InvoiceData struct {
	InvoiceNumber       OCRField[string] `json:"invoice_number,omitempty"`
	InvoiceDate         OCRField[string] `json:"invoice_date,omitempty"`
	VATTax              OCRField[int64]  `json:"vat_tax,omitempty"`
	IncomeTax           OCRField[int64]  `json:"income_tax,omitempty"`
	InvoiceTotalTax     OCRField[int64]  `json:"invoice_total_tax,omitempty"`
	InvoiceTotal        OCRField[int64]  `json:"invoice_total,omitempty"`
	InvoiceSubTotal     OCRField[int64]  `json:"invoice_sub_total,omitempty"`
	InvoiceDiscount     OCRField[int64]  `json:"invoice_discount,omitempty"`
	DownPayment         OCRField[int64]  `json:"down_payment,omitempty"`
	Voucher             OCRField[int64]  `json:"voucher,omitempty"`
	VendorName          OCRField[string] `json:"vendor_name,omitempty"`
	VendorAddress       OCRField[string] `json:"vendor_address,omitempty"`
	VendorTaxId         OCRField[string] `json:"vendor_tax_id,omitempty"`
	CustomerName        OCRField[string] `json:"customer_name,omitempty"`
	CustomerAddress     OCRField[string] `json:"customer_address,omitempty"`
	CustomerTaxId       OCRField[string] `json:"customer_tax_id,omitempty"`
	SenderName          OCRField[string] `json:"sender_name,omitempty"`
	Currency            OCRField[string] `json:"currency,omitempty"`
	PurchaseOrderNumber OCRField[string] `json:"purchase_order_number,omitempty"`

	PaymentBanks []PaymentBank `json:"payment_banks,omitempty"`
	Items        []InvoiceItem `json:"items,omitempty"`
}

type PaymentBank struct {
	PaymentBankName      OCRField[string] `json:"payment_bank_name,omitempty"`
	PaymentAccountNumber OCRField[string] `json:"payment_account_number,omitempty"`
	PaymentAccountName   OCRField[string] `json:"payment_account_name,omitempty"`
}

type InvoiceItem struct {
	ItemName         OCRField[string] `json:"item_name,omitempty"`
	ItemQuantity     OCRField[string] `json:"item_quantity,omitempty"`
	ItemTotalPrice   OCRField[int64]  `json:"item_total_price,omitempty"`
	ItemEachPrice    OCRField[int64]  `json:"item_each_price,omitempty"`
	ItemProductCode  OCRField[string] `json:"item_product_code,omitempty"`
	ItemQuantityUnit OCRField[string] `json:"item_quantity_unit,omitempty"`
}
