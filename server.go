package main

import (
	"database/sql"
	"fmt"

	"github.com/UoYMathSoc/2020-site/controllers"
	"github.com/UoYMathSoc/2020-site/database"
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

	q := NewDBFromConfig(c.Database)
	ss := models.NewSessionStore(c, q)

	// Routes go in here
	loginC := controllers.NewLoginController(c, q)
	getRouter.HandleFunc("/login", loginC.Get)
	getRouter.HandleFunc("/callback-gl", loginC.Callback)
	getRouter.HandleFunc("/google", loginC.Google)
	getRouter.HandleFunc("/destroy", loginC.Destroy)
	postRouter.HandleFunc("/login", loginC.Post)

	adminC := controllers.NewAdminController(c, q)
	getRouter.HandleFunc("/admin", adminC.Get)

	userC := controllers.NewUserController(c, q)
	getRouter.HandleFunc("/user/{id}", userC.Get)

	committeeC := controllers.NewCommitteeController(c, q)
	getRouter.HandleFunc("/committee", committeeC.Get)

	eventC := controllers.NewEventController(c, ss)
	getRouter.HandleFunc("/event/{id}", eventC.Get)

	postC := controllers.NewPostController(c, q)
	getRouter.HandleFunc("/post/{key}", postC.Get)

	calendarC := controllers.NewCalendarController(c, ss)
	getRouter.HandleFunc("/calendar/ical/MathSoc.ics", calendarC.GetICal)

	staticC := controllers.NewStaticController(c)
	getRouter.HandleFunc("/", staticC.GetIndex)
	getRouter.HandleFunc("/about", staticC.GetAbout)
	getRouter.HandleFunc("/contact", staticC.GetContact)
	// End routes

	s.UseHandler(router)

	return &s, nil

}

func NewDBFromConfig(db structs.Database) *database.Queries {
	dsn := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		db.Host, db.Port, db.User, db.Password, db.Name)
	conn, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

	return database.New(conn)
}
