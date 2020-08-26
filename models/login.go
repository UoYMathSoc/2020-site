package models

type LoginModel struct {
	Model
}

func NewLoginModel() *LoginModel {
	return new(LoginModel)
}

func (m *LoginModel) Post(formParams map[string][]string) (err error) {
	username := "jgd511"
	password := "password"
	NewUser(username, password)
	return
}

//func (m *LoginModel) Post(formParams map[string][]string) (err error) {
//	username := formParams["username"][0]
//	password := formParams["password"][0]
//
//	user := read(username)
//
//	return bcrypt.CompareHashAndPassword(user.HashedPassword, []byte(password))
//}
