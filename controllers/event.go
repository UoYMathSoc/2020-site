package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/UoYMathSoc/2020-site/models"
	"github.com/gorilla/mux"

	"github.com/UoYMathSoc/2020-site/structs"
)

type EventController struct {
	controller
}

func NewEventController(c *structs.Config, s *models.SessionStore) *EventController {
	return &EventController{controller{config: c, Sessions: s}}
}

func (eventC *EventController) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	event, err := eventC.Sessions.GetEvent(id)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "could not find event", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "%s\n", event.Name)
	if event.StartTime.Truncate(time.Hour*24) == event.EndTime.Truncate(time.Hour*24) { // Same day
		fmt.Fprintf(w, "%s - %s\n", event.StartTime.Format("2 Jan, 15:04"), event.EndTime.Format("15:04"))
	} else {
		fmt.Fprintf(w, "%s - %s\n", event.StartTime.Format("2 Jan, 15:04"), event.EndTime.Format("2 Jan, 15:04"))
	}
	fmt.Fprintf(w, "Location: %s\n", event.Location)
	fmt.Fprintf(w, "%s/n", event.Description)
}
