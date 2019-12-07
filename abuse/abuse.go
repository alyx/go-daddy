package abuse

import (
	"time"

	"github.com/alyx/godaddy"
)

// Ticket defines the content of a ticket generated from GetTicket()
type Ticket struct {
	Closed    bool
	ClosedAt  time.Time
	CreatedAt time.Time
	DomainIP  string
	Reporter  string
	Source    string
	Target    string
	TicketID  string
	Type      string
}

func Tickets(godaddy.Client) { return }

func NewTicket() { return }

func GetTicket() { return }
