package back_office

import (
	"github.com/gin-gonic/gin"
	backOffice "golang-blueprint-clean/app/layers/usecases/back_office"
)

type handler struct {
	BackOfficeUseCase backOffice.UseCase
}

// NewEndpointHandler routing
func NewEndpointHttpHandler(ginEngine *gin.Engine, backofficeUseCase backOffice.UseCase) {
	handler := &handler{
		BackOfficeUseCase: backofficeUseCase,
	}

	v1 := ginEngine.Group("v1")
	{

		v1.GET("/banners", handler.FindBanners)

	}
}
