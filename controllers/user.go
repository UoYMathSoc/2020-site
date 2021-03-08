package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/UoYMathSoc/2020-site/database"
	"github.com/UoYMathSoc/2020-site/models"
	"github.com/UoYMathSoc/2020-site/structs"
	"github.com/UoYMathSoc/2020-site/views"
	"github.com/gorilla/mux"
)

type UserController struct {
	controller
	Users models.UserStore
}

// NewUserController creates a new 'null' user controller
func NewUserController(c *structs.Config, q database.Querier) *UserController {
	us := models.NewUserStore(q)
	ss := models.NewSessionStore(c, q)
	v := views.New("base", "user", "navbar")
	con := controller{
		config:   c,
		Sessions: ss,
		View:     v,
	}
	return &UserController{controller: con, Users: us}
}

func (uc *UserController) Post(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, _ := strconv.Atoi(vars["id"])
	_ = models.User{
		ID:       id,
		Username: "user" + vars["id"],
		Name:     "Mr. " + vars["id"],
		Bio:      "I am not a proper user.",
	}
	//_, err := uc.Users.Create(&user)
	//if err != nil {
	//	return
	//}
	uc.View.Render(w, uc.config.PageContext, nil)
}

func (uc *UserController) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	user, positions, err := uc.Users.Get(id)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if len(positions) == 0 {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	data := struct {
		User      models.User
		Positions []models.Position
	}{
		User:      user,
		Positions: positions.ByDate(),
	}

	err = uc.View.Render(w, uc.config.PageContext, data)
	//err = utils.RenderContent(w, uc.config.PageContext, data, "user.gohtml")
	if err != nil {
		fmt.Println(err)
	}
}
