package controllers

import (
	"fmt"
	"net/http"

	"github.com/UoYMathSoc/2020-site/models"
	"github.com/UoYMathSoc/2020-site/structs"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type UserController struct {
	Controller
}

// NewUserController creates a new 'null' user controller
func NewUserController(c *structs.Config, db *gorm.DB) *UserController {
	return &UserController{Controller{config: c, database: db}}
}

func (userC *UserController) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]

	userM := models.NewUserModel(userC.database)
	err := userM.Get(username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Hello, %s", userM.Username)
}
