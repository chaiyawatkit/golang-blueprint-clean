package v1

// Authentication Create Role
// @Summary Create Role
// @Description Create a role
// @ID post-roles
// @Accept json
// @Produce json
// @Tags roles
// @Param X-Correlation-ID header string true "for request tracking"
// @Param body body models.CreateRoleRequestJSON true "All params related to role"
// @Success 200 {object} models.CreateRoleResponseSwagger
// @Failure 400 {object} utils.errorResponse
// @Failure 500 {object} utils.errorResponse
// @Router /v1/roles [post]
func CreateRole() {}
