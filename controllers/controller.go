package controllers

import (
	"github.com/UoYMathSoc/2020-site/models"
	"github.com/UoYMathSoc/2020-site/views"
	"net/http"

	"github.com/UoYMathSoc/2020-site/structs"
)

type controller struct {
	config   *structs.Config
	Sessions *models.SessionStore
	View     *views.View
}

// Controller defines the methods all controllers must implement.
type Controller interface {
	Get()
	Post()
	Delete()
	Put()
	Head()
	Patch()
	Options()
}

// The following methods will be refused unless specifically overwritten.

func (c *controller) Get(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
}

func (c *controller) Post(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
}

func (c *controller) Delete(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
}

func (c *controller) Put(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
}

func (c *controller) Head(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
}

func (c *controller) Patch(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
}

func (c *controller) Options(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
}
