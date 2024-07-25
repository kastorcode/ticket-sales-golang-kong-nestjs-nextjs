package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Partner2 struct {
	BaseURL string
}

type Partner2ReservationRequest struct {
	Spots      []string `json:"spots"`
	TicketKind string   `json:"ticketKind"`
	Email      string   `json:"email"`
}

type Partner2ReservationResponse struct {
	ID         string `json:"id"`
	EventID    string `json:"eventID"`
	Email      string `json:"email"`
	Spot       string `json:"spot"`
	TicketKind string `json:"ticketKind"`
	Status     string `json:"status"`
}

func (partner *Partner2) MakeReservation(request *ReservationRequest) ([]ReservationResponse, error) {
	partnerRequest := Partner2ReservationRequest{
		Spots:      request.Spots,
		TicketKind: request.TicketKind,
		Email:      request.Email,
	}
	body, err := json.Marshal(partnerRequest)
	if err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s/events/%s/reserve", partner.BaseURL, request.EventID)
	httpRequest, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	httpRequest.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	httpResponse, err := client.Do(httpRequest)
	if err != nil {
		return nil, err
	}
	defer httpResponse.Body.Close()
	if httpResponse.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("unexpected status code: %d", httpResponse.StatusCode)
	}
	var partnerResponse []Partner2ReservationResponse
	if err := json.NewDecoder(httpResponse.Body).Decode(&partnerResponse); err != nil {
		return nil, err
	}
	responses := make([]ReservationResponse, len(partnerResponse))
	for i, r := range partnerResponse {
		responses[i] = ReservationResponse{
			ID:         r.ID,
			EventID:    r.EventID,
			Email:      r.Email,
			Spot:       r.Spot,
			TicketKind: r.TicketKind,
			Status:     r.Status,
		}
	}
	return responses, nil
}
