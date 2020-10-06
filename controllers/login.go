package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/UoYMathSoc/2020-site/models"
	"github.com/UoYMathSoc/2020-site/structs"
)

type LoginController struct {
	Controller
}

func NewLoginController(c *structs.Config, s *models.Session) *LoginController {
	return &LoginController{Controller{config: c, session: s}}
}

func (loginC *LoginController) Post(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	id, err := loginC.session.ValidateUser(username, password)
	if err != nil {
		log.Println(err)
		http.Error(w, "The password that you have entered is incorrect", http.StatusUnauthorized)
		return
	}
	user, err := loginC.session.GetUser(id)
	if err != nil {
		log.Panicln(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "Hello, %s", user.Name)
}
