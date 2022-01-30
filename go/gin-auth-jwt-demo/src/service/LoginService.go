package service

type LoginService interface {
	LoginUser(email, password string) bool
}

type loginInformation struct {
	email    string
	password string
}

func StaticLoginService() LoginService {
	return &loginInformation{
		email:    "aaa@test.com",
		password: "abcdefghi",
	}
}

func (info *loginInformation) LoginUser(email, password string) bool {
	return info.email == email && info.password == password
}
