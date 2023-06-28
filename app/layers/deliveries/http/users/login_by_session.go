package users

import (
	"github.com/gin-gonic/gin"
	"golang-blueprint-clean/app/constants"
	"golang-blueprint-clean/app/layers/deliveries/http/users/models"
	"golang-blueprint-clean/app/utils"
)

func (h *Handler) LoginBySession(c *gin.Context) {

	loginByEmailRequestRaw, err := new(models.SignInBySessionRequestJSON).Parse(c)

	if err != nil {

		utils.JSONErrorResponse(c, err)
		return
	}

	loginByEmailRequest, errCheck := loginByEmailRequestRaw.IsValid()
	if errCheck != nil {

		utils.JSONErrorResponse(c, errCheck)
		return
	}

	token, err := h.UsersUseCase.LoginBySession(loginByEmailRequest.Entity())
	if err != nil {

		utils.JSONErrorResponse(c, err)
		return
	}

	user := new(models.LoginBySessionResponseJSON).Parse(token)
	c.Header(constants.XFinPlusAuth, user.AccessToken)
	utils.JSONSuccessResponse(c, user)
}
