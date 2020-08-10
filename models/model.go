package models

type ModelInterface interface {
	Get() (data *interface{}, err error)
}

type Model struct {
	session int
}
