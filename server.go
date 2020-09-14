package main

import (
	"github.com/UoYMathSoc/2020-site/controllers"
	"github.com/UoYMathSoc/2020-site/structs"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
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

	db := InitDatabase(c)

	// Routes go in here
	loginC := controllers.NewLoginController(c, db)
	postRouter.HandleFunc("/login/", loginC.Post)

	userC := controllers.NewUserController(c, db)
	getRouter.HandleFunc("/user/{username}", userC.Get)

	calendarC := controllers.NewCalendarController(c, db)
	getRouter.HandleFunc("/calendar/ical", calendarC.GetICal)

	staticC := controllers.NewStaticController(c)
	getRouter.HandleFunc("/", staticC.GetIndex)
	getRouter.HandleFunc("/login/", staticC.GetLogin)
	// End routes

	s.UseHandler(router)

	return &s, nil

}
