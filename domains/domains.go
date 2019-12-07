package domains

import (
	"encoding/json"

	"github.com/alyx/godaddy"
)

// List returns a list of DomainSummary objects for all domains owned
// by the user.
func List(c *godaddy.Client) ([]DomainSummary, error) {
	var domains []DomainSummary

	data, err := c.Get("/v1/domains")
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &domains)
	if err != nil {
		return nil, err
	}

	return domains, nil
}

func GetAgreements() { return }

func GetAvailable() { return }

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

func ValidateContactSchema()       { return }
func GetPurchaseSchema(tld string) { return }

func ValidatePurchaseSchema() { return }

func GetSuggestions() { return }

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
