package controllers

import (
	"log"
	"net/http"

	"github.com/UoYMathSoc/2020-site/database"
	"github.com/UoYMathSoc/2020-site/models"
	"github.com/UoYMathSoc/2020-site/structs"
)

type LoginController struct {
	Controller
}

func NewLoginController(c *structs.Config, q *database.Queries) *LoginController {
	return &LoginController{Controller{config: c, querier: q}}
}

func (loginC *LoginController) Post(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	loginM := models.NewLoginModel(loginC.querier)
	err := loginM.Post(username, password)
	if err != nil {
		log.Println(err)
		http.Error(w, "The password that you have entered is incorrect", http.StatusUnauthorized)
		return
	}

}
