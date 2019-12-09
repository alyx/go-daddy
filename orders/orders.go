package orders

import (
	"encoding/json"
	"strconv"

	"github.com/alyx/godaddy"
)

// List retrieves a list of orders for the authenticated shopper.
func List(c *godaddy.Client, periodStart string, periodEnd string, domain string,
	productGroupID int, paymentProfileID int, parentOrderID string, offset int,
	limit int, sort string) (*OrderList, error) {
	res := new(OrderList)
	uri, err := godaddy.BuildQuery("/v1/orders", map[string]string{
		"periodStart":      periodStart,
		"periodEnd":        periodEnd,
		"domain":           domain,
		"productGroupId":   strconv.Itoa(productGroupID),
		"paymentProfileId": strconv.Itoa(paymentProfileID),
		"parentOrderId":    parentOrderID,
		"offset":           strconv.Itoa(offset),
		"limit":            strconv.Itoa(limit),
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

	return res, err
}

// Get retrieves details for the specified order.
func Get(c *godaddy.Client, orderID string) (*Order, error) {
	res := new(Order)

	data, err := c.Get("/v1/orders/" + orderID)
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(data, &res)

	return res, err
}
