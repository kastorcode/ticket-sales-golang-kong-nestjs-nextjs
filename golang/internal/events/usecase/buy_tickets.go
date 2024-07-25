package usecase

import (
	"ticket_sales_api/internal/events/domain"
	"ticket_sales_api/internal/events/infra/service"
)

type BuyTicketsInputDTO struct {
	EventID    string   `json:"eventId"`
	Spots      []string `json:"spots"`
	TicketKind string   `json:"ticketKind"`
	CardHash   string   `json:"cardHash"`
	Email      string   `json:"email"`
}

type BuyTicketsOutputDTO struct {
	Tickets []TicketDTO `json:"tickets"`
}

type BuyTicketsUseCase struct {
	repository     domain.EventRepository
	partnerFactory service.PartnerFactory
}

func NewBuyTicketsUseCase(repository domain.EventRepository, partnerFactory service.PartnerFactory) *BuyTicketsUseCase {
	return &BuyTicketsUseCase{
		repository:     repository,
		partnerFactory: partnerFactory,
	}
}

func (useCase *BuyTicketsUseCase) Execute(input BuyTicketsInputDTO) (*BuyTicketsOutputDTO, error) {
	event, err := useCase.repository.FindEventByID(input.EventID)
	if err != nil {
		return nil, err
	}
	request := &service.ReservationRequest{
		EventID:    input.EventID,
		Spots:      input.Spots,
		TicketKind: input.TicketKind,
		CardHash:   input.CardHash,
		Email:      input.Email,
	}
	partnerService, err := useCase.partnerFactory.CreatePartner(event.PartnerID)
	if err != nil {
		return nil, err
	}
	response, err := partnerService.MakeReservation(request)
	if err != nil {
		return nil, err
	}
	tickets := make([]domain.Ticket, len(response))
	for i, reservation := range response {
		spot, err := useCase.repository.FindSpotByName(event.ID, reservation.Spot)
		if err != nil {
			return nil, err
		}
		ticket, err := domain.NewTicket(event, spot, domain.TicketKind(reservation.TicketKind))
		if err != nil {
			return nil, err
		}
		err = useCase.repository.CreateTicket(ticket)
		if err != nil {
			return nil, err
		}
		err = spot.Reserve(ticket.ID)
		if err != nil {
			return nil, err
		}
		err = useCase.repository.ReserveSpot(spot.ID, ticket.ID)
		if err != nil {
			return nil, err
		}
		tickets[i] = *ticket
	}
	ticketsDTO := make([]TicketDTO, len(tickets))
	for i, ticket := range tickets {
		ticketsDTO[i] = TicketDTO{
			ID:         ticket.ID,
			SpotID:     ticket.Spot.ID,
			TicketKind: string(ticket.TicketKind),
			Price:      ticket.Price,
		}
	}
	return &BuyTicketsOutputDTO{Tickets: ticketsDTO}, nil
}
