// Copyright 2019 A. Wolcott. All rights reserved.
//
// Use of this source code is governed by the ISC
// license that can be found in the LICENSE file.

package daddy

// AbuseTicket represents an abuse ticket
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

// AbuseTicketCreate is a representation of a new abuse ticket to be submitted
type AbuseTicketCreate struct {
	Info        string
	InfoURL     string
	Intentional bool
	Proxy       string
	Source      string
	Target      string
	Type        string
}

// AbuseTicketID represents an ID corresponding to an abuse ticket
type AbuseTicketID struct {
	TicketID string
}

// AbuseTicketList is a paginated listing of ticket IDs
type AbuseTicketList struct {
	Pagination Pagination
	TicketIDs  []string
}
