package models

import (
	"github.com/jinzhu/gorm"
)

type LoginModel struct {
	Model
}

// NewLoginModel returns a new UserModel with access to the database
func NewLoginModel(db *gorm.DB) *LoginModel {
	return &LoginModel{Model{database: db}}
}

// Post attempts to log in a user using the credentials given
func (m *LoginModel) Post(username string, password string) error {
	user := NewUserModel(m.database)
	err := user.Get(username)
	return err
}
