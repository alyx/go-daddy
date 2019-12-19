// Copyright 2019 A. Wolcott. All rights reserved.
//
// Use of this source code is governed by the ISC
// license that can be found in the LICENSE file.

package daddy

import (
	"encoding/json"
)

type AgreementsService service

// Get retrieves Legal Agreements for provided agreements keys
func (s *AgreementsService) Get(keys []string) ([]LegalAgreement, error) {
	var res []LegalAgreement

	uri, err := BuildQuery("/v1/agreements", map[string]interface{}{
		"keys": keys,
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
