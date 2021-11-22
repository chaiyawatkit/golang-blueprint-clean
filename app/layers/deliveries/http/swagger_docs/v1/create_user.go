package v1

// Authentication Create User
// @Summary Create User
// @Description Create a user
// @ID post-users
// @Accept json
// @Produce json
// @Tags users
// @Param X-Correlation-ID header string true "for request tracking"
// @Param body body models.CreateUserRequestJSON true "All params related to user"
// @Success 200 {object} models.CreateUserResponse
// @Failure 400 {object} utils.errorResponse
// @Failure 500 {object} utils.errorResponse
// @Router /v1/users [post]
func CreateUser() {}
