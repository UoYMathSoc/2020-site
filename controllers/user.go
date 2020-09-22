package controllers

import (
	"fmt"
	"github.com/UoYMathSoc/2020-site/utils"
	"net/http"
	"strconv"

	"github.com/UoYMathSoc/2020-site/database"
	"github.com/UoYMathSoc/2020-site/models"
	"github.com/UoYMathSoc/2020-site/structs"
	"github.com/gorilla/mux"
)

type UserController struct {
	Controller
}

// NewUserController creates a new 'null' user controller
func NewUserController(c *structs.Config, q *database.Queries) *UserController {
	return &UserController{Controller{config: c, querier: q}}
}

func (userC *UserController) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	userM := models.NewUserModel(userC.querier)
	user, positions, err := userM.Get(int32(id))
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
		User      *database.User
		Positions []database.Position
	}{
		User:      user,
		Positions: positions,
	}

	err = utils.RenderContent(w, userC.config.PageContext, data, "user.gohtml")
	if err != nil {
		fmt.Println(err)
		return
	}
}
