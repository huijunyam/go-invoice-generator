package generator

const (
	// Invoice define the "invoice" document type
	Invoice string = "INVOICE"

	// Quotation define the "quotation" document type
	Quotation string = "QUOTATION"

	// DeliveryNote define the "delievry note" document type
	DeliveryNote string = "DELIVERY_NOTE"

	// BaseMargin define base margin used in documents
	BaseMargin float64 = 10

	// BaseMarginTop define base margin top used in documents
	BaseMarginTop float64 = 20

	// HeaderMarginTop define base header margin top used in documents
	HeaderMarginTop float64 = 5

	// MaxPageHeight define the maximum height for a single page
	MaxPageHeight float64 = 260
)

// Cols offsets
const (
	// ItemColNameOffset ...
	ItemColNameOffset float64 = 10

	// ItemColSKUOffset ...
	ItemColSKUOffset float64 = 80

	// ItemColQuantityOffset ...
	ItemColQuantityOffset float64 = 110

	// ItemColUnitPriceOffset ...
	ItemColUnitPriceOffset float64 = 140

	// // ItemColTotalHTOffset ...
	// ItemColTotalHTOffset float64 = 113

	// // ItemColDiscountOffset ...
	// ItemColDiscountOffset float64 = 140

	// // ItemColTaxOffset ...
	// ItemColTaxOffset float64 = 157

	// ItemColTotalTTCOffset ...
	ItemColTotalTTCOffset float64 = 175
)

const (
	// ItemColNameOffset ...
	ItemShippingOrderColNameOffset float64 = 10

	// ItemColSKUOffset ...
	ItemShippingOrderColSKUOffset float64 = 80

	// ItemColQuantityOffset ...
	ItemShippingOrderColQuantityOffset float64 = 120
)

const (
	// ItemColNameOffset ...
	ItemQuoteColNameOffset float64 = 10

	// ItemColSKUOffset ...
	ItemQuoteColSKUOffset float64 = 80

	// ItemColMinQuantityOffset ...
	ItemQuoteColMinQuantityOffset float64 = 110

	// ItemColQuantityOffset ...
	ItemQuoteColQuantityOffset float64 = 135

	// ItemColUnitPriceOffset ...
	ItemQuoteColUnitPriceOffset float64 = 155

	// // ItemColTotalHTOffset ...
	// ItemColTotalHTOffset float64 = 113

	// // ItemColDiscountOffset ...
	// ItemColDiscountOffset float64 = 140

	// // ItemColTaxOffset ...
	// ItemColTaxOffset float64 = 157

	// ItemColTotalTTCOffset ...
	ItemQuoteColTotalTTCOffset float64 = 175
)

var (
	// BaseTextFontSize define the base font size for text in document
	BaseTextFontSize float64 = 8

	// SmallTextFontSize define the small font size for text in document
	SmallTextFontSize float64 = 7

	// ExtraSmallTextFontSize define the extra small font size for text in document
	ExtraSmallTextFontSize float64 = 6

	// LargeTextFontSize define the large font size for text in document
	LargeTextFontSize float64 = 10
)
