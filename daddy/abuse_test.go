// Copyright 2019 A. Wolcott. All rights reserved.
//
// Use of this source code is governed by the ISC
// license that can be found in the LICENSE file.

package daddy

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

type TestEndpointData struct {
	ID           string
	URI          string
	ResponseCode int
	DataFile     string
	Response     []byte
	Failure      bool
	JSONFailure  bool
}

func TestGetTicket(t *testing.T) {
	for _, tt := range []TestEndpointData{
		{
			ID:           "1",
			URI:          "/v1/abuse/tickets/1",
			ResponseCode: 200,
			DataFile:     "test_data/abuse_getticket_1.json",
			Failure:      false,
		},
		{
			ID:           "2",
			URI:          "/v1/abuse/tickets/2",
			ResponseCode: 404,
			DataFile:     "test_data/abuse_getticket_2.json",
			Failure:      true,
		},
		{
			ID:           "a",
			URI:          "/v1/abuse/tickets/a",
			ResponseCode: 200,
			DataFile:     "test_data/abuse_getticket_3.json",
			Failure:      true,
			JSONFailure:  true,
		},
	} {
		dat, err := ioutil.ReadFile(tt.DataFile)
		var enc map[string]interface{}
		json.Unmarshal(dat, &enc)

		if err != nil {
			if tt.JSONFailure != true {
				t.Errorf("BROKEN TEST: %s", err.Error())
				continue
			}
		}
		server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			if req.URL.String() != tt.URI {
				t.Errorf("abuse.GetTicket: want URI %s, got %s", tt.URI, req.URL.String())
			}
			if tt.ResponseCode <= 299 {
				rw.Write(dat)
			} else {
				http.Error(rw, http.StatusText(tt.ResponseCode), tt.ResponseCode)
			}
		}))

		defer server.Close()

		c, err := NewClientWithURL("abc", "def", server.URL)
		if err != nil {
			t.Errorf("BROKEN TEST: %s", err.Error())
			continue
		}

		out, err := c.Abuse.GetTicket(tt.ID)
		if err != nil {
			if tt.Failure != true {
				t.Errorf("abuse.GetTicket: error %s", err.Error())
			}
			continue
		}

		if out.Closed != enc["closed"] {
			t.Errorf("abuse.GetTicket().Closed: expected %t, got %t", enc["closed"], out.Closed)
		}
		if out.ClosedAt != enc["closedAt"] {
			t.Errorf("abuse.GetTicket().ClosedAt: expected %s, got %s", enc["closedAt"], out.ClosedAt)
		}
		if out.CreatedAt != enc["createdAt"] {
			t.Errorf("abuse.GetTicket().CreatedAt: expected %s, got %s", enc["createdAt"], out.CreatedAt)
		}
		if out.DomainIP != enc["domainIp"] {
			t.Errorf("abuse.GetTicket().DomainIP: expected %s, got %s", enc["domainIp"], out.DomainIP)
		}
		if out.Reporter != enc["reporter"] {
			t.Errorf("abuse.GetTicket().Reporter: expected %s, got %s", enc["reporter"], out.Reporter)
		}
		if out.Source != enc["source"] {
			t.Errorf("abuse.GetTicket().Source: expected %s, got %s", enc["source"], out.Source)
		}
		if out.Target != enc["target"] {
			t.Errorf("abuse.GetTicket().Target: expected %s, got %s", enc["target"], out.Target)
		}
		if out.TicketID != enc["ticketId"] {
			t.Errorf("abuse.GetTicket().TicketID: expected %s, got %s", enc["ticketId"], out.TicketID)
		}
		if out.Type != enc["type"] {
			t.Errorf("abuse.GetTicket().Type: expected %s, got %s", enc["type"], out.Type)
		}
	}
}
