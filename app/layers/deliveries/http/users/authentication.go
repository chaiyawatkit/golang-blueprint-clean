package users

import (
	"github.com/gin-gonic/gin"
	"golang-blueprint-clean/app/constants"
	"golang-blueprint-clean/app/errors"
	"golang-blueprint-clean/app/utils"
)

func (h *Handler) Authentication(c *gin.Context) {
	jwtAccessToken := c.GetHeader(constants.XFinPlusAuth)
	if jwtAccessToken == "" {
		//boredom.Error(c, errors.ParameterError{Message: constants.MissingFinPlusAuth})

		humanMsg := utils.GetHumanErrorCode(constants.MissingFinPlusAuth)
		utils.JSONErrorCodeResponse(c, 400, errors.ParameterError{Message: constants.MissingFinPlusAuth}, humanMsg)
		return
	}

	c.Header(constants.XFinPlusAuth, jwtAccessToken)
	humanMsg := utils.GetHumanSuccessCode("default")
	utils.JSONSuccessCodeResponse(c, map[string]interface{}{"accessToken": jwtAccessToken}, humanMsg)

}
