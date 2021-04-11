// Copyright 2019 A. Wolcott. All rights reserved.
//
// Use of this source code is governed by the ISC
// license that can be found in the LICENSE file.

package daddy

import (
	"encoding/json"
)

// AbuseService handles communication with the abuse related
// methods of the GoDaddy API.
//
// GoDaddy API docs: https://developer.godaddy.com/doc/endpoint/abuse
type AbuseService service

// ListTickets lists all abuse tickets ids that match user provided filters
func (s *AbuseService) ListTickets(ticketType string, closed bool, sourceDomainOrIP string,
	target string, createdStart string, createdEnd string, limit int,
	offset int) (*AbuseTicketList, error) {
	res := new(AbuseTicketList)

	uri, err := BuildQuery("/v1/abuse/tickets", map[string]interface{}{
		"type":             ticketType,
		"closed":           closed,
		"sourceDomainOrIp": sourceDomainOrIP,
		"target":           target,
		"createdStart":     createdStart,
		"createdEnd":       createdEnd,
		"limit":            limit,
		"offset":           offset,
	})

	if err != nil {
		return nil, err
	}

	data, err := s.client.Get(uri)
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(data, &res)

	return res, err
}

// NewTicket creates a new abuse ticket
func (s *AbuseService) NewTicket(ticket *AbuseTicketCreate) (*AbuseTicketID, error) {
	var id = new(AbuseTicketID)

	enc, err := json.Marshal(ticket)
	if err != nil {
		return nil, err
	}
	data, err := s.client.Post("/v1/abuse/tickets", enc)
	if err != nil {
		return id, err
	}

	err = json.Unmarshal(data, &id)

	return id, err
}

// GetTicket returns the abuse ticket data for a given ticket id
func (s *AbuseService) GetTicket(ticketID string) (*AbuseTicket, error) {
	var id = new(AbuseTicket)

	data, err := s.client.Get("/v1/abuse/tickets/" + ticketID)
	if err != nil {
		return id, err
	}

	err = json.Unmarshal(data, &id)

	return id, err
}
