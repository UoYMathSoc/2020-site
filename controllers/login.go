package controllers

import (
	"fmt"
	"github.com/UoYMathSoc/2020-site/database"
	"log"
	"net/http"

	"github.com/UoYMathSoc/2020-site/models"
	"github.com/UoYMathSoc/2020-site/structs"
)

type LoginController struct {
	controller
	Users models.UserStore
}

func NewLoginController(c *structs.Config, q database.Querier) *LoginController {
	us := models.NewUserStore(q)
	return &LoginController{controller: controller{config: c}, Users: us}
}

func (lc *LoginController) Post(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	id, err := lc.Users.Validate(username, password)
	if err != nil {
		log.Println(err)
		http.Error(w, "The password that you have entered is incorrect", http.StatusUnauthorized)
		return
	}
	user, err := lc.Users.Get(id)
	if err != nil {
		log.Panicln(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "Hello, %s", user.Name)
	lc.View.Render(w, lc.config.PageContext, nil)
}
