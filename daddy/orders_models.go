// Copyright 2019 A. Wolcott. All rights reserved.
//
// Use of this source code is governed by the ISC
// license that can be found in the LICENSE file.

package daddy

type BillTo struct {
	Contact Contact
	TaxID   string
}

type LineItem struct {
	Domains      []string
	Period       int
	PeriodUnit   string
	Pfid         int
	Pricing      LineItemPricing
	Quantity     int
	TaxCollector LineItemTaxCollector
}

type LineItemPricing struct {
	Discount int
	Fees     OrderFee
	List     int
	Sale     int
	Savings  int
	Subtotal int
	Taxes    int
	Unit     int //XXX: Is this right
}

type LineItemSummary struct {
	Label string
}

type LineItemTaxCollector struct {
	TaxCollectorID int
}

type LineItemPricingTaxDetail struct {
	Amount int
	Rate   int
}

type LineItemUnitPricing struct {
	Discount int
	Fees     OrderFee
	List     int
	Sale     int
	Savings  int
	Taxes    int
}

type Order struct {
	BillTo        BillTo
	CreatedAt     string
	Currency      string
	Items         []LineItem
	ParentOrderID string
	Payments      []OrderPayment
	Pricing       OrderPricing
}

type OrderFee struct {
	ICANN int
	Total int
}

type OrderList struct {
	Orders     []Order
	Pagination Pagination
}

type OrderPricing struct {
	Discount   int
	Fees       OrderFee
	ID         int
	List       int
	Savings    int
	Subtotal   int
	Taxes      int
	TaxDetails []LineItemPricingTaxDetail
	Total      int
}

type OrderSummary struct {
	CreatedAt     string
	Currency      string
	Items         []LineItemSummary
	OrderID       int
	ParentOrderID string
	Pricing       OrderSummaryPricing
}

type OrderSummaryPricing struct {
	Total string
}

type OrderPayment struct {
	Amount           int
	Category         string
	PaymentProfileID string
	Subcategory      string
}
