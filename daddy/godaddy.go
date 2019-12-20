// Copyright 2019 A. Wolcott. All rights reserved.
//
// Use of this source code is governed by the ISC
// license that can be found in the LICENSE file.

package daddy

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func (e *Error) Error() string {
	return e.Message
}

// GenericWithBody is handler for HTTP actions with potential body content
func (c *Client) GenericWithBody(action string, method string, body []byte) ([]byte, error) {
	client := new(http.Client)
	req, err := http.NewRequest(action, c.BaseURL+method, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("sso-key %s:%s", c.Key, c.Secret))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	if c.Shopper != "" {
		req.Header.Add("X-Shopper-Id", c.Shopper)
	}
	if c.Market != "" {
		req.Header.Add("X-Market-Id", c.Market)
	}
	if c.PrivateLabel != "" {
		req.Header.Add("X-Private-Label-Id", c.PrivateLabel)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode > 299 {
		e := new(Error)

		err = json.Unmarshal(data, &e)
		if err != nil {
			return nil, err
		}

		return data, e
	}

	return data, nil
}

// Get processes core API integration via HTTP GET requests
func (c *Client) Get(method string) ([]byte, error) {
	client := new(http.Client)
	req, err := http.NewRequest("GET", c.BaseURL+method, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("sso-key %s:%s", c.Key, c.Secret))
	req.Header.Add("Content-Type", "application/json")

	if c.Shopper != "" {
		req.Header.Add("X-Shopper-Id", c.Shopper)
	}
	if c.Market != "" {
		req.Header.Add("X-Market-Id", c.Market)
	}
	if c.PrivateLabel != "" {
		req.Header.Add("X-Private-Label-Id", c.PrivateLabel)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode > 299 {
		e := new(Error)

		err = json.Unmarshal(data, &e)
		if err != nil {
			return nil, err
		}

		return data, e
	}

	return data, nil
}

// Post processes core API integration via HTTP POST requests
func (c *Client) Post(method string, body []byte) ([]byte, error) {
	return c.GenericWithBody("POST", method, body)
}

// Delete processes core API integration via HTTP DELETE requests
func (c *Client) Delete(method string, body []byte) ([]byte, error) {
	return c.GenericWithBody("DELETE", method, body)
}

// Put processes core API integration via HTTP PUT requests
func (c *Client) Put(method string, body []byte) ([]byte, error) {
	return c.GenericWithBody("PUT", method, body)
}

// Patch processes core API integration via HTTP POST requests
func (c *Client) Patch(method string, body []byte) ([]byte, error) {
	return c.GenericWithBody("PATCH", method, body)
}

// BuildQuery builds an HTTP query from a map of string:string components
func BuildQuery(uri string, values map[string]interface{}) (string, error) {
	u, err := url.Parse(uri)
	if err != nil {
		return "", err
	}

	q, err := url.ParseQuery(u.RawQuery)
	if err != nil {
		return "", err
	}

	for key, value := range values {
		switch v := value.(type) {
		case int:
			if v > 0 {
				q.Add(key, strconv.Itoa(v))
			}
		case string:
			if v != "" {
				q.Add(key, v)
			}
		case []string:
			if len(v) > 0 {
				q.Add(key, strings.Join(v, ","))
			}
		case bool:
			q.Add(key, strconv.FormatBool(v))
		}
	}

	u.RawQuery = q.Encode()

	return u.String(), nil
}

// NewClient accepts three inputs:
// - key is a string containing your GoDaddy API key. If this value is empty,
//   NewClient will attempt to find GODADDY_API_KEY in your environmental
//   variables.
// - secret is a string containing your GoDaddy API secret. IF this value is
//   empty, NewClient will attept to find GODADDY_API_SECRET in your
//   environmental variables.
// - ote is a bool, true runs your API calls against the GoDaddy test
//   environment at https://api.ote-godaddy.com/, false runs your API
//   calls against the production environment, https://api.godaddy.com/
//
// NewClient will return a godaddy.Client object on success, and an error if
// it cannot determine either your api key or secret key.
//
//XXX: Update for preferring environment.
func NewClient(key string, secret string, ote bool) (*Client, error) {
	if ote {
		return NewClientWithURL(key, secret, "https://api.ote-godaddy.com")
	}

	return NewClientWithURL(key, secret, "https://api.godaddy.com")
}

// NewClientWithURL functions similarly to NewClient, with the exception of
// accepting any URL as the endpoint as opposed to pre-setting the endpoint URL
// based on official GoDaddy endpoints.
func NewClientWithURL(key string, secret string, url string) (*Client, error) {
	c := new(Client)

	if key == "" {
		return nil, errors.New("Missing GoDaddy API key")
	}
	c.Key = key

	if secret == "" {
		return nil, errors.New("Missing GoDaddy API secret")
	}
	c.Secret = secret

	if url == "" {
		return nil, errors.New("Missing API endpoint")
	}
	c.BaseURL = url

	c.common.client = c
	c.Abuse = (*AbuseService)(&c.common)
	c.Aftermarket = (*AftermarketService)(&c.common)
	c.Agreements = (*AgreementsService)(&c.common)
	c.Certificates = (*CertificatesService)(&c.common)
	c.Countries = (*CountriesService)(&c.common)
	c.Domains = (*DomainsService)(&c.common)
	c.Orders = (*OrdersService)(&c.common)
	c.Shoppers = (*ShoppersService)(&c.common)
	c.Subscriptions = (*SubscriptionsService)(&c.common)

	return c, nil
}
