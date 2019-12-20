// Copyright 2019 A. Wolcott. All rights reserved.
//
// Use of this source code is governed by the ISC
// license that can be found in the LICENSE file.

package daddy

// Shopper represents a customer, internally known as a shopper
type Shopper struct {
	ShopperID
	ShopperUpdate
}

// ShopperID represents the core IDs, customer ID and shopper ID, of a customer
type ShopperID struct {
	CustomerID string
	ShopperID  string
}

// ShopperStatus shows the current billing state of the customer
type ShopperStatus struct {
	BillingState string
}

// ShopperUpdate holds information used to update a customer's accoutn
type ShopperUpdate struct {
	Email      string
	ExternalID int
	MarketID   string
	NameFirst  string
	NameLast   string
}

// SubaccountCreate holds information used for creating a new customer subaccount
type SubaccountCreate struct {
	ShopperUpdate
	Password string
}

// Secret holds a new password
type Secret struct {
	Secret string
}
