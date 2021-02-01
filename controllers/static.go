package controllers

import (
	"github.com/UoYMathSoc/2020-site/views"
	"net/http"

	"github.com/UoYMathSoc/2020-site/structs"
)

type StaticController struct {
	controller
	views map[string]*views.View
}

func NewStaticController(c *structs.Config) *StaticController {
	return &StaticController{controller: controller{config: c}, views: map[string]*views.View{}}
}

func (sc *StaticController) get(w http.ResponseWriter, content string) {
	if _, ok := sc.views[content]; !ok {
		sc.views[content] = views.New("base", content, "navbar")
	}
	sc.views[content].Render(w, sc.config.PageContext, nil)
}

func (sc *StaticController) GetIndex(w http.ResponseWriter, r *http.Request) {
	sc.get(w, "index")
}

func (sc *StaticController) GetAbout(w http.ResponseWriter, r *http.Request) {
	sc.get(w, "about")
}

func (sc *StaticController) GetCommittee(w http.ResponseWriter, r *http.Request) {
	sc.get(w, "committee")
}

func (sc *StaticController) GetContact(w http.ResponseWriter, r *http.Request) {
	sc.get(w, "contact")
}

func (sc *StaticController) GetLogin(w http.ResponseWriter, r *http.Request) {
	if _, ok := sc.views["login"]; !ok {
		sc.views["login"] = views.New("base", "login", "adminbar")
	}
	sc.views["login"].Render(w, sc.config.PageContext, nil)
}
