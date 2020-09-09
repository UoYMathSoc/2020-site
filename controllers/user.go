package controllers

import (
	"fmt"
	"net/http"

	"github.com/UoYMathSoc/2020-site/models"
	"github.com/UoYMathSoc/2020-site/structs"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

type UserController struct {
	Controller
}

// NewUserController comment
func NewUserController(c *structs.Config, db *gorm.DB) *UserController {
	return &UserController{Controller{config: c, database: db}}
}

func (userC *UserController) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]

	userM := models.NewUserModel(userC.database)
	userM.Username = username
	userM.Get()

	fmt.Println(userM.Username + " ?= " + username)
	fmt.Fprintf(w, "Hello, %s", userM.Username)
}
