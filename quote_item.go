package generator

import (
	"fmt"

	"github.com/shopspring/decimal"
)

// Item represent a 'product' or a 'service'
type QuoteItem struct {
	Name            string            `json:"name,omitempty" validate:"required"`
	SKU             string            `json:"sku,required"`
	UnitCost        string            `json:"unit_cost,omitempty"`
	Quantity        string            `json:"quantity,omitempty"`
	MinQuantity     string            `json:"min_quantity,omitempty"`
	Total           string            `json:"total,omitempty"`
	PriceBreakItems []*PriceBreakItem `json:"price_break_items,omitempty"`
	// Tax         *Tax      `json:"tax,omitempty"`
	// Discount    *Discount `json:"discount,omitempty"`
	CurrencySymbol string `json:"currency_symbol,omitempty"`
	_unitCost      decimal.Decimal
	_quantity      decimal.Decimal
}

type PriceBreakItem struct {
	Quantity string `json:"quantity,omitempty"`
	UnitCost string `json:"unit_cost,omitempty"`
	Total    string `json:"total,omitempty"`
}

// Prepare convert strings to decimal
func (i *QuoteItem) Prepare() error {
	return nil
}

// appendColTo document doc
func (i *QuoteItem) appendColTo(options *Options, doc *Document) {
	// Get base Y (top of line)
	baseY := doc.pdf.GetY()

	// Name
	doc.pdf.SetX(ItemQuoteColNameOffset)
	doc.pdf.MultiCell(
		ItemQuoteColSKUOffset-ItemQuoteColNameOffset,
		3,
		doc.encodeString(i.Name),
		"",
		"",
		false,
	)

	// Compute line height
	colHeight := doc.pdf.GetY() - baseY

	// sku
	doc.pdf.SetY(baseY)
	doc.pdf.SetX(ItemQuoteColSKUOffset)
	doc.pdf.MultiCell(
		ItemQuoteColMinQuantityOffset-ItemQuoteColSKUOffset,
		3,
		doc.encodeString(i.SKU),
		"",
		"",
		false,
	)

	// Min quantity
	doc.pdf.SetY(baseY)
	doc.pdf.SetX(ItemQuoteColMinQuantityOffset)
	doc.pdf.CellFormat(
		ItemQuoteColQuantityOffset-ItemQuoteColMinQuantityOffset,
		colHeight,
		doc.encodeString(i.MinQuantity),
		"0",
		0,
		"",
		false,
		0,
		"",
	)

	// Quantity
	doc.pdf.SetY(baseY)
	doc.pdf.SetX(ItemQuoteColQuantityOffset)
	doc.pdf.CellFormat(
		ItemQuoteColUnitPriceOffset-ItemQuoteColQuantityOffset,
		colHeight,
		doc.encodeString(i.Quantity),
		"0",
		0,
		"",
		false,
		0,
		"",
	)

	// Unit price
	doc.pdf.SetY(baseY)
	doc.pdf.SetX(ItemQuoteColUnitPriceOffset)
	doc.pdf.CellFormat(
		ItemQuoteColTotalTTCOffset-ItemQuoteColUnitPriceOffset,
		colHeight,
		doc.encodeString(fmt.Sprintf("%s %s", i.CurrencySymbol, i.UnitCost)),
		"0",
		0,
		"",
		false,
		0,
		"",
	)

	// TOTAL TTC
	doc.pdf.SetX(ItemColTotalTTCOffset)
	doc.pdf.CellFormat(
		190-ItemColTotalTTCOffset,
		colHeight,
		doc.encodeString(fmt.Sprintf("%s %s", i.CurrencySymbol, i.Total)),
		"0",
		0,
		"",
		false,
		0,
		"",
	)

	// Set Y for next line
	doc.pdf.SetY(baseY + colHeight)

	if len(i.PriceBreakItems) > 0 {
		for _, priceBreakItem := range i.PriceBreakItems {
			addPriceBreakItem(i, priceBreakItem, options, doc, baseY, colHeight)
		}
	}
}

func addPriceBreakItem(i *QuoteItem, priceBreakItem *PriceBreakItem, options *Options, doc *Document, baseY float64, colHeight float64) {
	// Quantity
	doc.pdf.SetY(baseY)
	doc.pdf.SetX(ItemQuoteColQuantityOffset)
	doc.pdf.CellFormat(
		ItemQuoteColUnitPriceOffset-ItemQuoteColQuantityOffset,
		colHeight,
		doc.encodeString(priceBreakItem.Quantity),
		"0",
		0,
		"",
		false,
		0,
		"",
	)

	// Unit price
	doc.pdf.SetY(baseY)
	doc.pdf.SetX(ItemQuoteColUnitPriceOffset)
	doc.pdf.CellFormat(
		ItemQuoteColTotalTTCOffset-ItemQuoteColUnitPriceOffset,
		colHeight,
		doc.encodeString(fmt.Sprintf("%s %s", i.CurrencySymbol, priceBreakItem.UnitCost)),
		"0",
		0,
		"",
		false,
		0,
		"",
	)

	// TOTAL TTC
	doc.pdf.SetX(ItemColTotalTTCOffset)
	doc.pdf.CellFormat(
		190-ItemColTotalTTCOffset,
		colHeight,
		doc.encodeString(fmt.Sprintf("%s %s", i.CurrencySymbol, priceBreakItem.Total)),
		"0",
		0,
		"",
		false,
		0,
		"",
	)

	doc.pdf.SetY(baseY + colHeight)
}
