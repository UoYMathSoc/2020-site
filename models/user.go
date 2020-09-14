package models

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// UserModel is used when accessing the site's users
type UserModel struct {
	Model
	Username       string
	hashedPassword []byte
}

type Users []UserModel

// NewUserModel returns a new UserModel with access to the database
func NewUserModel(db *gorm.DB) *UserModel {
	return &UserModel{Model: Model{database: db}}
}

// Get populates a user's UserModel. Its uses user's username unless one is provided as a parameter
func (user *UserModel) Get(usernames ...string) error {
	if len(usernames) > 0 {
		user.Username = usernames[0]
	}
	username := user.Username
	if user.read().Username == username {
		return nil
	}
	return errors.New("could not find specified user in the database")
}

//func (user *UserModel) GetUsers() (users *[]UserModel, err error) {
//	user.database.Find(users)
//	for _, thisUser := range *users {
//		thisUser.database = user.database
//	}
//	return users, err
//}

func (user *UserModel) Register(username string, password string) error {
	user.Username = username
	err := user.Get()
	if err == nil {
		return errors.New("user already exists")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return fmt.Errorf("could not generate hash for provided password: %w", err)
	}
	user.hashedPassword = hashedPassword
	user.create()
	err = user.Get()
	if err != nil {
		return errors.New("could not add user to database")
	}
	return nil
}

func (user *UserModel) Validate(password string) error {
	return bcrypt.CompareHashAndPassword(user.hashedPassword, []byte(password))
}

func (user *UserModel) create() *UserModel {
	user.database.Create(&user)
	return user
}

func (user *UserModel) read() *UserModel {
	user.database.Where("username = ?", user.Username).Find(user)
	return user
}

func (user *UserModel) update() *UserModel {
	user.database.Where("username = ?", user.Username).Updates(user)
	return user
}

func (user *UserModel) delete() *UserModel {
	user.database.Where("username = ?", user.Username).Delete(&user)
	return user
}
