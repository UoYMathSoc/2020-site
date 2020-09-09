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
func (m *LoginModel) Post(username string, password string) (err error) {
	user := NewUserModel(m.database)
	user.Get(username)
	return nil
}
func (m *LoginModel) Post(formParams map[string][]string) (err error) {
	username := "jgd511"
	password := "password"
	NewUser(username, password)
	return
}

//func (m *LoginModel) Post(formParams map[string][]string) (err error) {
//	username := formParams["username"][0]
//	password := formParams["password"][0]
//
//	user := read(username)
//
//	return bcrypt.CompareHashAndPassword(user.HashedPassword, []byte(password))
//}
