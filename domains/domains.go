package domains

import (
	"encoding/json"

	godaddy "github.com/alyx/go-daddy"
	"github.com/alyx/go-daddy/agreements"
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

// Purchase and register the specified Domain
func Purchase(c *godaddy.Client, body *DomainPurchase) (*DomainPurchaseResponse, error) {
	res := new(DomainPurchaseResponse)

	enc, err := json.Marshal(body)
	if err != nil {
		return res, err
	}

	data, err := c.Post("/v1/domains/purchase", enc)
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(data, &res)

	return res, err
}

// Delete cancels a purchased domain
func Delete(c *godaddy.Client, domain string) error {
	_, err := c.Delete("/v1/domains/"+domain, nil)

	return err
}

// Get retrieves details for the specified Domain
func Get(c *godaddy.Client, domain string) (*DomainDetail, error) {
	res := new(DomainDetail)

	data, err := c.Get("/v1/domains/" + domain)
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(data, &res)

	return res, err
}

// Update details for the specified Domain
func Update(c *godaddy.Client, domain string, body *DomainUpdate) error {
	enc, err := json.Marshal(body)
	if err != nil {
		return err
	}

	_, err = c.Patch("/v1/domains/"+domain, enc)

	return err
}

// UpdateContacts updates contacts for the specified Domain
func UpdateContacts(c *godaddy.Client, domain string, body *DomainContacts) error {
	enc, err := json.Marshal(body)
	if err != nil {
		return err
	}

	_, err = c.Patch("/v1/domains/"+domain+"/contacts", enc)

	return err
}

// DeletePrivacy submits a privacy cancellation request for the given Domain
func DeletePrivacy(c *godaddy.Client, domain string) error {
	_, err := c.Delete("/v1/domains/"+domain+"/privacy", nil)

	return err
}

// PurchasePrivacy purchases privacy for a specified Domain
func PurchasePrivacy(c *godaddy.Client, domain string, body *PrivacyPurchase) (*DomainPurchaseResponse, error) {
	res := new(DomainPurchaseResponse)

	enc, err := json.Marshal(body)
	if err != nil {
		return res, err
	}

	data, err := c.Post("/v1/domains/"+domain+"/privacy/purchase", enc)
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(data, &res)

	return res, err
}

// AddRecords adds the specified DNS Records to the specified Domain
func AddRecords(c *godaddy.Client, domain string, body []DNSRecord) error {
	enc, err := json.Marshal(body)
	if err != nil {
		return err
	}

	_, err = c.Patch("/v1/domains/"+domain+"/records", enc)

	return err
}

// ReplaceRecords replaces all DNS Records for the specified Domain
func ReplaceRecords(c *godaddy.Client, domain string, body []DNSRecord) error {
	enc, err := json.Marshal(body)
	if err != nil {
		return err
	}

	_, err = c.Put("/v1/domains/"+domain+"/records", enc)

	return err
}

// ReplaceRecordsByType replaces all DNS Records for the specified Domain with
// the specified Type
func ReplaceRecordsByType(c *godaddy.Client, domain string, dnstype string, body []DNSRecord) error {
	enc, err := json.Marshal(body)
	if err != nil {
		return err
	}

	_, err = c.Put("/v1/domains/"+domain+"/records/"+dnstype, enc)

	return err
}

// ReplaceRecordsByTypeAndName replaces all DNS Records for the specified
// Domain with the specified Type and Name
func ReplaceRecordsByTypeAndName(c *godaddy.Client, domain string, dnstype string, name string, body []DNSRecord) error {
	enc, err := json.Marshal(body)
	if err != nil {
		return err
	}

	_, err = c.Put("/v1/domains/"+domain+"/records/"+dnstype+"/"+name, enc)

	return err
}

func GetRecords() { return }

// Renew purchases a renewal for the specified Domain
func Renew(c *godaddy.Client, domain string, body *DomainRenew) (*DomainPurchaseResponse, error) {
	res := new(DomainPurchaseResponse)

	enc, err := json.Marshal(body)
	if err != nil {
		return res, err
	}

	data, err := c.Post("/v1/domains/"+domain+"/renew", enc)
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(data, &res)

	return res, err
}

// Transfer purchases and starts or restarts a transfer process
func Transfer(c *godaddy.Client, domain string, body *DomainTransferIn) (*DomainPurchaseResponse, error) {
	res := new(DomainPurchaseResponse)

	enc, err := json.Marshal(body)
	if err != nil {
		return res, err
	}

	data, err := c.Post("/v1/domains/"+domain+"/transfer", enc)
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(data, &res)

	return res, err
}

// SendRegistrantEmail re-sends Contact E-Mail Verification for specified Domain
func SendRegistrantEmail(c *godaddy.Client, domain string) error {
	_, err := c.Post("/v1/domains/"+domain+"/verifyRegistrantEmail", nil)

	return err
}

// v2

// RemoveForward submits a forwarding cancellation request for the given fqdn
func RemoveForward(c *godaddy.Client, customer string, fqdn string) error {
	_, err := c.Delete("/v2/customers/"+customer+"/domains/forwards/"+fqdn, nil)

	return err
}

// GetForward retrieves the forwarding information for the given fqdn
func GetForward(c *godaddy.Client, customer string, fqdn string, subdomains bool) ([]DomainForwarding, error) {
	var res []DomainForwarding
	uri, err := godaddy.BuildQuery("/v2/customers/"+customer+"/domains/forwards/"+fqdn,
		map[string]interface{}{"includeSubs": subdomains})
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

// UpdateForward modifies the forwarding information for the given fqdn
func UpdateForward(c *godaddy.Client, customer string, fqdn string, body *DomainForwardingCreate) error {
	enc, err := json.Marshal(body)
	if err != nil {
		return err
	}

	_, err = c.Put("/v2/customers/"+customer+"/domains/forwards/"+fqdn, enc)

	return err
}
