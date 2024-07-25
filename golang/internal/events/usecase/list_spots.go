package usecase

import "ticket_sales_api/internal/events/domain"

type ListSpotsInputDTO struct {
	EventID string `json:"eventId"`
}

type ListSpotsOutputDTO struct {
	Event EventDTO  `json:"event"`
	Spots []SpotDTO `json:"spots"`
}

type ListSpotsUseCase struct {
	repository domain.EventRepository
}

func NewListSpotsUseCase(repository domain.EventRepository) *ListSpotsUseCase {
	return &ListSpotsUseCase{repository: repository}
}

func (useCase *ListSpotsUseCase) Execute(input ListSpotsInputDTO) (*ListSpotsOutputDTO, error) {
	event, err := useCase.repository.FindEventByID(input.EventID)
	if err != nil {
		return nil, err
	}
	spots, err := useCase.repository.FindSpotsByEventID(input.EventID)
	if err != nil {
		return nil, err
	}
	spotsDTO := make([]SpotDTO, len(spots))
	for i, spot := range spots {
		spotsDTO[i] = SpotDTO{
			ID:       spot.ID,
			Name:     spot.Name,
			Status:   string(spot.Status),
			TicketID: spot.TicketID,
		}
	}
	eventDTO := EventDTO{
		ID:           event.ID,
		Name:         event.Name,
		Location:     event.Location,
		Organization: event.Organization,
		Rating:       string(event.Rating),
		Date:         event.Date.Format("2006-01-02 15:04:05"),
		Capacity:     event.Capacity,
		Price:        event.Price,
		PartnerID:    event.PartnerID,
		ImageURL:     event.ImageURL,
	}
	return &ListSpotsOutputDTO{Event: eventDTO, Spots: spotsDTO}, nil
}
