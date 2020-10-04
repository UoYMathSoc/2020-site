package controllers

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/UoYMathSoc/2020-site/utils"

	"github.com/UoYMathSoc/2020-site/database"
	"github.com/UoYMathSoc/2020-site/models"
	"github.com/UoYMathSoc/2020-site/structs"
	"github.com/gorilla/mux"
)

type UserController struct {
	Controller
}

// NewUserController creates a new 'null' user controller
func NewUserController(c *structs.Config, q *database.Queries) *UserController {
	return &UserController{Controller{config: c, querier: q}}
}

func (userC *UserController) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var user *database.User
	var positions []database.Position
	var err error
	if vars["id"] != "test" {
		id, _ := strconv.Atoi(vars["id"])

		userM := models.NewUserModel(userC.querier)
		user, positions, err = userM.Get(int32(id))
		if len(positions) == 0 {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}

		if err != nil {
			fmt.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		user = &database.User{
			ID:       69,
			Username: "abc123",
			Name:     "Dick Testin",
			Bio:      sql.NullString{String: "I am a test user lol", Valid: true},
		}
		positions = append(positions, database.Position{
			ID:          69,
			Name:        sql.NullString{String: "Tester", Valid: true},
			Alias:       "tester",
			Ordering:    69,
			Description: sql.NullString{String: "Does a test", Valid: true},
			Executive:   sql.NullBool{Bool: false, Valid: true},
		},
		)
	}

	data := struct {
		User      *database.User
		Positions []database.Position
	}{
		User:      user,
		Positions: positions,
	}

	err = utils.RenderContent(w, userC.config.PageContext, data, "user.gohtml")
	if err != nil {
		fmt.Println(err)
		return
	}
}
