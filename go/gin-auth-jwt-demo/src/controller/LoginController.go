package controller

import (
	"gin-auth-jwt-demo/src/locred"
	"gin-auth-jwt-demo/src/service"

	"github.com/gin-gonic/gin"
)

type LoginController interface {
	Login(ctx *gin.Context) string
}

type loginController struct {
	loginService service.LoginService
	jwtService   service.JWTService
}

func LoginHandler(loginService service.LoginService, jwtService service.JWTService) LoginController {
	return &loginController{
		loginService: loginService,
		jwtService:   jwtService,
	}
}

func (controller *loginController) Login(ctx *gin.Context) string {
	var credential locred.LoginCredentials
	err := ctx.ShouldBind(&credential)
	if err != nil {
		return "no data found"
	}

	isUserAuthenticated := controller.loginService.LoginUser(credential.Email, credential.Password)
	if isUserAuthenticated {
		return controller.jwtService.GeneraleToken(credential.Email, isUserAuthenticated)
	}

	return ""
}
