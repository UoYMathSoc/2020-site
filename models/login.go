package models

import (
	"golang.org/x/crypto/bcrypt"
)

type LoginModel struct {
	Model
}

func (m *LoginModel) Post(formParams map[string][]string) (err error) {
	username := formParams["username"][0]
	password := formParams["password"][0]

	user := getUser(username)

	return bcrypt.CompareHashAndPassword(user.HashedPassword, []byte(password))
}
