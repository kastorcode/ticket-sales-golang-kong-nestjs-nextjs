package usecase

import (
	"ticket_sales_api/internal/events/domain"
)

type GetEventInputDTO struct {
	ID string
}

type GetEventOutputDTO struct {
	ID           string  `json:"id"`
	Name         string  `json:"name"`
	Location     string  `json:"location"`
	Organization string  `json:"organization"`
	Rating       string  `json:"rating"`
	Date         string  `json:"date"`
	Capacity     int     `json:"capacity"`
	Price        float64 `json:"price"`
	PartnerID    int     `json:"partnerId"`
}

type GetEventUseCase struct {
	repository domain.EventRepository
}

func NewGetEventUseCase(repository domain.EventRepository) *GetEventUseCase {
	return &GetEventUseCase{repository: repository}
}

func (useCase *GetEventUseCase) Execute(input GetEventInputDTO) (*GetEventOutputDTO, error) {
	event, err := useCase.repository.FindEventByID(input.ID)
	if err != nil {
		return nil, err
	}
	return &GetEventOutputDTO{
		ID:           event.ID,
		Name:         event.Name,
		Location:     event.Location,
		Organization: event.Organization,
		Rating:       string(event.Rating),
		Date:         event.Date.Format("2006-01-02 15:04:05"),
		Capacity:     event.Capacity,
		Price:        event.Price,
		PartnerID:    event.PartnerID,
	}, nil
}
