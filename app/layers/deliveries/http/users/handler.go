package users

import (
	"github.com/gin-gonic/gin"
	"golang-blueprint-clean/app/layers/usecases/users"
)

type Handler struct {
	UsersUseCase users.UseCase
}

func NewEndpointHttpHandler(ginEngine *gin.Engine, usersUC users.UseCase) {
	handler := &Handler{
		UsersUseCase: usersUC,
	}

	v1 := ginEngine.Group("v1")
	{

		v1.GET("/users.session.login", handler.LoginBySession)
		v1.POST("/users.auth", handler.Authentication)

	}
}
