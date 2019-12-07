package godaddy

// SPDX-License-Identifier: ISC

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// Client is the core wrapper around the GoDaddy API client.
type Client struct {
	Key    string
	Secret string
	OTE    bool
}

// Error presents a generic error class
type Error struct {
	Code    string
	Fields  []ErrorField
	Message string
}

func (e *Error) Error() string {
	return e.Message
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

// GetURL is a helper function returning the API base URL for all API calls.
// If the GODADDY_API_URL is defined, this value will override the standard
// options.
func (c *Client) GetURL() string {
	e := os.Getenv("GODADDY_API_URL")
	if e != "" {
		return e
	} else if c.OTE {
		return "https://api.ote-godaddy.com"
	}

	return "https://api.godaddy.com"
}

// Get processes core API integration via HTTP GET requests
func (c *Client) Get(method string) ([]byte, error) {
	client := new(http.Client)
	req, err := http.NewRequest("GET", c.GetURL()+method, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("sso-key %s:%s", c.Key, c.Secret))
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
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
	client := new(http.Client)
	req, err := http.NewRequest("POST", c.GetURL()+method, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("sso-key %s:%s", c.Key, c.Secret))
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		e := new(Error)

		err = json.Unmarshal(data, &e)
		if err != nil {
			return nil, err
		}

		return data, e
	}

	return data, nil
}

func checkEnvOrGiveString(env string, value string) string {
	v := os.Getenv(env)
	if v == "" {
		return value
	}

	return v
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
	client := new(Client)

	client.Key = checkEnvOrGiveString("GODADDY_API_KEY", key)
	if client.Key == "" {
		return nil, errors.New("Missing GoDaddy API key")
	}

	client.Secret = checkEnvOrGiveString("GODADDY_API_SECRET", secret)
	if client.Secret == "" {
		return nil, errors.New("Missing GoDaddy API secret key")
	}

	client.OTE = ote

	return client, nil
}
