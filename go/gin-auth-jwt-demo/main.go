package main

import (
	"gin-auth-jwt-demo/src/controller"
	"gin-auth-jwt-demo/src/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	var loginService service.LoginService = service.StaticLoginService()
	var jwtService service.JWTService = service.JWTAuthService()
	var loginController controller.LoginController = controller.LoginHandler(loginService, jwtService)

	serv := gin.New()

	serv.POST("/login", func(ctx *gin.Context) {
		token := loginController.Login(ctx)
		if token != "" {
			ctx.JSON(http.StatusOK, gin.H{
				"token": token,
			})
		} else {
			ctx.JSON(http.StatusUnauthorized, nil)
		}
	})

	serv.Run(":8000")
}
