package v1

// USER Find One USERS DATA
// @Summary Find One USERS DATA
// @Description Find One a USERS DATA
// @ID post-users-find-one-data
// @Accept json
// @Produce json
// @Tags users
// @Param X-Correlation-ID header string true "for request tracking"
// @Param id query string false "filter from a id"
// @Param email query string false "filter from a email"
// @Success 200 {object} models.FindOneUserDataResponseJSON
// @Failure 400 {object} utils.errorResponse
// @Failure 500 {object} utils.errorResponse
// @Router /v1/users.data [get]
func FindOneUserData() {}
