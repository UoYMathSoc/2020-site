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
	err := utils.RenderTemplate(w, staticC.config.PageContext, nil, "index.tmpl")
	if err != nil {
		log.Println(err)
		return
	}
}
