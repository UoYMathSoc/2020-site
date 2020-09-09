package models

import (
	"errors"

	"github.com/jinzhu/gorm"
)

// UserModel is used when accessing the site's users
type UserModel struct {
	Model
	Username       string
	HashedPassword []byte
}

// NewUserModel returns a new UserModel with access to the database
func NewUserModel(db *gorm.DB) *UserModel {
	model := Model{database: db}
	return &UserModel{Model: model, Username: "", HashedPassword: []byte("")}
}

// Get populates a user's UserModel. Uses user's username unless one is provided as a parameter
func (user *UserModel) Get(usernames ...string) (*UserModel, error) {
	username := user.Username
	if len(usernames) > 0 {
		username = usernames[0]
	}
	user.getUser()
	if username == user.Username {
		return user, nil
	}
	return nil, errors.New("User does note exist in database")
}

func (user *UserModel) getUsers() (users *[]UserModel) {
	user.database.Find(users)
	for _, thisUser := range *users {
		thisUser.database = user.database
	}
	return users
}

func (user *UserModel) newUser() *UserModel {
	user.database.Create(user)
	return user
}

func (user *UserModel) getUser() *UserModel {
	user.database.Where("username = ?", user.Username).Find(&user)
	return user
}

func (user *UserModel) updateUser() *UserModel {
	user.database.Where("username = ?", user.Username).Updates(user)
	return user
}

func (user *UserModel) deleteUser() *UserModel {
	user.database.Where("username = ?", user.Username).Delete(UserModel{})
	return user
}
