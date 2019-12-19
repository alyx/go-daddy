// Copyright 2019 A. Wolcott. All rights reserved.
//
// Use of this source code is governed by the ISC
// license that can be found in the LICENSE file.

package daddy

import (
	"encoding/json"
)

type ShoppersService service

// Create a Subaccount owned by the authenticated Reseller
func (s *ShoppersService) Create(body *SubaccountCreate) (*ShopperID, error) {
	res := new(ShopperID)

	enc, err := json.Marshal(body)
	if err != nil {
		return res, err
	}

	data, err := s.client.Post("/v1/shoppers/subaccount", enc)
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(data, &res)

	return res, err
}

// Get details for the specified Shopper
func (s *ShoppersService) Get(shopper string, includes []string) (*Shopper, error) {
	res := new(Shopper)

	uri, err := BuildQuery("/v1/shoppers/"+shopper, map[string]interface{}{
		"includes": includes,
	})
	if err != nil {
		return res, err
	}

	data, err := s.client.Get(uri)
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(data, &res)
	if err != nil {
		return res, err
	}

	return res, nil
}

// Update details for the specified Shopper
func (s *ShoppersService) Update(shopper string, body *ShopperUpdate) (*ShopperID, error) {
	res := new(ShopperID)

	enc, err := json.Marshal(body)
	if err != nil {
		return res, err
	}

	data, err := s.client.Post("/v1/shoppers/"+shopper, enc)
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(data, &res)

	return res, err
}

// Delete sends a request for deletion of a Shopper profile
func (s *ShoppersService) Delete(shopper string, auditClientIP string) error {
	uri, err := BuildQuery("/v1/shoppers/"+shopper, map[string]interface{}{
		"auditClientIp": auditClientIP,
	})

	_, err = s.client.Delete(uri, nil)

	return err
}

// Status retrieves details on the specified Shopper
func (s *ShoppersService) Status(shopper string, auditClientIP string) (*ShopperStatus, error) {
	res := new(ShopperStatus)

	uri, err := BuildQuery("/v1/shoppers/"+shopper+"/status", map[string]interface{}{
		"auditClientIp": auditClientIP,
	})
	if err != nil {
		return res, err
	}

	data, err := s.client.Get(uri)
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(data, &res)
	if err != nil {
		return res, err
	}

	return res, nil
}

// SetPassword sets password for a Subaccount
func (s *ShoppersService) SetPassword(shopper string, body *Secret) (*ShopperID, error) {
	res := new(ShopperID)

	enc, err := json.Marshal(body)
	if err != nil {
		return res, err
	}

	data, err := s.client.Put("/v1/shoppers/"+shopper+"/factors/password", enc)
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(data, &res)

	return res, err
}
