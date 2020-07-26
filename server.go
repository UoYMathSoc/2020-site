package main

import (
	//"fmt"
	//"net/http"
	//"strconv"

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
	//postRouter := router.Methods("POST").Subrouter()
	//headRouter := router.Methods("HEAD").Subrouter()
	
	// Routes go in here
	staticC := controllers.NewStaticController(c)
	getRouter.HandleFunc("/", staticC.GetIndex)
	// End routes

	s.UseHandler(router)

	return &s, nil
	
}
