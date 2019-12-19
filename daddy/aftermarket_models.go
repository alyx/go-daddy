// Copyright 2019 A. Wolcott. All rights reserved.
//
// Use of this source code is governed by the ISC
// license that can be found in the LICENSE file.

package daddy

// ListingAction contains the resulting action ID from performing a change
// to the Aftermarket listings
type AftermarketListingAction struct {
	ListingActionID int
}

// ListingExpiryCreate handles the input structure for creating a new listing
// in the Aftermarket registry.
type AftermarketListingExpiryCreate struct {
	Domain            string
	ExpiresAt         string
	LosingRegistrarID int
	PageViewsMonthly  int
	RevenueMonthly    int
}
