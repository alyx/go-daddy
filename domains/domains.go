package domains

import (
	"encoding/json"
	"net/url"
	"strconv"
	"strings"

	"github.com/alyx/godaddy"
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

// GetSuggestions suggests alternate Domain names based on a seed Domain,
// a set of keywords, or the shopper's purchase history
func GetSuggestions(c *godaddy.Client, query string, country string, city string,
	sources []string, tlds []string, lengthMax int, lengthMin int, limit int,
	waitMs int) ([]DomainSuggestion, error) {
	var suggestions []DomainSuggestion
	u, err := url.Parse("/v1/domains/suggest")
	if err != nil {
		return nil, err
	}
	q, err := url.ParseQuery(u.RawQuery)
	if err != nil {
		return nil, err
	}

	if query != "" {
		q.Add("query", query)
	}
	if country != "" {
		q.Add("country", country)
	}
	if city != "" {
		q.Add("city", city)
	}
	if len(sources) > 0 {
		q.Add("sources", strings.Join(sources, ","))
	}
	if len(tlds) > 0 {
		q.Add("tlds", strings.Join(tlds, ","))
	}
	if lengthMax > 0 {
		q.Add("lengthMax", strconv.Itoa(lengthMax))
	}
	if lengthMin > 0 {
		q.Add("lengthMin", strconv.Itoa(lengthMin))
	}
	if limit > 0 {
		q.Add("limit", strconv.Itoa(limit))
	}
	if waitMs > 0 {
		q.Add("waitMs", strconv.Itoa(waitMs))
	}

	u.RawQuery = q.Encode()

	data, err := c.Get(u.String())
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
