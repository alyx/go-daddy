// Copyright 2019 A. Wolcott. All rights reserved.
//
// Use of this source code is governed by the ISC
// license that can be found in the LICENSE file.

package daddy

// Ticket defines the content of a ticket generated from GetTicket()
type AbuseTicket struct {
	Closed    bool
	ClosedAt  string
	CreatedAt string
	DomainIP  string
	Reporter  string
	Source    string
	Target    string
	TicketID  string
	Type      string
}

type AbuseTicketCreate struct {
	Info        string
	InfoURL     string
	Intentional bool
	Proxy       string
	Source      string
	Target      string
	Type        string
}

type AbuseTicketID struct {
	TicketID string
}

type AbuseTicketList struct {
	Pagination Pagination
	TicketIDs  []string
}
