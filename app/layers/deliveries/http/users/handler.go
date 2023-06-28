package users

import (
	"github.com/gin-gonic/gin"
	"golang-blueprint-clean/app/layers/usecases/users"
	_middlewareHttp "golang-blueprint-clean/app/middlewares/http"
)

type ResponseError struct {
	Message string `json:"message"`
}

type Handler struct {
	UsersUseCase users.UseCase
}

func NewEndpointHttpHandler(ginEngine *gin.Engine, authMiddleware _middlewareHttp.AuthMiddleware, usersUC users.UseCase) {
	handler := &Handler{
		UsersUseCase: usersUC,
	}

	v1Auth := ginEngine.Group("v1").
		Use(authMiddleware.Authentication)
	{
		v1Auth.POST("/users.auth", handler.Authentication)
		v1Auth.POST("/demo.token", handler.DemoToken)

	}

	v1 := ginEngine.Group("v1")
	{

		v1.GET("/users.signin.session", handler.SignInBySession)

	}
}
