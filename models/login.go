package models

import (
	"fmt"
	"gorm.io/gorm"
)

type LoginModel struct {
	Model
}

// NewLoginModel returns a new UserModel with access to the database
func NewLoginModel(db *gorm.DB) *LoginModel {
	return &LoginModel{Model: Model{database: db}}
}

// Post attempts to log in a user using the credentials given
func (m *LoginModel) Post(username string, password string) error {
	user := NewUserModel(m.database)
	err := user.Register(username, password)
	if err != nil {
		return fmt.Errorf("Unable to register user: %w", err)
	}
	err = user.Get(username)
	if err != nil {
		return err
	}
	err = user.Validate(password)
	if err != nil {
		return fmt.Errorf("could not validate user's credentails: %w", err)
	}
	return err
}
