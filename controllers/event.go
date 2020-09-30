package controllers

import (
	"fmt"
	"github.com/UoYMathSoc/2020-site/models"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"time"

	"github.com/UoYMathSoc/2020-site/database"
	"github.com/UoYMathSoc/2020-site/structs"
)

type EventController struct {
	Controller
}

func NewEventController(c *structs.Config, q *database.Queries) *EventController {
	return &EventController{Controller{config: c, querier: q}}
}

func (eventC *EventController) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	event, err := eventC.session.GetEvent(id)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "%s\n", event.Name)
	fmt.Fprintf(w, "%s-%s\n", event.StartDate.String(), event.EndDate.String())
	fmt.Fprintf(w, "Location: %s\n", event.Location)
	fmt.Fprintf(w, "%s/n", event.Description)

	//fmt.Fprintf(w, "%s\n", event.Name)
	//fmt.Fprintf(w, "We are going to %s on %s, so get ready!\n", event.Location.String, event.Date.Weekday())
	//fmt.Fprintf(w, "%s\n", event.Description.String)
	//fmt.Fprintf(w, "%d:%d-%d:%d", event.StartTime.Hour(), event.StartTime.Minute(), event.EndTime.Time.Hour(), event.EndTime.Time.Minute())
}
