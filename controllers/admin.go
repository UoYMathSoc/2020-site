package controllers

import (
	"github.com/UoYMathSoc/2020-site/database"
	"github.com/UoYMathSoc/2020-site/models"
	"github.com/UoYMathSoc/2020-site/structs"
	"github.com/UoYMathSoc/2020-site/views"
	"golang.org/x/oauth2"
	"net/http"
)

type AdminController struct {
	controller
	Users models.UserStore
}

func NewAdminController(c *structs.Config, q database.Querier) *AdminController {
	us := models.NewUserStore(q)
	ss := models.NewSessionStore(c, q)
	return &AdminController{controller: controller{
		config:   c,
		Sessions: ss,
		View:     views.New("base", "admin", "adminbar"),
	}, Users: us}
}

func (ac *AdminController) Get(w http.ResponseWriter, r *http.Request) {
	//TODO: move this somewhere as it is likely to be called at the start of all admin pages
	session, err := ac.Sessions.Get(r, sessionStoreKey)
	if err != nil {
		//TODO: Feed in some 'Not logged in' message
		http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
		return
	}
	username, ok := session.Values["username"].(string)
	if !ok {
		token, ok := session.Values["googleAccessToken"].(*oauth2.Token)
		if !ok {
			//TODO: Feed in some 'Not logged in' message
			http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
			return
		}
		username, err = ac.Sessions.GetUsername(token)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		session.Values["username"] = username
		session.Save(r, w)
	}

	user, _, err := ac.Users.GetByUsername(username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	data := struct {
		ActiveUser models.User
	}{
		ActiveUser: user,
	}
	err = ac.View.Render(w, ac.config.PageContext, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
