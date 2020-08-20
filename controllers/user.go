package controllers

import (
	"fmt"
	"github.com/UoYMathSoc/2020-site/models"
	"github.com/UoYMathSoc/2020-site/structs"
	"github.com/gorilla/mux"
	"net/http"
)

type UserController struct {
	Controller
}

func NewUserController(c *structs.Config) *UserController {
	return &UserController{Controller{config: c}}
}

func (userC *UserController) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]

	userM, err := models.GetUser(username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Hello, %s", userM.Username)
}
