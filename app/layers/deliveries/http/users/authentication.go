package users

import (
	"github.com/gin-gonic/gin"
	"golang-blueprint-clean/app/constants"
	"golang-blueprint-clean/app/errors"
	"golang-blueprint-clean/app/utils"
)

func (h *Handler) Authentication(c *gin.Context) {
	accessToken := c.GetHeader(constants.XFinPlusAuth)
	if accessToken == "" {

		utils.JSONErrorResponse(c, errors.ParameterError{Message: "ArgumentIsMissing"})
		return
	}

	_, err := h.UsersUseCase.FindUserByAccessToken(accessToken)

	if err != nil {

		utils.JSONErrorResponse(c, err)
		return
	}

	c.Header(constants.XFinPlusAuth, accessToken)
	utils.JSONSuccessResponse(c, map[string]interface{}{"accessToken": accessToken})

}
