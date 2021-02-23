package controllers

import (
	"fmt"
	"net/http"

	"github.com/UoYMathSoc/2020-site/database"
	"github.com/UoYMathSoc/2020-site/models"
	"github.com/UoYMathSoc/2020-site/structs"
	"github.com/UoYMathSoc/2020-site/views"
)

type CommitteeController struct {
	controller
	Users models.UserStore
}

func NewCommitteeController(c *structs.Config, q database.Querier) *CommitteeController {
	us := models.NewUserStore(q)
	ss := models.NewSessionStore(c, q)
	v := views.New("base", "committee", "navbar")
	con := controller{
		config:   c,
		Sessions: ss,
		View:     v,
	}
	return &CommitteeController{controller: con, Users: us}
}

func (cc *CommitteeController) Get(w http.ResponseWriter, r *http.Request) {
	exec, committee, err := cc.Users.GetCommittee()
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := struct {
		Exec      models.Committee
		Committee models.Committee
	}{
		Exec:      exec,
		Committee: committee,
	}
	err = cc.View.Render(w, cc.config.PageContext, data)
	if err != nil {
		fmt.Println(err)
	}
}
