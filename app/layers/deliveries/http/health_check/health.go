package healthcheck

import (
	"github.com/gin-gonic/gin"
	"golang-blueprint-clean/app/utils"
)

func (handler *handler) Health(c *gin.Context) {
	utils.JSONSuccessResponse(c, nil)
}
