package repository

import (
	"database/sql"
	"errors"
	"time"

	"ticket_sales_api/internal/events/domain"

	_ "github.com/lib/pq"
)

type EventRepositoryPostgres struct {
	db *sql.DB
}

func NewEventRepositoryPostgres(db *sql.DB) (domain.EventRepository, error) {
	return &EventRepositoryPostgres{db: db}, nil
}

func (repository *EventRepositoryPostgres) ListEvents() ([]domain.Event, error) {
	query := `
		SELECT 
			e.id, e.name, e.location, e.organization, e.rating, e.date, e.imageUrl, e.capacity, e.price, e.partnerId,
			s.id, s.eventId, s.name, s.status, s.ticketId,
			t.id, t.eventId, t.spotId, t.ticketKind, t.price
		FROM events e
		LEFT JOIN spots s ON e.id = s.eventId
		LEFT JOIN tickets t ON s.id = t.spotId
	`
	rows, err := repository.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	eventMap := make(map[string]*domain.Event)
	spotMap := make(map[string]*domain.Spot)
	for rows.Next() {
		var eventID, eventName, eventLocation, eventOrganization, eventRating, eventImageURL, spotID, spotEventID, spotName, spotStatus, spotTicketID, ticketID, ticketEventID, ticketSpotID, ticketKind sql.NullString
		var eventDate sql.NullString
		var eventCapacity int
		var eventPrice, ticketPrice sql.NullFloat64
		var partnerID sql.NullInt32
		err := rows.Scan(
			&eventID, &eventName, &eventLocation, &eventOrganization, &eventRating, &eventDate, &eventImageURL, &eventCapacity, &eventPrice, &partnerID,
			&spotID, &spotEventID, &spotName, &spotStatus, &spotTicketID,
			&ticketID, &ticketEventID, &ticketSpotID, &ticketKind, &ticketPrice,
		)
		if err != nil {
			return nil, err
		}
		if !eventID.Valid || !eventName.Valid || !eventLocation.Valid || !eventOrganization.Valid || !eventRating.Valid || !eventDate.Valid || !eventImageURL.Valid || !eventPrice.Valid || !partnerID.Valid {
			continue
		}
		event, exists := eventMap[eventID.String]
		if !exists {
			eventDateParsed, err := time.Parse("2006-01-02T15:04:05Z", eventDate.String)
			if err != nil {
				return nil, err
			}
			event = &domain.Event{
				ID:           eventID.String,
				Name:         eventName.String,
				Location:     eventLocation.String,
				Organization: eventOrganization.String,
				Rating:       domain.Rating(eventRating.String),
				Date:         eventDateParsed,
				ImageURL:     eventImageURL.String,
				Capacity:     eventCapacity,
				Price:        eventPrice.Float64,
				PartnerID:    int(partnerID.Int32),
				Spots:        []domain.Spot{},
				Tickets:      []domain.Ticket{},
			}
			eventMap[eventID.String] = event
		}
		if spotID.Valid {
			spot, spotExists := spotMap[spotID.String]
			if !spotExists {
				spot = &domain.Spot{
					ID:       spotID.String,
					EventID:  spotEventID.String,
					Name:     spotName.String,
					Status:   domain.SpotStatus(spotStatus.String),
					TicketID: spotTicketID.String,
				}
				event.Spots = append(event.Spots, *spot)
				spotMap[spotID.String] = spot
			}

			if ticketID.Valid {
				ticket := domain.Ticket{
					ID:         ticketID.String,
					EventID:    ticketEventID.String,
					Spot:       spot,
					TicketKind: domain.TicketKind(ticketKind.String),
					Price:      ticketPrice.Float64,
				}
				event.Tickets = append(event.Tickets, ticket)
			}
		}
	}
	events := make([]domain.Event, 0, len(eventMap))
	for _, event := range eventMap {
		events = append(events, *event)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return events, nil
}

func (repository *EventRepositoryPostgres) FindEventByID(eventID string) (*domain.Event, error) {
	query := `
		SELECT 
			e.id, e.name, e.location, e.organization, e.rating, e.date, e.imageUrl, e.capacity, e.price, e.partnerId,
			s.id, s.eventId, s.name, s.status, s.ticketId,
			t.id, t.eventId, t.spotId, t.ticketKind, t.price
		FROM events e
		LEFT JOIN spots s ON e.id = s.eventId
		LEFT JOIN tickets t ON s.id = t.spotId
		WHERE e.id = $1
	`
	rows, err := repository.db.Query(query, eventID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var event *domain.Event
	for rows.Next() {
		var eventIDStr, eventName, eventLocation, eventOrganization, eventRating, eventImageURL, spotID, spotEventID, spotName, spotStatus, spotTicketID, ticketID, ticketEventID, ticketSpotID, ticketKind sql.NullString
		var eventDate sql.NullString
		var eventCapacity int
		var eventPrice, ticketPrice sql.NullFloat64
		var partnerID sql.NullInt32
		err := rows.Scan(
			&eventIDStr, &eventName, &eventLocation, &eventOrganization, &eventRating, &eventDate, &eventImageURL, &eventCapacity, &eventPrice, &partnerID,
			&spotID, &spotEventID, &spotName, &spotStatus, &spotTicketID,
			&ticketID, &ticketEventID, &ticketSpotID, &ticketKind, &ticketPrice,
		)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, domain.ErrEventNotFound
			}
			return nil, err
		}
		if !eventIDStr.Valid || !eventName.Valid || !eventLocation.Valid || !eventOrganization.Valid || !eventRating.Valid || !eventDate.Valid || !eventImageURL.Valid || !eventPrice.Valid || !partnerID.Valid {
			continue
		}
		if event == nil {
			eventDateParsed, err := time.Parse("2006-01-02T15:04:05Z", eventDate.String)
			if err != nil {
				return nil, err
			}
			event = &domain.Event{
				ID:           eventIDStr.String,
				Name:         eventName.String,
				Location:     eventLocation.String,
				Organization: eventOrganization.String,
				Rating:       domain.Rating(eventRating.String),
				Date:         eventDateParsed,
				ImageURL:     eventImageURL.String,
				Capacity:     eventCapacity,
				Price:        eventPrice.Float64,
				PartnerID:    int(partnerID.Int32),
				Spots:        []domain.Spot{},
				Tickets:      []domain.Ticket{},
			}
		}
		if spotID.Valid {
			spot := domain.Spot{
				ID:       spotID.String,
				EventID:  spotEventID.String,
				Name:     spotName.String,
				Status:   domain.SpotStatus(spotStatus.String),
				TicketID: spotTicketID.String,
			}
			event.Spots = append(event.Spots, spot)
			if ticketID.Valid {
				ticket := domain.Ticket{
					ID:         ticketID.String,
					EventID:    ticketEventID.String,
					Spot:       &spot,
					TicketKind: domain.TicketKind(ticketKind.String),
					Price:      ticketPrice.Float64,
				}
				event.Tickets = append(event.Tickets, ticket)
			}
		}
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	if event == nil {
		return nil, domain.ErrEventNotFound
	}
	return event, nil
}

func (repository *EventRepositoryPostgres) CreateEvent(event *domain.Event) error {
	query := `
		INSERT INTO events (id, name, location, organization, rating, date, imageUrl, capacity, price, partnerId)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	`
	_, err := repository.db.Exec(query, event.ID, event.Name, event.Location, event.Organization, event.Rating, event.Date.Format("2006-01-02 15:04:05"), event.ImageURL, event.Capacity, event.Price, event.PartnerID)
	return err
}

func (repository *EventRepositoryPostgres) FindSpotByID(spotID string) (*domain.Spot, error) {
	query := `
		SELECT
			s.id, s.eventId, s.name, s.status, s.ticketId,
			t.id, t.eventId, t.spotId, t.ticketKind, t.price
		FROM spots s
		LEFT JOIN tickets t ON s.id = t.spotId
		WHERE s.id = $1
	`
	row := repository.db.QueryRow(query, spotID)
	var spot domain.Spot
	var ticket domain.Ticket
	var ticketID, ticketEventID, ticketSpotID, ticketKind sql.NullString
	var ticketPrice sql.NullFloat64
	err := row.Scan(
		&spot.ID, &spot.EventID, &spot.Name, &spot.Status, &spot.TicketID,
		&ticketID, &ticketEventID, &ticketSpotID, &ticketKind, &ticketPrice,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrSpotNotFound
		}
		return nil, err
	}
	if ticketID.Valid {
		ticket.ID = ticketID.String
		ticket.EventID = ticketEventID.String
		ticket.Spot = &spot
		ticket.TicketKind = domain.TicketKind(ticketKind.String)
		ticket.Price = ticketPrice.Float64
		spot.TicketID = ticket.ID
	}
	return &spot, nil
}

