package models

import (
	"github.com/jinzhu/gorm"
)

type Model struct {
	gorm.Model
	database *gorm.DB
}
