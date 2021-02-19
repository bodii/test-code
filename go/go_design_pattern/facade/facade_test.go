package facade

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserService_Login(t *testing.T) {
	service := UserService{}
	user, err := service.Login(13011112222, 123456)
	assert.NoError(t, err)
	assert.Equal(t, &User{Name: "test login"}, user)
}

func TestUserService_Register(t *testing.T) {
	service := UserService{}
	user, err := service.Register(15611112222, 123456)
	assert.NoError(t, err)
	assert.Equal(t, &User{Name: "test register"}, user)
}

func TestUserService_LoginOrRegister(t *testing.T) {
	service := UserService{}
	user, err := service.LoginOrRegister(18011112222, 123456)
	assert.NoError(t, err)
	assert.Equal(t, &User{Name: "test login"}, user)
}
