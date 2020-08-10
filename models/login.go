package models

type LoginModel struct {
	Model
}

func NewLoginModel(s int) *LoginModel {
	return &LoginModel{Model{s}}
}

func (m *LoginModel) Post(formParams map[string][]string) (err error) {
	return nil
}
