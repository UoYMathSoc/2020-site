package controllers

import (
	"bytes"
	"github.com/UoYMathSoc/2020-site/models"
	"log"

	"github.com/UoYMathSoc/2020-site/structs"
	"github.com/randoomjd/goics"
	"gorm.io/gorm"
	"net/http"
)

type CalendarController struct {
	Controller
}

func NewCalendarController(c *structs.Config, db *gorm.DB) *CalendarController {
	return &CalendarController{Controller{config: c, database: db}}
}

func (calendarC *CalendarController) GetICal(w http.ResponseWriter, r *http.Request) {
	eventM := models.NewEventModel(calendarC.database)
	events, err := eventM.GetEvents()
	if err != nil {
		log.Println(err)
		http.Error(w, "Calendar could not be generated. Please try again later.", http.StatusInternalServerError)
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
