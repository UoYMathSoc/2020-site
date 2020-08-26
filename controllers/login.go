package controllers

import (
	"fmt"
	"github.com/UoYMathSoc/2020-site/structs"
	"net/http"
)

type LoginController struct {
	Controller
}

func NewLoginController(c *structs.Config) *LoginController {
	return &LoginController{Controller{config: c}}
}

func (loginC *LoginController) Post(w http.ResponseWriter, r *http.Request) {
	//err := r.ParseForm()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//log.Printf("%+v\n", r.Form)
	//formParams := r.Form
	//
	//loginM := models.NewLoginModel()
	//err = loginM.Post(formParams)
	//if err != nil {
	//	w.WriteHeader(http.StatusUnauthorized)
	//}
	fmt.Fprintf(w, "Login Endpoint Hit")
}
