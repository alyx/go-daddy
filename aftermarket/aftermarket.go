package aftermarket

import (
	"encoding/json"

	"github.com/alyx/godaddy"
)

// DeleteListings removes listings from GoDaddy Auction
func DeleteListings(c *godaddy.Client, domains []string) (*ListingAction, error) {
	var res = new(ListingAction)

	enc, err := json.Marshal(domains)
	if err != nil {
		return res, err
	}

	data, err := c.Delete("/v1/aftermarket/listings", enc)
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(data, &res)

	return res, err
}

// AddListings adds expiry listings into GoDaddy Auction
func AddListings(c *godaddy.Client, listings []ListingExpiryCreate) (*ListingAction, error) {
	var res = new(ListingAction)

	enc, err := json.Marshal(listings)
	if err != nil {
		return res, err
	}

	data, err := c.Post("/v1/aftermarket/listings/expiry", enc)
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(data, &res)

	return res, err
}
