// Copyright 2019 A. Wolcott. All rights reserved.
//
// Use of this source code is governed by the ISC
// license that can be found in the LICENSE file.

package daddy

// Client is the core wrapper around the GoDaddy API client.
type Client struct {
	PrivateLabel string
	Market       string
	Shopper      string
	Key          string
	Secret       string
	BaseURL      string

	common service
	// Services used for talking to different parts of the GoDaddy API
	Abuse         *AbuseService
	Aftermarket   *AftermarketService
	Agreements    *AgreementsService
	Certificates  *CertificatesService
	Countries     *CountriesService
	Domains       *DomainsService
	Orders        *OrdersService
	Shoppers      *ShoppersService
	Subscriptions *SubscriptionsService
}

type service struct {
	client *Client
}

// Error presents a generic error class
type Error struct {
	Code    string
	Fields  []ErrorField
	Message string
}

// ErrorField components are used to specify precise errors from an Error
type ErrorField struct {
	Code        string
	Message     string
	Path        string
	PathRelated string
}

// ErrorLimit is presented specifically when errors occur due to rate limiting
type ErrorLimit struct {
	Code          string
	Fields        []ErrorField
	Message       string
	RetryAfterSec int
}

// Address represents a billing/mailing address for an entity
type Address struct {
	Address1   string
	Address2   string
	City       string
	Country    string
	PostalCode string
	State      string
}

// Contact represents a domain registrant/contact/customer/etc.
type Contact struct {
	AddressMailing Address
	Email          string
	Fax            string
	JobTitle       string
	NameFirst      string
	NameLast       string
	NameMiddle     string
	Organization   string
	Phone          string
}

// Pagination represents the list of pages in a multi-page response
type Pagination struct {
	First    string
	Last     string
	Next     string
	Previous string
	Total    int
}
