package proxy

import (
	"log"
	"time"
)

// IUser user interface class
type IUser interface {
	Login(username, password string) error
}

// User user struct
type User struct {
}

// Login user login function
func (u *User) Login(username, password string) error {
	return nil
}

// UserProxy user proxy struct
type UserProxy struct {
	user *User
}

// NewUserProxy create an new user proxy
func NewUserProxy(user *User) *UserProxy {
	return &UserProxy{user: user}
}

// Login implement user proxy login func
func (p *UserProxy) Login(username, password string) error {
	start := time.Now()

	if err := p.user.Login(username, password); err != nil {
		return err
	}

	log.Printf("user login cost time: %s", time.Now().Sub(start))

	return nil
}