func (repository *EventRepositoryPostgres) CreateSpot(spot *domain.Spot) error {
	query := `
		INSERT INTO spots (id, eventId, name, status, ticketId)
		VALUES ($1, $2, $3, $4, $5)
	`
	_, err := repository.db.Exec(query, spot.ID, spot.EventID, spot.Name, spot.Status, spot.TicketID)
	return err
}

func (repository *EventRepositoryPostgres) CreateTicket(ticket *domain.Ticket) error {
	query := `
		INSERT INTO tickets (id, eventId, spotId, ticketKind, price)
		VALUES ($1, $2, $3, $4, $5)
	`
	_, err := repository.db.Exec(query, ticket.ID, ticket.EventID, ticket.Spot.ID, ticket.TicketKind, ticket.Price)
	return err
}

func (repository *EventRepositoryPostgres) ReserveSpot(spotID, ticketID string) error {
	query := `
		UPDATE spots
		SET status = $1, ticketId = $2
		WHERE id = $3
	`
	_, err := repository.db.Exec(query, domain.SpotStatusSold, ticketID, spotID)
	return err
}

func (repository *EventRepositoryPostgres) FindSpotsByEventID(eventID string) ([]*domain.Spot, error) {
	query := `
		SELECT id, eventId, name, status, ticketId
		FROM spots
		WHERE eventId = $1
	`
	rows, err := repository.db.Query(query, eventID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var spots []*domain.Spot
	for rows.Next() {
		var spot domain.Spot
		if err := rows.Scan(&spot.ID, &spot.EventID, &spot.Name, &spot.Status, &spot.TicketID); err != nil {
			return nil, err
		}
		spots = append(spots, &spot)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return spots, nil
}

func (repository *EventRepositoryPostgres) FindSpotByName(eventID, name string) (*domain.Spot, error) {
	query := `
		SELECT 
			s.id, s.eventId, s.name, s.status, s.ticketId,
			t.id, t.eventId, t.spotId, t.ticketKind, t.price
		FROM spots s
		LEFT JOIN tickets t ON s.id = t.spotId
		WHERE s.eventId = $1 AND s.name = $2
	`
	row := repository.db.QueryRow(query, eventID, name)
	var spot domain.Spot
	var ticket domain.Ticket
	var ticketID, ticketEventID, ticketSpotID, ticketKind sql.NullString
	var ticketPrice sql.NullFloat64
	err := row.Scan(
		&spot.ID, &spot.EventID, &spot.Name, &spot.Status, &spot.TicketID,
		&ticketID, &ticketEventID, &ticketSpotID, &ticketKind, &ticketPrice,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrSpotNotFound
		}
		return nil, err
	}
	if ticketID.Valid {
		ticket.ID = ticketID.String
		ticket.EventID = ticketEventID.String
		ticket.Spot = &spot
		ticket.TicketKind = domain.TicketKind(ticketKind.String)
		ticket.Price = ticketPrice.Float64
		spot.TicketID = ticket.ID
	}
	return &spot, nil
}
