package models

import (
	"github.com/jinzhu/gorm"
)

type Model struct {
	gorm.Model
}

type table interface {
	create()
	read()
	update()
	delete()
}
