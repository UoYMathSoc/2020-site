package main

import (
	"github.com/UoYMathSoc/2020-site/controllers"
	"github.com/UoYMathSoc/2020-site/structs"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

type Server struct {
	*negroni.Negroni
}

func NewServer(c *structs.Config) (*Server, error) {

	s := Server{negroni.Classic()}

	router := mux.NewRouter().StrictSlash(true)

	getRouter := router.Methods("GET").Subrouter()
	//postRouter := router.Methods("POST").Subrouter()
	//headRouter := router.Methods("HEAD").Subrouter()

	// Routes go in here
	//loginC := controllers.NewLoginController(c)
	//postRouter.HandleFunc("/login/", loginC.Post)

	userC := controllers.NewUserController(c)
	getRouter.HandleFunc("/user/{username}", userC.Get)

	staticC := controllers.NewStaticController(c)
	getRouter.HandleFunc("/", staticC.GetIndex)
	getRouter.HandleFunc("/login/", staticC.GetLogin)
	// End routes

	s.UseHandler(router)

	return &s, nil

}
