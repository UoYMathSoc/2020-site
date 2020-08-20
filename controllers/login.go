package controllers

//type LoginController struct {
//	Controller
//}
//
//func NewLoginController(s *database.Session, c *structs.Config) *LoginController {
//	return &LoginController{Controller{session: s,config:  c}}
//}
//
//func (loginC *LoginController) Post(w http.ResponseWriter, r *http.Request) {
//	r.ParseForm()
//	formParams := r.Form
//
//	loginM := models.NewLoginModel(loginC.session)
//	err := loginM.Post(formParams)
//
//	userTable, err := database.NewUserTable()
//	if err != nil {
//		w.WriteHeader(http.StatusInternalServerError)
//		return
//	}
//
//	retUser := new(database.User)
//	err = userTable.Select(username, retUser)
//	if err != nil {
//		if err == sql.ErrNoRows {
//			w.WriteHeader(http.StatusUnauthorized)
//			return
//		}
//		w.WriteHeader(http.StatusInternalServerError)
//		return
//	}
//
//	if err = bcrypt.CompareHashAndPassword(retUser.HashedPassword, []byte(password)); err!= nil {
//		w.WriteHeader(http.StatusUnauthorized)
//	}
//}
