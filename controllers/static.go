package controllers

import (
	"net/http"

	"github.com/UoYMathSoc/2020-site/structs"
	"github.com/UoYMathSoc/2020-site/utils"
)

type StaticController struct {
	Controller
}

func NewStaticController(c *structs.Config) *StaticController {
	return &StaticController{Controller{config: c}}
}

func (staticC *StaticController) Get(w http.ResponseWriter, r *http.Request, content string) {
	err := utils.RenderContent(w, staticC.config.PageContext, nil, content)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (staticC *StaticController) GetIndex(w http.ResponseWriter, r *http.Request) {
	staticC.Get(w, r, "index.gohtml")
}

func (staticC *StaticController) GetAbout(w http.ResponseWriter, r *http.Request) {
	staticC.Get(w, r, "about.gohtml")
}

func (staticC *StaticController) GetContact(w http.ResponseWriter, r *http.Request) {
	staticC.Get(w, r, "contact.gohtml")
}

func (staticC *StaticController) GetLogin(w http.ResponseWriter, r *http.Request) {
	staticC.Get(w, r, "internal/login.gohtml")
}
