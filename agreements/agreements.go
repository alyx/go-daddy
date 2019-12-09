package agreements

import (
	"encoding/json"
	"strings"

	"github.com/alyx/godaddy"
)

// Get retrieves Legal Agreements for provided agreements keys
func Get(c *godaddy.Client, keys []string) ([]LegalAgreement, error) {
	var res []LegalAgreement

	uri, err := godaddy.BuildQuery("/v1/agreements", map[string]string{
		"keys": strings.Join(keys, ","),
	})
	if err != nil {
		return res, err
	}

	data, err := godaddy.Get(uri)
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(data, &res)

	return res, err
}
