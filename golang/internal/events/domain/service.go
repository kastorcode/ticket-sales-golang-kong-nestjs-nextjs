package domain

import (
	"errors"
	"fmt"
)

type SpotService struct{}

var (
	ErrQuantityInvalid = errors.New("quantity must be greater than zero")
)

func NewSpotService() *SpotService {
	return &SpotService{}
}

func (spotService *SpotService) GenerateSpots(event *Event, quantity int) error {
	if quantity <= 0 {
		return ErrQuantityInvalid
	}
	for i := range quantity {
		spotName := fmt.Sprintf("%c%d", 'A'+i/10, i%10+1)
		spot, err := NewSpot(event, spotName)
		if err != nil {
			return err
		}
		event.Spots = append(event.Spots, *spot)
	}
	return nil
}
