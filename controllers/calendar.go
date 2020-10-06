package controllers

import (
	"fmt"
	"net/http"

	"github.com/UoYMathSoc/2020-site/utils"

	"github.com/UoYMathSoc/2020-site/models"
	"github.com/UoYMathSoc/2020-site/structs"
)

type CalendarController struct {
	Controller
}

func NewCalendarController(c *structs.Config, s *models.Session) *CalendarController {
	return &CalendarController{Controller{config: c, session: s}}
}

func (calendarC *CalendarController) GetICal(w http.ResponseWriter, r *http.Request) {
	events, err := calendarC.session.ListEvents()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-type", "text/calendar")
	w.Header().Set("charset", "utf-8")
	w.Header().Set("Content-Disposition", "inline")
	w.Header().Set("filename", "mathsoc.ics")

	data := struct {
		Events []models.Event
	}{
		Events: events,
	}

	err = utils.RenderICal(w, data, "ical.tmpl")
	if err != nil {
		fmt.Println(err)
		return
	}
}
