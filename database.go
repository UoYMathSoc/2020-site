package main

import (
	"fmt"
	"github.com/UoYMathSoc/2020-site/models"
	"github.com/jinzhu/gorm"
)

// TODO: Find a better way of passing these in
var ActiveModels = [...]interface{}{
	models.UserModel{},
	models.LoginModel{},
}

func MigrateModels(models ...interface{}) {
	db, err := gorm.Open("postgres", "dbname=mydb sslmode=disable")
	if err != nil {
		fmt.Println(err)
		panic("failed to connect to database")
	}
	defer db.Close()

	for _, model := range models {
		db.AutoMigrate(model)
	}
}
