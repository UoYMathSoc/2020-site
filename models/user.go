package models

import (
<<<<<<< HEAD
	"errors"

	"github.com/jinzhu/gorm"
=======
	"fmt"
	"log"

	"github.com/UoYMathSoc/2020-site/database"
	"golang.org/x/crypto/bcrypt"
>>>>>>> 39e4d53fb0b04b54d765f2d71237fb145da2fb1c
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
	return nil, errors.New("User does not exist in database")
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

func (user *UserModel) NewUser(username string, password string) (*UserModel, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	user.Username = username
	user.HashedPassword = hashedPassword
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
