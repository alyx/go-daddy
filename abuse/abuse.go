package abuse

import (
	"encoding/json"

	"github.com/alyx/godaddy"
)

// ListTickets lists all abuse tickets ids that match user provided filters
func ListTickets(c *godaddy.Client, ticketType string, closed bool, sourceDomainOrIP string,
	target string, createdStart string, createdEnd string, limit int,
	offset int) (*TicketList, error) {
	res := new(TicketList)

	uri, err := godaddy.BuildQuery("/v1/abuse/tickets", map[string]interface{}{
		"type":             ticketType,
		"closed":           closed,
		"sourceDomainOrIp": sourceDomainOrIP,
		"target":           target,
		"createdStart":     createdStart,
		"createdEnd":       createdEnd,
		"limit":            limit,
		"offset":           offset,
	})

	data, err := c.Get(uri)
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(data, &res)

	return res, err
}

// NewTicket creates a new abuse ticket
func NewTicket(c *godaddy.Client, ticket *TicketCreate) (*TicketID, error) {
	var id = new(TicketID)

	enc, err := json.Marshal(ticket)
	data, err := c.Post("/v1/abuse/tickets", enc)
	if err != nil {
		return id, err
	}

	err = json.Unmarshal(data, &id)

	return id, err
}

// GetTicket returns the abuse ticket data for a given ticket id
func GetTicket(c *godaddy.Client, ticketID string) (*Ticket, error) {
	var id = new(Ticket)

	data, err := c.Get("/v1/abuse/tickets/" + ticketID)
	if err != nil {
		return id, err
	}

	err = json.Unmarshal(data, &id)

	return id, err
}
