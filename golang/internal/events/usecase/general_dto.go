package usecase

type EventDTO struct {
	ID           string  `json:"id"`
	Name         string  `json:"name"`
	Location     string  `json:"location"`
	Organization string  `json:"organization"`
	Rating       string  `json:"rating"`
	Date         string  `json:"date"`
	ImageURL     string  `json:"imageUrl"`
	Capacity     int     `json:"capacity"`
	Price        float64 `json:"price"`
	PartnerID    int     `json:"partnerId"`
}

type SpotDTO struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	EventID  string `json:"eventId"`
	Reserved bool   `json:"reserved"`
	Status   string `json:"status"`
	TicketID string `json:"ticketId"`
}

type TicketDTO struct {
	ID         string  `json:"id"`
	SpotID     string  `json:"spotId"`
	TicketKind string  `json:"ticketKind"`
	Price      float64 `json:"price"`
}
