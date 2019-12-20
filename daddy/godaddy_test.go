// Copyright 2019 A. Wolcott. All rights reserved.
//
// Use of this source code is governed by the ISC
// license that can be found in the LICENSE file.

package daddy

import (
	"testing"
)

type TestNewClientDataIn struct {
	Key     string
	Secret  string
	OTE     bool
	BaseURL string
}

type TestNewClientDataOut struct {
	Key     string
	Secret  string
	BaseURL string
	Error   bool
}

type TestNewClientData struct {
	Input  *TestNewClientDataIn
	Output *TestNewClientDataOut
}

func TestNewClient(t *testing.T) {
	for _, tt := range []TestNewClientData{
		// All params valid, live server, no err
		{
			Input: &TestNewClientDataIn{
				"abc",
				"def",
				false,
				"",
			},
			Output: &TestNewClientDataOut{
				"abc",
				"def",
				"https://api.godaddy.com",
				false,
			},
		},
		// All params valid, OTE server, no err
		{
			Input: &TestNewClientDataIn{
				"abc",
				"def",
				true,
				"",
			},
			Output: &TestNewClientDataOut{
				"abc",
				"def",
				"https://api.ote-godaddy.com",
				false,
			},
		},
	} {
		c, err := NewClient(tt.Input.Key, tt.Input.Secret, tt.Input.OTE)
		if err != nil {
			if !tt.Output.Error {
				t.Errorf("daddy.NewClient: expected success, had error: %s", err.Error())
			}
			continue
		}

		if c.Key != tt.Output.Key {
			t.Errorf("daddy.NewClient: expected key %s, got %s", tt.Output.Key, c.Key)
		}

		if c.Secret != tt.Output.Secret {
			t.Errorf("daddy.NewClient: expected secret %s, got %s", tt.Output.Key, c.Key)
		}

		if c.BaseURL != tt.Output.BaseURL {
			t.Errorf("daddy.NewClient: expected base URL %s, got %s", tt.Output.Key, c.Key)
		}
	}
}

func TestNewClientWithURL(t *testing.T) {
	for _, tt := range []TestNewClientData{
		// No key, other params valid, err returned
		{
			Input: &TestNewClientDataIn{
				"",
				"def",
				true,
				"https://example",
			},
			Output: &TestNewClientDataOut{
				"",
				"",
				"",
				true,
			},
		},
		// No secret, other params valid, err returned
		{
			Input: &TestNewClientDataIn{
				"abc",
				"",
				true,
				"https://example",
			},
			Output: &TestNewClientDataOut{
				"",
				"",
				"",
				true,
			},
		},
		// No key or secret, err returned
		{
			Input: &TestNewClientDataIn{
				"",
				"",
				true,
				"https://example",
			},
			Output: &TestNewClientDataOut{
				"",
				"",
				"",
				true,
			},
		},
		// Valid key and secret, no URL, err returned
		{
			Input: &TestNewClientDataIn{
				"abc",
				"def",
				true,
				"",
			},
			Output: &TestNewClientDataOut{
				"abc",
				"def",
				"",
				true,
			},
		},
	} {
		c, err := NewClientWithURL(tt.Input.Key, tt.Input.Secret, tt.Input.BaseURL)
		if err != nil {
			if !tt.Output.Error {
				t.Errorf("daddy.NewClient: expected success, had error: %s", err.Error())
			}
			continue
		}

		if c.Key != tt.Output.Key {
			t.Errorf("daddy.NewClient: expected key %s, got %s", tt.Output.Key, c.Key)
		}

		if c.Secret != tt.Output.Secret {
			t.Errorf("daddy.NewClient: expected secret %s, got %s", tt.Output.Key, c.Key)
		}

		if c.BaseURL != tt.Output.BaseURL {
			t.Errorf("daddy.NewClient: expected base URL %s, got %s", tt.Output.Key, c.Key)
		}
	}
}

func TestError(t *testing.T) {
	err := Error{Message: "abcdef"}
	if err.Error() != "abcdef" {
		t.Errorf("daddy.Error: expected 'abcdef', got: %s", err.Error())
	}
}
