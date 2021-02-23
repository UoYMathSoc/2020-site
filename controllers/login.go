package controllers

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/gob"
	"fmt"
	"golang.org/x/oauth2"
	"log"
	"net/http"

	"github.com/UoYMathSoc/2020-site/database"
	"github.com/UoYMathSoc/2020-site/models"
	"github.com/UoYMathSoc/2020-site/structs"
	"github.com/UoYMathSoc/2020-site/views"
)

// Important: Allow token to be saved in session data
func init() {
	gob.Register(&oauth2.Token{})
}

type LoginController struct {
	controller
	Users models.UserStore
}

func NewLoginController(c *structs.Config, q database.Querier) *LoginController {
	us := models.NewUserStore(q)
	ss := models.NewSessionStore(c, q)
	return &LoginController{controller: controller{
		config:   c,
		Sessions: ss,
		View:     views.New("base", "login", "adminbar"),
	}, Users: us}
}

func (lc *LoginController) Get(w http.ResponseWriter, r *http.Request) {
	var user models.User
	session, err := lc.Sessions.Get(r, sessionStoreKey)
	if err == nil {
		username, ok := session.Values["username"].(string)
		if ok {
			user, _, _ = lc.Users.GetByUsername(username)
		}
	}

	data := struct {
		ActiveUser models.User
		LoggedIn   bool
	}{
		ActiveUser: user,
		LoggedIn:   user.Username != "",
	}
	err = lc.View.Render(w, lc.config.PageContext, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (lc *LoginController) Post(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	_, err := lc.Users.Validate(username, password)
	if err != nil {
		log.Println(err)
		http.Error(w, "The password that you have entered is incorrect", http.StatusUnauthorized)
		return
	}

	http.Redirect(w, r, "/admin", http.StatusTemporaryRedirect)
}

func (lc *LoginController) Callback(w http.ResponseWriter, r *http.Request) {
	session, err := lc.Sessions.Get(r, sessionStoreKey)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	if r.URL.Query().Get("state") != session.Values["state"] {
		http.Error(w, "Invalid State", http.StatusInternalServerError)
	}
	token, err := lc.Sessions.GenToken(r.URL.Query().Get("code"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	if !token.Valid() {
		http.Error(w, "Invalid Token", http.StatusInternalServerError)
	}

	username, err := lc.Sessions.GetUsername(token)
	if err != nil {
		http.Error(w, "Invalid Token", http.StatusInternalServerError)
	}

	session.Values["username"] = username
	session.Values["googleAccessToken"] = token
	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/admin", http.StatusTemporaryRedirect)
}

const sessionStoreKey = "MathSocSession"

func (lc *LoginController) Google(w http.ResponseWriter, r *http.Request) {
	b := make([]byte, 16)
	rand.Read(b)

	state := base64.URLEncoding.EncodeToString(b)

	session, _ := lc.Sessions.Get(r, sessionStoreKey)
	session.Values["state"] = state
	session.Save(r, w)

	url := lc.Sessions.AuthCodeURL(state)
	http.Redirect(w, r, url, http.StatusFound)
}

func (lc *LoginController) Destroy(w http.ResponseWriter, r *http.Request) {
	session, err := lc.Sessions.Get(r, sessionStoreKey)
	if err != nil {
		fmt.Fprintln(w, "aborted")
		return
	}

	session.Options.MaxAge = -1

	session.Save(r, w)
	http.Redirect(w, r, "/", 302)

}
