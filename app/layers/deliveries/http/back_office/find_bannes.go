package back_office

import (
	"github.com/gin-gonic/gin"
	"golang-blueprint-clean/app/layers/deliveries/http/back_office/models"
	"golang-blueprint-clean/app/utils"
)

func (h *handler) FindBanners(c *gin.Context) {
	//boredom.HandlerInfo(c, models.FindBannerDataRequestJSON{})

	findBannersDataRequest, err := new(models.FindBannerDataRequestJSON).Parse(c)
	if err != nil {
		//boredom.Error(c, err)
		utils.JSONErrorResponse(c, err)
		return
	}

	findBannersValid, err := findBannersDataRequest.IsValid()
	if err != nil {
		utils.JSONErrorResponse(c, err)
		return
	}

	//boredom.FuncDebug(c, h.BackOfficeUseCase.FindBanners, FindBannerDataRequestJSON)
	Banners, err := h.BackOfficeUseCase.FindBanners(*findBannersValid.Segment)
	if err != nil {
		//boredom.Error(c, err)
		utils.JSONErrorResponse(c, err)
		return
	}
	user, _ := new(models.FindBannerListResponseJSON).Parse(Banners)

	utils.JSONSuccessResponse(c, user)
}
