package http

import (
	"github.com/gin-gonic/gin"
)

type AuthMiddleware interface {
	Authentication(c *gin.Context)
}

type authMiddleware struct {
	UsersUseCase interface{}
}

func (middleware *authMiddleware) Authentication(c *gin.Context) {

}

func InitAuthMiddleware(usersUseCase interface{}) AuthMiddleware {
	return &authMiddleware{
		UsersUseCase: usersUseCase,
	}
}
