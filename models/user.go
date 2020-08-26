package models

import (
	"fmt"
	"log"

	"github.com/UoYMathSoc/2020-site/database"
	"golang.org/x/crypto/bcrypt"
)

type UserModel struct {
	Model
	Username       string
	HashedPassword []byte
}

func NewUserModel() *UserModel {
	return new(UserModel)
}

func NewUser(username string, password string) (*UserModel, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	user := &UserModel{
		Username:       username,
		HashedPassword: hashedPassword,
	}
	user.create()
	if _, err := user.Get(username); err != nil {
		err := fmt.Errorf("User: %s could not be registered", username)
		return nil, err
	}
	return user, nil
}

func (user *UserModel) Get(username string) (*UserModel, error) {
	if user.read(username).Username == username {
		return user, nil
	}
	return nil, fmt.Errorf("Database does not contain User: %s", username)
}

func (user *UserModel) create() *UserModel {
	db := database.Instance
	db.Create(user)
	return user
}

func (user *UserModel) read(username string) *UserModel {
	db := database.Instance
	db.Where("username = ?", username).Find(&user)
	return user
}

func (user *UserModel) update(username string) *UserModel {
	db := database.Instance
	db.Where("username = ?", username).Updates(user)
	return user
}

func (user *UserModel) delete() *UserModel {
	db := database.Instance
	db.Where("username = ?", user.Username).Delete(user)
	return user
}
