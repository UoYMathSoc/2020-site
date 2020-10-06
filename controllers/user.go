package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/UoYMathSoc/2020-site/utils"
	"github.com/UoYMathSoc/2020-site/database"
	"github.com/UoYMathSoc/2020-site/models"
	"github.com/UoYMathSoc/2020-site/structs"
	"github.com/gorilla/mux"
)

type UserController struct {
	Controller
}

// NewUserController creates a new 'null' user controller
func NewUserController(c *structs.Config, s *models.Session) *UserController {
	return &UserController{Controller{config: c, session: s}}
}

func (userC *UserController) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r){
  id, _ := strconv.Atoi(vars["id"])
  
	user, err := userC.session.GetUser(id)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	positions, err := userC.session.GetUserPositions(id)
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

	err = utils.RenderContent(w, userC.config.PageContext, data, "user.gohtml")
	if err != nil {
		fmt.Println(err)
		return
	}
}
