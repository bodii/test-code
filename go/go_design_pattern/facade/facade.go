package facade

// User user struct
type User struct {
	Name string
}

// IUser user interface
type IUser interface {
	Login(phone int, code int) (*User, error)
	Register(phone int, code int) (*User, error)
}

// IUserFacade user facade interface
type IUserFacade interface {
	LoginOrRegister(phone int, code int) (*User, error)
}

// UserService user service struct
type UserService struct{}

// Login login function
func (u UserService) Login(phone int, code int) (*User, error) {
	return &User{Name: "test login"}, nil
}

// Register register function
func (u UserService) Register(phone int, code int) (*User, error) {
	return &User{Name: "test register"}, nil
}

// LoginOrRegister login or register function
func (u UserService) LoginOrRegister(phone int, code int) (*User, error) {
	user, err := u.Login(phone, code)
	if err != nil {
		return nil, err
	}

	if user != nil {
		return user, nil
	}

	return u.Register(phone, code)
}
