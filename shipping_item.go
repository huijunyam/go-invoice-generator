package generator

// Item represent a 'product' or a 'service'
type ShippingItem struct {
	Name     string `json:"name,omitempty" validate:"required"`
	SKU      string `json:"sku,required"`
	Quantity string `json:"quantity,omitempty"`
}

// Prepare convert strings to decimal
func (i *ShippingItem) Prepare() error {
	return nil
}

// appendColTo document doc
func (i *ShippingItem) appendColTo(options *Options, doc *Document) {
	// Get base Y (top of line)
	baseY := doc.pdf.GetY()

	// Name
	doc.pdf.SetX(ItemShippingOrderColNameOffset)
	doc.pdf.MultiCell(
		ItemShippingOrderColSKUOffset-ItemShippingOrderColNameOffset,
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
	doc.pdf.SetX(ItemShippingOrderColSKUOffset)
	doc.pdf.MultiCell(
		ItemShippingOrderColQuantityOffset-ItemShippingOrderColSKUOffset,
		3,
		doc.encodeString(i.SKU),
		"",
		"",
		false,
	)

	// Quantity
	doc.pdf.SetY(baseY)
	doc.pdf.SetX(ItemShippingOrderColQuantityOffset)
	doc.pdf.CellFormat(
		190-ItemShippingOrderColQuantityOffset,
		colHeight,
		doc.encodeString(i.Quantity),
		"0",
		0,
		"",
		false,
		0,
		"",
	)
	// Set Y for next line
	doc.pdf.SetY(baseY + colHeight)
}
