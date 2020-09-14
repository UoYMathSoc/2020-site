package models

import (
	"time"

	"github.com/randoomjd/goics"
	"gorm.io/gorm"
)

type EventModel struct {
	Model
	key      string
	name     string
	time     time.Time
	duration time.Duration
	location string
}

type Events []EventModel

var (
	Event1 = EventModel{
		key:      "test1",
		name:     "Test Event",
		time:     time.Now(),
		duration: time.Hour,
		location: "here",
	}

	Event2 = EventModel{
		key:      "test2",
		name:     "Tested Event",
		time:     time.Now().Add(time.Hour),
		duration: time.Hour,
		location: "there",
	}
)

func NewEventModel(db *gorm.DB) *EventModel {
	return &EventModel{Model: Model{database: db}}
}

func (m *EventModel) GetEvents() (*Events, error) {
	//var events []EventModel
	//m.database.Find(&events)
	//for _, event := range events {
	//	event.database = m.database
	//}
	//return &events, nil
	return &Events{Event1, Event2}, nil
}

func (events Events) EmitICal() goics.Componenter {
	c := goics.NewComponent()
	c.SetType("VCALENDAR")
	c.AddProperty("CALSCAL", "GREGORIAN")

	for _, event := range events {
		s := goics.NewComponent()
		s.SetType("VEVENT")
		dtend := event.time.Add(event.duration)
		k, v := goics.FormatDateTimeField("DTEND", dtend)
		s.AddProperty(k, v)
		k, v = goics.FormatDateTimeField("DTSTART", event.time)
		s.AddProperty(k, v)
		domain := "yums.org.uk"
		s.AddProperty("UID", event.key+"@calendar."+domain)
		s.AddProperty("SUMMARY", event.name)
		s.AddProperty("LOCATION", event.location)
		s.AddProperty("URL", domain+"/events/"+event.key)
		c.AddComponent(s)
	}
	return c
}
