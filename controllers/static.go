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

func (staticC *StaticController) GetIndex(w http.ResponseWriter, r *http.Request) {
	err := utils.RenderContent(w, staticC.config.PageContext, nil, "index.gohtml")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (staticC *StaticController) GetLogin(w http.ResponseWriter, r *http.Request) {
	err := utils.RenderTemplates(w, staticC.config.PageContext, nil, "internal/login.gohtml", "elements/adminbar.gohtml")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
