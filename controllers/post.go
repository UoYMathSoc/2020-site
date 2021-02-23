package controllers

import (
	"fmt"
	"github.com/UoYMathSoc/2020-site/database"
	"github.com/UoYMathSoc/2020-site/models"
	"github.com/UoYMathSoc/2020-site/structs"
	"github.com/UoYMathSoc/2020-site/views"
	"github.com/gorilla/mux"
	"net/http"
)

type PostController struct {
	controller
	Posts models.PostStore
}

func NewPostController(c *structs.Config, q database.Querier) *PostController {
	ps := models.NewPostStore(q)
	ss := models.NewSessionStore(c, q)
	v := views.New("base", "post", "navbar")
	con := controller{
		config:   c,
		Sessions: ss,
		View:     v,
	}
	return &PostController{controller: con, Posts: ps}
}

func (pc *PostController) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]

	post, err := pc.Posts.Get(key)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := struct {
		Post *models.Post
	}{
		Post: post,
	}

	err = pc.View.Render(w, pc.config.PageContext, data)
	if err != nil {
		fmt.Println(err)
		return
	}
}
