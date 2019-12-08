package abuse

import (
	"encoding/json"

	"github.com/alyx/godaddy"
)

// Tickets lists all abuse tickets ids that match user provided filters
func Tickets() {
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
