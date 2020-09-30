package controllers

import (
	"github.com/UoYMathSoc/2020-site/models"
	"net/http"

	"github.com/UoYMathSoc/2020-site/structs"
)

type Controller struct {
	config  *structs.Config
	session *models.Session
}

// ControllerInterface defines the methods all controllers must implement.
type ControllerInterface interface {
	Get()
	Post()
	Delete()
	Put()
	Head()
	Patch()
	Options()
}

// The following methods will be refused unless specifically overwritten.

func (c *Controller) Get(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
}

func (c *Controller) Post(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
}

func (c *Controller) Delete(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
}

func (c *Controller) Put(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
}

func (c *Controller) Head(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
}

func (c *Controller) Patch(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
}

func (c *Controller) Options(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
}
