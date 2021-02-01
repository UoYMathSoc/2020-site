package controllers

import (
	"fmt"
	"github.com/UoYMathSoc/2020-site/database"
	"github.com/UoYMathSoc/2020-site/views"
	"net/http"
	"strconv"

	"github.com/UoYMathSoc/2020-site/models"
	"github.com/UoYMathSoc/2020-site/structs"
	"github.com/gorilla/mux"
)

type UserController struct {
	controller
	Users models.UserStore
}

// NewUserController creates a new 'null' user controller
func NewUserController(c *structs.Config, q database.Querier) *UserController {
	us := models.NewUserStore(q)
	ss := models.NewSessionStore(q)
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
	user := models.User{
		ID:       id,
		Username: "user" + vars["id"],
		Name:     "Mr. " + vars["id"],
		Bio:      "I am not a proper user.",
	}
	err := uc.Users.Create(&user)
	if err != nil {
		return
	}
	uc.View.Render(w, uc.config.PageContext, nil)
}

func (uc *UserController) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		return
	}

	user, err := uc.Users.Get(id)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	positions, err := uc.Users.GetPositions(id)
	if len(positions) == 0 {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := struct {
		User      *models.User
		Positions []models.Position
	}{
		User:      user,
		Positions: positions,
	}

	err = uc.View.Render(w, uc.config.PageContext, data)
	//err = utils.RenderContent(w, uc.config.PageContext, data, "user.gohtml")
	if err != nil {
		fmt.Println(err)
		return
	}
}
