package forms

import (
	"errors"
	"fmt"
	"strings"

	"github.com/UoYMathSoc/2020-site/controllers"
)

func NewUser(ac *controllers.AdminController) (string, controllers.Response) {
	return "New User", NewUserResponse{}
}

type NewUserResponse struct {
	Name  string
	Email string `form:"type=email"`

	Password     string `form:"type=password"`
	Confirmation string `form:"display=Password Confirmation;type=password"`

	Bio string `form:"type=textarea;placeholder=Tell us a bit about yourself"`
}

var ErrPasswordMismatch = errors.New("passwords do not match")
var ErrNotStudentEmail = errors.New("used unmanaged email")

func (nur NewUserResponse) Do(ac *controllers.AdminController) error {
	emailSuffix := "york.ac.uk"

	if nur.Password != nur.Confirmation {
		return ErrPasswordMismatch
	}
	if !strings.HasSuffix(nur.Email, emailSuffix) {
		return ErrNotStudentEmail
	}
	username := strings.TrimSuffix(nur.Email, emailSuffix)
	_, err := ac.Users.Create(username, nur.Password)
	if err != nil {
		return fmt.Errorf("could not create user: %w", err)
	}
	return nil
}

var _ controllers.Form = NewUser
var _ controllers.Response = NewUserResponse{}
