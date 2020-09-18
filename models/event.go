package models

import (
	"context"
	"github.com/UoYMathSoc/2020-site/database"
	"github.com/randoomjd/goics"
)

type EventModel struct {
	Model
}

func NewEventModel(q *database.Queries) *EventModel {
	return &EventModel{Model{q}}
}

func (m EventModel) List() (Events, error) {
	return m.querier.ListEvents(context.Background())
}

type Events []database.Event

func (events Events) EmitICal() goics.Componenter {
	c := goics.NewComponent()
	c.SetType("VCALENDAR")
	c.AddProperty("CALSCAL", "GREGORIAN")

	for _, event := range events {
		s := goics.NewComponent()
		s.SetType("VEVENT")
		k, v := goics.FormatDateTimeField("DTEND", event.EndTime.Time)
		s.AddProperty(k, v)
		k, v = goics.FormatDateTimeField("DTSTART", event.StartTime)
		s.AddProperty(k, v)
		domain := "yums.org.uk"
		s.AddProperty("UID", event.Key+"@calendar."+domain)
		s.AddProperty("SUMMARY", event.Name)
		s.AddProperty("LOCATION", event.Location.String)
		s.AddProperty("URL", domain+"/events/"+event.Key)
		c.AddComponent(s)
	}
	return c
}
