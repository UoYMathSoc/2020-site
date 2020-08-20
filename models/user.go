package models

import (
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type UserModel struct {
	Model
	Username       string
	HashedPassword []byte
}

func GetUser(username string) (userM *UserModel, err error) {
	db, err := newDatabase()
	if err != nil {
		fmt.Println(err)
		return userM, err
	}
	defer db.Close()

	db.Where("username = ?", username).Find(&userM)
	return userM, err
}
