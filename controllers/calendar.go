package controllers

import (
	"bytes"
	"net/http"

	"github.com/UoYMathSoc/2020-site/database"
	"github.com/UoYMathSoc/2020-site/models"
	"github.com/UoYMathSoc/2020-site/structs"
	"github.com/randoomjd/goics"
)

type CalendarController struct {
	Controller
}

func NewCalendarController(c *structs.Config, q *database.Queries) *CalendarController {
	return &CalendarController{Controller{config: c, querier: q}}
}

func (calendarC *CalendarController) GetICal(w http.ResponseWriter, r *http.Request) {
	eventM := models.NewEventModel(calendarC.querier)
	events, err := eventM.List()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-type", "text/calendar")
	w.Header().Set("charset", "utf-8")
	w.Header().Set("Content-Disposition", "inline")
	w.Header().Set("filename", "mathsoc.ics")

	b := bytes.Buffer{}
	goics.NewICalEncode(&b).Encode(events)
	w.WriteHeader(http.StatusOK)
	w.Write(b.Bytes())
}
