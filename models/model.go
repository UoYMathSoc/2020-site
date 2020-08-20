package models

import (
	"github.com/jinzhu/gorm"
)

type Model struct {
	gorm.Model
}

func newDatabase() (*gorm.DB, error) {
	return gorm.Open("postgres", "dbname=mydb sslmode=disable")
}
