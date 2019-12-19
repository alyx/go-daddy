package agreements

import (
	"encoding/json"

	godaddy "github.com/alyx/go-daddy"
)

// Get retrieves Legal Agreements for provided agreements keys
func Get(c *godaddy.Client, keys []string) ([]LegalAgreement, error) {
	var res []LegalAgreement

	uri, err := godaddy.BuildQuery("/v1/agreements", map[string]interface{}{
		"keys": keys,
	})
	if err != nil {
		return res, err
	}

	data, err := c.Get(uri)
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(data, &res)

	return res, err
}
