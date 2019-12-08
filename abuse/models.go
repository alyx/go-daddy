package abuse

import "github.com/alyx/godaddy"

// Ticket defines the content of a ticket generated from GetTicket()
type Ticket struct {
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

type TicketCreate struct {
	Info        string
	InfoUrl     string
	Intentional bool
	Proxy       string
	Source      string
	Target      string
	Type        string
}

type TicketId struct {
	TicketId string
}

type TicketList struct {
	Pagination godaddy.Pagination
	TicketIds  []string
}
