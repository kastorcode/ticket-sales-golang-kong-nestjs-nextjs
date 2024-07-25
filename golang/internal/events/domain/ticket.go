package domain

import (
	"errors"

	"github.com/google/uuid"
)

type Ticket struct {
	ID         string
	EventID    string
	Spot       *Spot
	TicketKind TicketKind
	Price      float64
}

type TicketKind string

const (
	TicketKindHalf TicketKind = "half"
	TicketKindFull TicketKind = "full"
)

var (
	ErrTicketPriceZero   = errors.New("price must be greater than zero")
	ErrTicketKindInvalid = errors.New("ticket kind invalid")
)

func NewTicket(event *Event, spot *Spot, ticketKind TicketKind) (*Ticket, error) {
	if !IsValidTicketKind(ticketKind) {
		return nil, ErrTicketKindInvalid
	}
	ticket := &Ticket{
		ID:         uuid.New().String(),
		EventID:    event.ID,
		Spot:       spot,
		TicketKind: ticketKind,
		Price:      event.Price,
	}
	ticket.CalculatePrice()
	if err := ticket.Validate(); err != nil {
		return nil, err
	}
	return ticket, nil
}

func IsValidTicketKind(ticketKind TicketKind) bool {
	return ticketKind == TicketKindHalf || ticketKind == TicketKindFull
}

func (ticket *Ticket) CalculatePrice() {
	if ticket.TicketKind == TicketKindHalf {
		ticket.Price /= 2
	}
}

func (ticket *Ticket) Validate() error {
	if ticket.Price <= 0 {
		return ErrTicketPriceZero
	}
	return nil
}
