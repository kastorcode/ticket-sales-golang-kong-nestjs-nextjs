package http

import (
	"encoding/json"
	"net/http"
	"ticket_sales_api/internal/events/usecase"
)

type EventsHandler struct {
	buyTicketsUseCase *usecase.BuyTicketsUseCase
	getEventUseCase   *usecase.GetEventUseCase
	listEventsUseCase *usecase.ListEventsUseCase
	listSpotsUseCase  *usecase.ListSpotsUseCase
}

func NewEventsHandler(
	buyTicketsUseCase *usecase.BuyTicketsUseCase,
	getEventUseCase *usecase.GetEventUseCase,
	listEventsUseCase *usecase.ListEventsUseCase,
	listSpotsUseCase *usecase.ListSpotsUseCase,
) *EventsHandler {
	return &EventsHandler{
		buyTicketsUseCase: buyTicketsUseCase,
		getEventUseCase:   getEventUseCase,
		listEventsUseCase: listEventsUseCase,
		listSpotsUseCase:  listSpotsUseCase,
	}
}

func (handler *EventsHandler) BuyTickets(writer http.ResponseWriter, request *http.Request) {
	var input usecase.BuyTicketsInputDTO
	if err := json.NewDecoder(request.Body).Decode(&input); err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	tickets, err := handler.buyTicketsUseCase.Execute(input)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(tickets)
}

func (handler *EventsHandler) GetEvent(writer http.ResponseWriter, request *http.Request) {
	eventID := request.PathValue("eventID")
	input := usecase.GetEventInputDTO{ID: eventID}
	event, err := handler.getEventUseCase.Execute(input)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(event)
}

func (handler *EventsHandler) ListEvents(writer http.ResponseWriter, request *http.Request) {
	events, err := handler.listEventsUseCase.Execute()
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(events)
}

func (handler *EventsHandler) ListSpots(writer http.ResponseWriter, request *http.Request) {
	eventID := request.PathValue("eventID")
	input := usecase.ListSpotsInputDTO{EventID: eventID}
	spots, err := handler.listSpotsUseCase.Execute(input)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(spots)
}
