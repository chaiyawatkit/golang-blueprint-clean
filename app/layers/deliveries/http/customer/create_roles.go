package customer

import (
	"github.com/chaiyawatkit/ginney"
	"github.com/gin-gonic/gin"
	"golang-blueprint-clean/app/layers/deliveries/http/customer/models"
	"golang-blueprint-clean/app/utils"
)

func (h *handler) CreateRoles(c *gin.Context) {
	//boredom.HandlerInfo(c, nil)

	createRoleRequest, err := new(models.CreateRoleRequestJSON).Parse(c)
	if err != nil {
		//boredom.Error(c, err)
		humanMsg := utils.GetHumanErrorCode(err.Error())
		ginney.JSONErrorCodeResponse(c, 401, err, humanMsg)
		return
	}

	//entityRole := createRoleRequest.Entity()

	//boredom.FuncDebug(c, h.CustomerUseCase.CreateRole, createRoleRequest)
	roleOutput, err := h.CustomerUseCase.CreateRole(createRoleRequest.Entity())
	if err != nil {
		//boredom.Error(c, err)
		utils.JSONErrorResponse(c, err)
		return
	}

	createRole, _ := new(models.CreateRoleResponseJSON).Parse(roleOutput)

	//ginney.JSONSuccessResponse(c, createRole)
	utils.JSONSuccessResponse(c, createRole)
}
