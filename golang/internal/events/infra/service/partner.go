package service

type ReservationRequest struct {
	EventID    string   `json:"eventId"`
	Spots      []string `json:"spots"`
	TicketKind string   `json:"ticketKind"`
	CardHash   string   `json:"cardHash"`
	Email      string   `json:"email"`
}

type ReservationResponse struct {
	ID         string `json:"id"`
	EventID    string `json:"eventId"`
	Email      string `json:"email"`
	Spot       string `json:"spot"`
	TicketKind string `json:"ticketKind"`
	Status     string `json:"status"`
}

type Partner interface {
	MakeReservation(request *ReservationRequest) ([]ReservationResponse, error)
}
