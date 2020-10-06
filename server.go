package main

import (
	"github.com/UoYMathSoc/2020-site/controllers"
	"github.com/UoYMathSoc/2020-site/models"
	"github.com/UoYMathSoc/2020-site/structs"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

// Server is the type of the main 2020site web application.
type Server struct {
	*negroni.Negroni
}

// NewServer creates a 2020site server based on the config c.
func NewServer(c *structs.Config) (*Server, error) {

	s := Server{negroni.Classic()}

	router := mux.NewRouter().StrictSlash(true)

	getRouter := router.Methods("GET").Subrouter()
	postRouter := router.Methods("POST").Subrouter()
	//headRouter := router.Methods("HEAD").Subrouter()

	session := models.NewSessionFromConfig(c.Database)

	// Routes go in here
	loginC := controllers.NewLoginController(c, session)
	postRouter.HandleFunc("/login/", loginC.Post)

	userC := controllers.NewUserController(c, session)
	getRouter.HandleFunc("/user/{id}", userC.Get)

	eventC := controllers.NewEventController(c, session)
	getRouter.HandleFunc("/events/{id}", eventC.Get)

	calendarC := controllers.NewCalendarController(c, session)
	getRouter.HandleFunc("/calendar/ical/MathSoc.ics", calendarC.GetICal)

	staticC := controllers.NewStaticController(c)
	getRouter.HandleFunc("/", staticC.GetIndex)
	getRouter.HandleFunc("/about/", staticC.GetAbout)
	getRouter.HandleFunc("/committee", staticC.GetCommittee)
	getRouter.HandleFunc("/contact", staticC.GetContact)
	getRouter.HandleFunc("/login/", staticC.GetLogin)
	// End routes

	s.UseHandler(router)

	return &s, nil

}
