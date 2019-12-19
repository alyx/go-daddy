// Copyright 2019 A. Wolcott. All rights reserved.
//
// Use of this source code is governed by the ISC
// license that can be found in the LICENSE file.

package daddy

import (
	"encoding/json"
)

type SubscriptionsService service

// List of Subscriptions for a specified Shopper
func (s *SubscriptionsService) List(productGroupKeys []string, includes []string, offset int, limit int, sort string) (*SubscriptionList, error) {
	res := new(SubscriptionList)

	uri, err := BuildQuery("/v1/subscriptions", map[string]interface{}{
		"productGroupKeys": productGroupKeys,
		"includes":         includes,
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
	if err != nil {
		return res, err
	}

	return res, nil
}

// ListProductGroups retrieves a list of Product Groups for a specific Shopper
func (s *SubscriptionsService) ListProductGroups() ([]ProductGroup, error) {
	var res []ProductGroup

	data, err := s.client.Get("/v1/subscriptions/productGroups")
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(data, &res)
	if err != nil {
		return res, err
	}

	return res, nil
}

// Delete will cancel the specified Subscription
func (s *SubscriptionsService) Delete(id string) error {
	_, err := s.client.Delete("/v1/subscriptions/"+id, nil)

	return err
}

// Get will retrieve details for the specified Subscription
func (s *SubscriptionsService) Get(id string) (*Subscription, error) {
	res := new(Subscription)

	data, err := s.client.Get("/v1/subscriptions/" + id)
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(data, &res)
	if err != nil {
		return res, err
	}

	return res, nil
}

// Update will update details for the specified Subscription
// func Update(c *godaddy.Client, id string, body *SubscriptionUpdate) error {
// 	enc, err := json.Marshal(body)
// 	if err != nil {
// 		return err
// 	}

// 	_, err = c.Patch("/v1/subscriptions/"+id, enc)

// 	return err
// }
