package domain

import (
	"errors"

	"github.com/google/uuid"
)

type Spot struct {
	ID       string
	EventID  string
	TicketID string
	Name     string
	Status   SpotStatus
}

type SpotStatus string

const (
	SpotStatusAvailable SpotStatus = "available"
	SpotStatusSold      SpotStatus = "sold"
)

var (
	ErrSpotNameRequired    = errors.New("spot name is required")
	ErrSpotNameTooShort    = errors.New("spot name must be at least two characters long")
	ErrSpotNameStartLetter = errors.New("spot name must start with a letter")
	ErrSpotNameEndNumber   = errors.New("spot name must end with a number")
	ErrSpotNumberInvalid   = errors.New("spot number invalid")
	ErrSpotNotFound        = errors.New("spot not found")
	ErrSpotAlreadyReserved = errors.New("spot already reserved")
)

func (spot *Spot) Validate() error {
	if len(spot.Name) == 0 {
		return ErrSpotNameRequired
	}
	if len(spot.Name) < 2 {
		return ErrSpotNameTooShort
	}
	if spot.Name[0] < 'A' || spot.Name[0] > 'Z' {
		return ErrSpotNameStartLetter
	}
	if spot.Name[1] < '0' || spot.Name[1] > '9' {
		return ErrSpotNameEndNumber
	}
	return nil
}

func NewSpot(event *Event, name string) (*Spot, error) {
	spot := &Spot{
		ID:      uuid.New().String(),
		EventID: event.ID,
		Name:    name,
		Status:  SpotStatusAvailable,
	}
	hasError := spot.Validate()
	if hasError != nil {
		return nil, hasError
	}
	return spot, nil
}

func (spot *Spot) Reserve(ticketID string) error {
	if spot.Status == SpotStatusSold {
		return ErrSpotAlreadyReserved
	}
	spot.Status = SpotStatusSold
	spot.TicketID = ticketID
	return nil
}
