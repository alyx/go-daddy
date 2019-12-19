package subscriptions

import (
	"encoding/json"

	godaddy "github.com/alyx/go-daddy"
)

// List of Subscriptions for a specified Shopper
func List(c *godaddy.Client, productGroupKeys []string, includes []string, offset int, limit int, sort string) (*SubscriptionList, error) {
	res := new(SubscriptionList)

	uri, err := godaddy.BuildQuery("/v1/subscriptions", map[string]interface{}{
		"productGroupKeys": productGroupKeys,
		"includes":         includes,
		"offset":           offset,
		"limit":            limit,
		"sort":             sort,
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

// ListProductGroups retrieves a list of Product Groups for a specific Shopper
func ListProductGroups(c *godaddy.Client) ([]ProductGroup, error) {
	var res []ProductGroup

	data, err := c.Get("/v1/subscriptions/productGroups")
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
func Delete(c *godaddy.Client, id string) error {
	_, err := c.Delete("/v1/subscriptions/"+id, nil)

	return err
}

// Get will retrieve details for the specified Subscription
func Get(c *godaddy.Client, id string) (*Subscription, error) {
	res := new(Subscription)

	data, err := c.Get("/v1/subscriptions/" + id)
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
func Update(c *godaddy.Client, id string, body *SubscriptionUpdate) error {
	enc, err := json.Marshal(body)
	if err != nil {
		return err
	}

	_, err = c.Patch("/v1/subscriptions/"+id, enc)

	return err
}
