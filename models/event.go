package models

import (
	"context"
	"github.com/UoYMathSoc/2020-site/database"
	"time"
)

type Event struct {
	ID          int
	Name        string
	StartDate   time.Time
	EndDate     time.Time
	Location    string
	Description string
}

func (s *Session) ListEvents() ([]Event, error) {
	events, err := s.querier.ListEvents(context.Background())
	if err != nil {
		return nil, err
	}
	var result []Event
	for _, event := range events {
		sanitiseEvent(&event)

		startDate := event.Date.Add(event.StartTime.Sub(event.StartTime.Truncate(time.Hour * 24)))
		endDate := event.Date.Add(event.EndTime.Time.Sub(event.EndTime.Time.Truncate(time.Hour * 24)))
		event := Event{
			ID:          int(event.ID),
			Name:        event.Name,
			StartDate:   startDate,
			EndDate:     endDate,
			Location:    event.Location.String,
			Description: event.Description.String,
		}
		result = append(result, event)
	}
	return result, nil
}

func (s *Session) GetEvent(id int) (*Event, error) {
	event, err := s.querier.GetEvent(context.Background(), int32(id))
	if err != nil {
		return nil, err
	}

	sanitiseEvent(&event)

	startDate := event.Date.Add(event.StartTime.Sub(event.StartTime.Truncate(time.Hour * 24)))
	endDate := event.Date.Add(event.EndTime.Time.Sub(event.EndTime.Time.Truncate(time.Hour * 24)))
	return &Event{
		ID:          int(event.ID),
		Name:        event.Name,
		StartDate:   startDate,
		EndDate:     endDate,
		Location:    event.Location.String,
		Description: event.Description.String,
	}, nil
}

// Maybe use sql defaults?
func sanitiseEvent(event *database.Event) {
	if !event.EndTime.Valid {
		event.EndTime.Time = event.StartTime.Add(time.Hour)
		event.EndTime.Valid = true
	}
	if !event.Description.Valid {
		event.Description.String = ""
		event.Description.Valid = true
	}
	if !event.Location.Valid {
		event.Location.String = ""
		event.Location.Valid = true
	}
}
