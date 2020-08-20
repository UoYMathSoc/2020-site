package models

import (
	"github.com/jinzhu/gorm"
)

type LoginModel struct {
	gorm.Model
}

func (m *LoginModel) Post(formParams map[string][]string) (err error) {
	return nil
}
