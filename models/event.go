package models

import (
	"context"
	"github.com/UoYMathSoc/2020-site/database"
	"time"
)

type Event struct {
	ID          int
	Name        string
	StartTime   time.Time
	EndTime     time.Time
	Location    string
	Description string
}

func (ss *SessionStore) ListEvents() ([]Event, error) {
	events, err := ss.querier.ListEvents(context.Background())
	if err != nil {
		return nil, err
	}
	var result []Event
	for _, event := range events {
		sanitiseEvent(&event)

		startTime := event.StartTime
		endTime := event.EndTime.Time
		event := Event{
			ID:          int(event.ID),
			Name:        event.Name,
			StartTime:   startTime,
			EndTime:     endTime,
			Location:    event.Location.String,
			Description: event.Description.String,
		}
		result = append(result, event)
	}
	return result, nil
}

func (ss *SessionStore) GetEvent(id int) (*Event, error) {
	event, err := ss.querier.GetEvent(context.Background(), int32(id))
	if err != nil {
		return nil, err
	}

	sanitiseEvent(&event)

	startTime := event.StartTime
	endTime := event.EndTime.Time
	return &Event{
		ID:          int(event.ID),
		Name:        event.Name,
		StartTime:   startTime,
		EndTime:     endTime,
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
