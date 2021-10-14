package controllers

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/UoYMathSoc/2020-site/database"
	"github.com/UoYMathSoc/2020-site/models"
	"github.com/UoYMathSoc/2020-site/structs"
	"github.com/UoYMathSoc/2020-site/views"
	"github.com/joncalhoun/form"
	"golang.org/x/oauth2"
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

func (ac *AdminController) activeUser(w http.ResponseWriter, r *http.Request) (string, error) {
	session, err := ac.Sessions.Get(r, sessionStoreKey)
	if err != nil {
		return "", fmt.Errorf("could not get session from request: %w", err)
	}
	username, ok := session.Values["username"].(string)
	if !ok {
		token, ok := session.Values["googleAccessToken"].(*oauth2.Token)
		if !ok {
			return "", fmt.Errorf("could not find google access token: %w", err)
		}
		username, err = ac.Sessions.GetUsername(token)
		if err != nil {
			return "", fmt.Errorf("could not get username from token: %w", err)
		}
		session.Values["username"] = username
		session.Save(r, w)
	}
	return username, nil
}

func (ac *AdminController) Get(w http.ResponseWriter, r *http.Request) {
	username, err := ac.activeUser(w, r)
	if err != nil {
		// Not logged in
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	user, _, err := ac.Users.GetByUsername(username)
	if err != nil {
		// Could not get user data from database
		http.Error(w, err.Error(), http.StatusInternalServerError)
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

var inputTpl = `
<div class="form-group">
	<label {{with .ID}}for="{{.}}"{{end}}>
		{{.Label}}
	</label>
	{{if eq .Type "textarea"}}
		<textarea class="form-control" {{with .ID}}id="{{.}}"{{end}} name="{{.Name}}" rows="3" placeholder="{{.Placeholder}}">{{with .Value}}{{.}}{{end}}</textarea>
	{{else}}
		<input type="{{.Type}}" class="form-control" {{with .ID}}id="{{.}}"{{end}} name="{{.Name}}" placeholder="{{.Placeholder}}" {{with .Value}}value="{{.}}"{{end}}>
	{{end}}
	{{with .Footer}}
		<small class="form-text text-muted">
			{{.}}
		</small>
	{{end}}
</div>`

type Form func(ac *AdminController) (name string, defaults Response)

type Response interface {
	Do(ac *AdminController) error
}

func (ac *AdminController) GetForm(f Form) http.HandlerFunc {
	tpl := template.Must(template.New("").Parse(inputTpl))
	fb := form.Builder{
		InputTemplate: tpl,
	}
	view := views.New("base", "action", "adminbar").Funcs(fb.FuncMap())
	return func(w http.ResponseWriter, r *http.Request) {
		user, _, err := ac.Users.GetByUsername("jgd511")
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		name, defaults := f(ac)
		data := struct {
			ActiveUser models.User
			Title      string
			Defaults   interface{}
		}{
			ActiveUser: user,
			Title:      name,
			Defaults:   defaults,
		}
		err = view.Render(w, ac.config.PageContext, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

// func (ac *AdminController) PostForm(f Form) http.HandlerFunc {

// }
