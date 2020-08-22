package models

import (
	"github.com/UoYMathSoc/2020-site/structs"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Model struct {
	gorm.Model
	config *structs.Config
}
