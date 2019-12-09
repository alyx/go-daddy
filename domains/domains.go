package domains

import (
	"encoding/json"

	"github.com/alyx/godaddy"
	"github.com/alyx/godaddy/agreements"
)

// List returns a list of DomainSummary objects for all domains owned
// by the user.
func List(c *godaddy.Client, statuses []string, statusGroups []string, limit int,
	marker string, includes []string, modifiedDate string) ([]DomainSummary, error) {
	var res []DomainSummary

	uri, err := godaddy.BuildQuery("/v1/domains", map[string]interface{}{
		"statuses":     statuses,
		"statusGroups": statusGroups,
		"limit":        limit,
		"marker":       marker,
		"includes":     includes,
		"modifiedDate": modifiedDate,
	})
	if err != nil {
		return res, err
	}

	data, err := c.Get(uri)
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(data, &res)
	if err != nil {
		return res, err
	}

	return res, nil
}

// GetAgreements retrieves the legal agreement(s) required to purchase the specified TLD and add-ons
func GetAgreements(c *godaddy.Client, tlds []string, privacy bool, forTransfer bool) ([]agreements.LegalAgreement, error) {
	var res []agreements.LegalAgreement

	uri, err := godaddy.BuildQuery("/v1/domains/agreements", map[string]interface{}{
		"tlds":        tlds,
		"privacy":     privacy,
		"forTransfer": forTransfer,
	})
	if err != nil {
		return res, err
	}

	data, err := c.Get(uri)
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(data, &res)
	if err != nil {
		return res, err
	}

	return res, nil
}

// GetAvailable determines whether or not the specified domain is available for purchase
func GetAvailable(c *godaddy.Client, domain string, checkType string, forTransfer bool) (*DomainAvailableResponse, error) {
	res := new(DomainAvailableResponse)

	uri, err := godaddy.BuildQuery("/v1/domains/available", map[string]interface{}{
		"domain":      domain,
		"checkType":   checkType,
		"forTransfer": forTransfer,
	})
	if err != nil {
		return res, err
	}

	data, err := c.Get(uri)
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(data, &res)
	if err != nil {
		return res, err
	}

	return res, nil
}

// GetAvailableBulk determine whether or not the specified domains are
// available for purchase
func GetAvailableBulk(c *godaddy.Client, domains []string) (*DomainAvailableBulkMixed, error) {
	var res = new(DomainAvailableBulkMixed)

	converted, err := json.Marshal(domains)
	if err != nil {
		return &DomainAvailableBulkMixed{}, err
	}

	data, err := c.Post("/v1/domains/available", converted)
	if err != nil {
		return &DomainAvailableBulkMixed{}, err
	}

	err = json.Unmarshal(data, &res)
	if err != nil {
		return &DomainAvailableBulkMixed{}, err
	}

	return res, nil
}

// ValidateContactSchema validates the request body using the Domain Contact Validation Schema for specified domains.
func ValidateContactSchema(c *godaddy.Client, marketID string, body *DomainContactsBulk) (bool, error) {
	uri, err := godaddy.BuildQuery("/v1/domains/contacts/validate", map[string]interface{}{
		"marketId": marketID,
	})
	if err != nil {
		return false, err
	}

	enc, err := json.Marshal(body)
	if err != nil {
		return false, err
	}

	_, err = c.Post(uri, enc)
	if err != nil {
		return false, err
	}

	return true, nil
}

func GetPurchaseSchema(tld string) { return }

func ValidatePurchaseSchema() { return }

// GetSuggestions suggests alternate Domain names based on a seed Domain,
// a set of keywords, or the shopper's purchase history
func GetSuggestions(c *godaddy.Client, query string, country string, city string,
	sources []string, tlds []string, lengthMax int, lengthMin int, limit int,
	waitMs int) ([]DomainSuggestion, error) {
	var suggestions []DomainSuggestion

	uri, err := godaddy.BuildQuery("/v1/domains/suggest", map[string]interface{}{
		"query":     query,
		"country":   country,
		"city":      city,
		"sources":   sources,
		"tlds":      tlds,
		"lengthMax": lengthMax,
		"lengthMin": lengthMin,
		"limit":     limit,
		"waitMs":    waitMs,
	})
	if err != nil {
		return nil, err
	}

	data, err := c.Get(uri)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &suggestions)
	if err != nil {
		return nil, err
	}

	return suggestions, nil
}

// GetTLDs retrieves a list of TLDs supported and enabled for sale
func GetTLDs(c *godaddy.Client) ([]TldSummary, error) {
	var tlds []TldSummary

	data, err := c.Get("/v1/domains/tlds")
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &tlds)
	if err != nil {
		return nil, err
	}

	return tlds, nil
}

// Individual domain settings

func Purchase()       { return }
func Delete()         { return }
func Get()            { return }
func Update()         { return }
func UpdateContacts() { return }

// Privacy
func DeletePrivacy()   { return }
func PurchasePrivacy() { return }

func AddRecords()                  { return }
func ReplaceRecords()              { return }
func ReplaceRecordsByType()        { return }
func ReplaceRecordsByTypeAndName() { return }
func GetRecords()                  { return }

func Renew()               { return }
func Transfer()            { return }
func SendRegistrantEmail() { return }

// v2

func RemoveForward() { return }
func GetForward()    { return }
func UpdateForward() { return }
