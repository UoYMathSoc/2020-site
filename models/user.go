package models

import (
	"github.com/UoYMathSoc/2020-site/database"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type UserModel struct {
	Model
	Username       string
	HashedPassword []byte
}

func Get(username string) (user *UserModel) {
	return getUser(username)
}

func getUsers() (users *[]UserModel) {
	db := database.Instance
	db.Find(users)
	return users
}

func newUser(user UserModel) {
	db := database.Instance
	db.Create(user)
}

func getUser(username string) (user *UserModel) {
	db := database.Instance
	db.Where("username = ?", username).Find(&user)
	return user
}

func updateUser(username string, user UserModel) {
	db := database.Instance
	db.Where("username = ?", username).Updates(user)
}

func deleteUser(username string) {
	db := database.Instance
	db.Where("username = ?", username).Delete(UserModel{})
}
