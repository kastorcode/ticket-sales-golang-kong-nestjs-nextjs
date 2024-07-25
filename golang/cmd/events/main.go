package main

import (
	"database/sql"
	"net/http"
	httpHandler "ticket_sales_api/internal/events/infra/http"
	"ticket_sales_api/internal/events/infra/repository"
	"ticket_sales_api/internal/events/infra/service"
	"ticket_sales_api/internal/events/usecase"
)

func main() {

	db, err := sql.Open("postgres", "postgresql://root:root@host.docker.internal:5432/golang?sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	eventRepository, err := repository.NewEventRepositoryPostgres(db)
	if err != nil {
		panic(err)
	}

	partnerFactory := service.NewPartnerFactory(map[int]string{
		1: "http://host.docker.internal:8000/partner1",
		2: "http://host.docker.internal:8000/partner2",
	})

	eventsHandler := httpHandler.NewEventsHandler(
		usecase.NewBuyTicketsUseCase(eventRepository, partnerFactory),
		usecase.NewGetEventUseCase(eventRepository),
		usecase.NewListEventsUseCase(eventRepository),
		usecase.NewListSpotsUseCase(eventRepository),
	)

	serveMux := http.NewServeMux()
	serveMux.HandleFunc("GET /events", eventsHandler.ListEvents)
	serveMux.HandleFunc("GET /events/{eventID}", eventsHandler.GetEvent)
	serveMux.HandleFunc("GET /events/{eventID}/spots", eventsHandler.ListSpots)
	serveMux.HandleFunc("POST /checkout", eventsHandler.BuyTickets)

	println("Listening on port 8080")
	if err := http.ListenAndServe(":8080", serveMux); err != nil {
		panic(err)
	}
}
