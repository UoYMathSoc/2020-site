package models

import (
	"gorm.io/gorm"
)

type Model struct {
	gorm.Model
	database *gorm.DB
}
