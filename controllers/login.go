package controllers

import (
	"github.com/UoYMathSoc/2020-site/models"
	"github.com/UoYMathSoc/2020-site/structs"
	"net/http"
)

type LoginController struct {
	Controller
}

func NewLoginController(c *structs.Config) *LoginController {
	return &LoginController{Controller{config: c}}
}

func (loginC *LoginController) Pots(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	formParams := r.Form

	loginM := models.NewLoginModel(0)
	err := loginM.Post(formParams)
	if err != nil {

	}
}
