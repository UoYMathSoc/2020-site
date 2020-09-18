package models

import (
	"fmt"
	"github.com/UoYMathSoc/2020-site/database"
)

type LoginModel struct {
	Model
}

// NewLoginModel returns a new UserModel with access to the database
func NewLoginModel(q *database.Queries) *LoginModel {
	return &LoginModel{Model{q}}
}

// Post attempts to log in a user using the credentials given
func (m *LoginModel) Post(username string, password string) error {
	userM := NewUserModel(m.querier)
	_, err := userM.Validate(username, password)
	if err != nil {
		return fmt.Errorf("could not validate user's credentails: %w", err)
	}
	return nil
}
