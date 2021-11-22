package customer

import (
	"github.com/gin-gonic/gin"
	//"github.com/golang-blueprint-clean/boredom"
	"golang-blueprint-clean/app/layers/deliveries/http/customer/models"
	"golang-blueprint-clean/app/utils"
)

func (h *handler) CreateUser(c *gin.Context) {
	//boredom.HandlerInfo(c, nil)

	createUserRequest, err := new(models.CreateUserRequestJSON).Parse(c)
	if err != nil {
		//boredom.Error(c, err)
		utils.JSONErrorResponse(c, err)
		return
	}

	//boredom.FuncDebug(c, h.CustomerUseCase.CreateUser, createUserRequest)
	userOutput, err := h.CustomerUseCase.CreateUser(createUserRequest.Entity())
	if err != nil {
		//boredom.Error(c, err)
		utils.JSONErrorResponse(c, err)
		return
	}

	//ginney.JSONSuccessResponse(c, userOutput)
	utils.JSONSuccessResponse(c, userOutput)
}
