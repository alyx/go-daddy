package shoppers

import (
	"encoding/json"

	godaddy "github.com/alyx/go-daddy"
)

// Create a Subaccount owned by the authenticated Reseller
func Create(c *godaddy.Client, body *SubaccountCreate) (*ShopperID, error) {
	res := new(ShopperID)

	enc, err := json.Marshal(body)
	if err != nil {
		return res, err
	}

	data, err := c.Post("/v1/shoppers/subaccount", enc)
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(data, &res)

	return res, err
}

// Get details for the specified Shopper
func Get(c *godaddy.Client, shopper string, includes []string) (*Shopper, error) {
	res := new(Shopper)

	uri, err := godaddy.BuildQuery("/v1/shoppers/"+shopper, map[string]interface{}{
		"includes": includes,
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

// Update details for the specified Shopper
func Update(c *godaddy.Client, shopper string, body *ShopperUpdate) (*ShopperID, error) {
	res := new(ShopperID)

	enc, err := json.Marshal(body)
	if err != nil {
		return res, err
	}

	data, err := c.Post("/v1/shoppers/"+shopper, enc)
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(data, &res)

	return res, err
}

// Delete sends a request for deletion of a Shopper profile
func Delete(c *godaddy.Client, shopper string, auditClientIP string) error {
	uri, err := godaddy.BuildQuery("/v1/shoppers/"+shopper, map[string]interface{}{
		"auditClientIp": auditClientIP,
	})

	_, err = c.Delete(uri, nil)

	return err
}

// Status retrieves details on the specified Shopper
func Status(c *godaddy.Client, shopper string, auditClientIP string) (*ShopperStatus, error) {
	res := new(ShopperStatus)

	uri, err := godaddy.BuildQuery("/v1/shoppers/"+shopper+"/status", map[string]interface{}{
		"auditClientIp": auditClientIP,
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

// SetPassword sets password for a Subaccount
func SetPassword(c *godaddy.Client, shopper string, body *Secret) (*ShopperID, error) {
	res := new(ShopperID)

	enc, err := json.Marshal(body)
	if err != nil {
		return res, err
	}

	data, err := c.Put("/v1/shoppers/"+shopper+"/factors/password", enc)
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(data, &res)

	return res, err
}
