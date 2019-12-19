// Copyright 2019 A. Wolcott. All rights reserved.
//
// Use of this source code is governed by the ISC
// license that can be found in the LICENSE file.

package daddy

import (
	"encoding/json"
)

type OrdersService service

// List retrieves a list of orders for the authenticated shopper.
func (s *OrdersService) List(periodStart string, periodEnd string, domain string,
	productGroupID int, paymentProfileID int, parentOrderID string, offset int,
	limit int, sort string) (*OrderList, error) {
	res := new(OrderList)
	uri, err := BuildQuery("/v1/orders", map[string]interface{}{
		"periodStart":      periodStart,
		"periodEnd":        periodEnd,
		"domain":           domain,
		"productGroupId":   productGroupID,
		"paymentProfileId": paymentProfileID,
		"parentOrderId":    parentOrderID,
		"offset":           offset,
		"limit":            limit,
		"sort":             sort,
	})
	if err != nil {
		return res, err
	}

	data, err := s.client.Get(uri)
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(data, &res)

	return res, err
}

// Get retrieves details for the specified order.
func (s *OrdersService) Get(orderID string) (*Order, error) {
	res := new(Order)

	data, err := s.client.Get("/v1/orders/" + orderID)
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(data, &res)

	return res, err
}
