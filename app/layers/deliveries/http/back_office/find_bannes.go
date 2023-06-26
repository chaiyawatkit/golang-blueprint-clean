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
		humanMsg := utils.GetHumanErrorCode(err.Error())
		utils.JSONErrorCodeResponse(c, 400, err, humanMsg)
		return
	}

	//boredom.FuncDebug(c, h.BackOfficeUseCase.FindBanners, FindBannerDataRequestJSON)

	Banners, err, errMsg := h.BackOfficeUseCase.FindBanners(findBannersDataRequest.ToEntity())
	if err != nil {
		//boredom.Error(c, err)
		humanMsg := utils.GetHumanErrorCode(errMsg.Error())
		utils.JSONErrorCodeResponse(c, 400, err, humanMsg)
		return
	}
	bannerList, _ := new(models.FindBannerListResponseJSON).Parse(Banners)
	humanMsg := utils.GetHumanSuccessCode("default")
	utils.JSONSuccessCodeResponse(c, bannerList, humanMsg)
}
