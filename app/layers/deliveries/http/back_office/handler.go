package back_office

import (
	"github.com/gin-gonic/gin"
	backOffice "golang-blueprint-clean/app/layers/usecases/back_office"
	_middlewareHttp "golang-blueprint-clean/app/middlewares/http"
)

type handler struct {
	BackOfficeUseCase backOffice.UseCase
}

func NewEndpointHttpHandler(ginEngine *gin.Engine, authMiddleware _middlewareHttp.AuthMiddleware, backofficeUseCase backOffice.UseCase) {
	handler := &handler{
		BackOfficeUseCase: backofficeUseCase,
	}

	v1Auth := ginEngine.Group("v1").
		Use(authMiddleware.Authentication)
	{
		v1Auth.GET("/banners", handler.FindBanners)

	}

	//v1 := ginEngine.Group("v1")
	//{
	//
	//}
}
