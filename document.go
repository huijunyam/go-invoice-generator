package generator

import (
	"github.com/go-pdf/fpdf"
	"github.com/leekchan/accounting"
)

// Document define base document
type Document struct {
	pdf *fpdf.Fpdf
	ac  accounting.Accounting

	Options *Options      `json:"options,omitempty"`
	Header  *HeaderFooter `json:"header,omitempty"`
	Footer  *HeaderFooter `json:"footer,omitempty"`
	Type    string        `json:"type,omitempty" validate:"required,oneof=INVOICE DELIVERY_NOTE QUOTATION"`
	Ref     string        `json:"ref,omitempty" validate:"required,min=1,max=32"`
	// Version      string        `json:"version,omitempty" validate:"max=32"`
	ClientRef          string          `json:"client_ref,omitempty" validate:"max=64"`
	Description        string          `json:"description,omitempty" validate:"max=1024"`
	Notes              string          `json:"notes,omitempty"`
	Company            *Contact        `json:"company,omitempty" validate:"required"`
	Customer           *Contact        `json:"customer,omitempty" validate:"required"`
	Items              []*Item         `json:"items,omitempty"`
	ShippingItems      []*ShippingItem `json:"shipping_items,omitempty"`
	QuoteItems         []*QuoteItem    `json:"quote_items,omitempty"`
	Date               string          `json:"date,omitempty"`
	DeliveryDate       string          `json:"delivery_date,omitempty"`
	ValidityDate       string          `json:"validity_date,omitempty"`
	PaymentTerm        string          `json:"payment_term,omitempty"`
	DiscountPercentage string          `json:"discount_percentage,omitempty"`
	DiscountAmount     string          `json:"discount_amount,omitempty"`
	TaxPercentage      string          `json:"tax_percentage,omitempty"`
	TaxAmount          string          `json:"tax_amount,omitempty"`
	Subtotal           string          `json:"subtotal,omitempty"`
	Total              string          `json:"total,omitempty"`
	Shipping           string          `json:"shipping,omitempty"`
	CurrencySymbol     string          `json:"currency_symbol,omitempty"`
}

// Pdf returns the underlying *fpdf.Fpdf used to build document
func (doc *Document) Pdf() *fpdf.Fpdf {
	return doc.pdf
}

// SetUnicodeTranslator to use
// See https://pkg.go.dev/github.com/go-pdf/fpdf#UnicodeTranslator
func (doc *Document) SetUnicodeTranslator(fn UnicodeTranslateFunc) {
	doc.Options.UnicodeTranslateFunc = fn
}

// encodeString encodes the string using doc.Options.UnicodeTranslateFunc
func (doc *Document) encodeString(str string) string {
	return doc.Options.UnicodeTranslateFunc(str)
}

// typeAsString return the document type as string
func (d *Document) typeAsString() string {
	if d.Type == Invoice {
		return d.Options.TextTypeInvoice
	}

	if d.Type == Quotation {
		return d.Options.TextTypeQuotation
	}

	return d.Options.TextTypeDeliveryNote
}
