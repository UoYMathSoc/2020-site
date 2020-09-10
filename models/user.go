package models

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"log"

	"github.com/jinzhu/gorm"
)

// UserModel is used when accessing the site's users
type UserModel struct {
	Model
	Username       string
	hashedPassword []byte
}

// NewUserModel returns a new UserModel with access to the database
func NewUserModel(db *gorm.DB) *UserModel {
	model := Model{database: db}
	return &UserModel{Model: model, Username: "", hashedPassword: []byte("")}
}

// Get populates a user's UserModel. Uses user's username unless one is provided as a parameter
func (user *UserModel) Get(usernames ...string) error {
	username := user.Username
	if len(usernames) > 0 {
		username = usernames[0]
	}
	if user.read().Username == username {
		return nil
	}
	return errors.New("user does not exist in database")
}

//func (user *UserModel) GetUsers() (users *[]UserModel, err error) {
//	user.database.Find(users)
//	for _, thisUser := range *users {
//		thisUser.database = user.database
//	}
//	return users, err
//}

func (user *UserModel) NewUser(username string, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		log.Print(err)
		return err
	}
	user.Username = username
	user.hashedPassword = hashedPassword
	user.create()
	if err := user.Get(username); err != nil {
		err := errors.New("user could not be registered")
		return err
	}
	return nil
}

func (user *UserModel) create() *UserModel {
	user.database.Create(user)
	return user
}

func (user *UserModel) read() *UserModel {
	user.database.Where("username = ?", user.Username).Find(&user)
	return user
}

func (user *UserModel) update() *UserModel {
	user.database.Where("username = ?", user.Username).Updates(user)
	return user
}

func (user *UserModel) delete() *UserModel {
	user.database.Where("username = ?", user.Username).Delete(user)
	return user
}
