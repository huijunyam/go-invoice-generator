package generator

// UnicodeTranslateFunc ...
type UnicodeTranslateFunc func(string) string

// Options for Document
type Options struct {
	AutoPrint bool `json:"auto_print,omitempty"`

	CurrencySymbol    string `default:"€ " json:"currency_symbol,omitempty"`
	CurrencyPrecision int    `default:"2" json:"currency_precision,omitempty"`
	CurrencyDecimal   string `default:"." json:"currency_decimal,omitempty"`
	CurrencyThousand  string `default:" " json:"currency_thousand,omitempty"`

	TextTypeInvoice      string `default:"INVOICE" json:"text_type_invoice,omitempty"`
	TextTypeQuotation    string `default:"QUOTATION" json:"text_type_quotation,omitempty"`
	TextTypeDeliveryNote string `default:"DELIVERY NOTE" json:"text_type_delivery_note,omitempty"`

	TextRefTitle          string `default:"Reference" json:"text_ref_title,omitempty"`
	TextClientRefTitle    string `default:"Customer Reference" json:"text_client_ref_title,omitempty"`
	TextVersionTitle      string `default:"Version" json:"text_version_title,omitempty"`
	TextDateTitle         string `default:"Order Date" json:"text_date_title,omitempty"`
	TextExpiresDateTitle  string `default:"Expiry Date" json:"text_expires_date_title,omitempty"`
	TextDeliveryDateTitle string `default:"Delivery Date" json:"text_delivery_date_title,omitempty"`
	TextPaymentTermTitle  string `default:"Payment term" json:"text_payment_term_title,omitempty"`

	TextItemsNameTitle        string `default:"Name" json:"text_items_name_title,omitempty"`
	TextItemsSKUTitle         string `default:"SKU" json:"text_items_sku_title,omitempty"`
	TextItemsUnitCostTitle    string `default:"Unit price" json:"text_items_unit_cost_title,omitempty"`
	TextItemsMinQuantityTitle string `default:"Min order qty" json:"text_items_min_quantity_title,omitempty"`
	TextItemsQuantityTitle    string `default:"Qty" json:"text_items_quantity_title,omitempty"`
	TextItemsTotalHTTitle     string `default:"Total no tax" json:"text_items_total_ht_title,omitempty"`
	TextItemsTaxTitle         string `default:"Tax" json:"text_items_tax_title,omitempty"`
	TextItemsDiscountTitle    string `default:"Discount" json:"text_items_discount_title,omitempty"`
	TextItemsTotalTTCTitle    string `default:"Total" json:"text_items_total_ttc_title,omitempty"`

	TextTotalTotal      string `default:"SUBTOTAL" json:"text_total_total,omitempty"`
	TextTotalDiscounted string `default:"DISCOUNT" json:"text_total_discounted,omitempty"`
	TextTotalTax        string `default:"TAX" json:"text_total_tax,omitempty"`
	TextTotalShipping   string `default:"FREIGHT" json:"text_total_shipping,omitempty"`
	TextTotalWithTax    string `default:"TOTAL" json:"text_total_with_tax,omitempty"`

	BaseTextColor []int `default:"[35,35,35]" json:"base_text_color,omitempty"`
	GreyTextColor []int `default:"[82,82,82]" json:"grey_text_color,omitempty"`
	GreyBgColor   []int `default:"[232,232,232]" json:"grey_bg_color,omitempty"`
	DarkBgColor   []int `default:"[212,212,212]" json:"dark_bg_color,omitempty"`

	Font     string `default:"Helvetica"`
	BoldFont string `default:"Helvetica"`

	UnicodeTranslateFunc UnicodeTranslateFunc
}
