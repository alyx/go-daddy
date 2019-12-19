// Copyright 2019 A. Wolcott. All rights reserved.
//
// Use of this source code is governed by the ISC
// license that can be found in the LICENSE file.

package daddy

type Shopper struct {
	ShopperID
	ShopperUpdate
}

type ShopperID struct {
	CustomerID string
	ShopperID  string
}

type ShopperStatus struct {
	BillingState string
}

type ShopperUpdate struct {
	Email      string
	ExternalID int
	MarketID   string
	NameFirst  string
	NameLast   string
}

type SubaccountCreate struct {
	ShopperUpdate
	Password string
}

type Secret struct {
	Secret string
}
