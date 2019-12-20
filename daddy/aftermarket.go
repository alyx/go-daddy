// Copyright 2019 A. Wolcott. All rights reserved.
//
// Use of this source code is governed by the ISC
// license that can be found in the LICENSE file.

package daddy

import (
	"encoding/json"
)

// AftermarketService handles communication with the aftermarket related
// methods of the GoDaddy API.
//
// GoDaddy API docs: https://developer.godaddy.com/doc/endpoint/aftermarket
type AftermarketService service

// DeleteListings removes listings from GoDaddy Auction
func (s *AftermarketService) DeleteListings(domains []string) (*AftermarketListingAction, error) {
	var res = new(AftermarketListingAction)

	enc, err := json.Marshal(domains)
	if err != nil {
		return res, err
	}

	data, err := s.client.Delete("/v1/aftermarket/listings", enc)
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(data, &res)

	return res, err
}

// AddListings adds expiry listings into GoDaddy Auction
func (s *AftermarketService) AddListings(listings []AftermarketListingExpiryCreate) (*AftermarketListingAction, error) {
	var res = new(AftermarketListingAction)

	enc, err := json.Marshal(listings)
	if err != nil {
		return res, err
	}

	data, err := s.client.Post("/v1/aftermarket/listings/expiry", enc)
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(data, &res)

	return res, err
}
