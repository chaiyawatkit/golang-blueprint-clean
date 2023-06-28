package users

import (
	"github.com/gin-gonic/gin"
	"golang-blueprint-clean/app/constants"
	"golang-blueprint-clean/app/env"
	"golang-blueprint-clean/app/layers/deliveries/http/users/models"
	"golang-blueprint-clean/app/utils"
)

func (h *Handler) SignInBySession(c *gin.Context) {

	loginRequestRaw, err := new(models.SignInBySessionRequestJSON).Parse(c)
	if err != nil {

		humanMsg := utils.GetHumanErrorCode(err.Error())
		utils.JSONErrorCodeResponse(c, 400, err, humanMsg)
		return
	}

	errCheckUsername := loginRequestRaw.SessionIdIsValid()
	if errCheckUsername != nil {

		humanMsg := utils.GetHumanErrorCode(errCheckUsername.Error())
		utils.JSONErrorCodeResponse(c, 400, errCheckUsername, humanMsg)
		return
	}

	c.Request.Header.Set(constants.CorrelationHeader, env.XCorrelationID)
	accessToken, err, errMsg := h.UsersUseCase.SignInBySession(c.Request.Context(), loginRequestRaw.ToEntity())
	if err != nil {
		humanMsg := utils.GetHumanErrorCode(errMsg.Error())
		utils.JSONErrorCodeResponse(c, 401, err, humanMsg)
		return
	}

	c.Header(constants.XFinPlusAuth, accessToken.AccessToken)
	humanMsg := utils.GetHumanSuccessCode("default")
	output := new(models.SignInBySessionResponseJSON).Parse(accessToken.AccessToken)
	utils.JSONSuccessCodeResponse(c, output, humanMsg)
}
