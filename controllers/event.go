package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/UoYMathSoc/2020-site/models"
	"github.com/gorilla/mux"

	"github.com/UoYMathSoc/2020-site/structs"
)

type EventController struct {
	Controller
}

func NewEventController(c *structs.Config, s *models.Session) *EventController {
	return &EventController{Controller{config: c, session: s}}
}

func (eventC *EventController) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	event, err := eventC.session.GetEvent(id)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "could not find event", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "%s\n", event.Name)
	fmt.Fprintf(w, "%s-%s\n", event.StartDate.String(), event.EndDate.String())
	fmt.Fprintf(w, "Location: %s\n", event.Location)
	fmt.Fprintf(w, "%s/n", event.Description)
}
