package controllers

import (
	"log"
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
	err := utils.RenderContent(w, staticC.config.PageContext, nil, "index.tmpl")
	if err != nil {
		log.Println(err)
		return
	}
}

func (staticC *StaticController) GetLogin(w http.ResponseWriter, r *http.Request) {
	err := utils.RenderTemplates(w, staticC.config.PageContext, nil, "admin/login.tmpl", "elements/adminbar.tmpl")
	if err != nil {
		log.Println(err)
		return
	}
}
