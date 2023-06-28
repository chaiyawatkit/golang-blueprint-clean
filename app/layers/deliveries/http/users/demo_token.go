package users

import (
	"github.com/gin-gonic/gin"
	"golang-blueprint-clean/app/constants"
	"golang-blueprint-clean/app/utils"
)

func (h *Handler) DemoToken(c *gin.Context) {

	jwtRawData, _ := c.Get(constants.JWTDataKey)

	humanMsg := utils.GetHumanSuccessCode("default")
	utils.JSONSuccessCodeResponse(c, jwtRawData, humanMsg)

}
