// Copyright 2019 A. Wolcott. All rights reserved.
//
// Use of this source code is governed by the ISC
// license that can be found in the LICENSE file.

package daddy

import (
	"encoding/json"
)

type CountriesService service

// Get retrieves summary country information for the provided marketID and
// filters.
func (s *CountriesService) Get(marketID string, regionTypeID int, regionName string, sort string, order string) ([]CountrySummary, error) {
	var res []CountrySummary
	uri, err := BuildQuery("/v1/countries", map[string]interface{}{
		"marketId":     marketID,
		"regionTypeId": regionTypeID,
		"regionName":   regionName,
		"sort":         sort,
		"order":        order,
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

// GetByKey retrieves country and summary state information for provided
// countryKey
func (s *CountriesService) GetByKey(countryKey string, marketID string, sort string, order string) ([]Country, error) {
	var res []Country
	uri, err := BuildQuery("/v1/countries/"+countryKey, map[string]interface{}{
		"marketId": marketID,
		"sort":     sort,
		"order":    order,
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
