package v1

// Authentication Banner List
// @Summary Banner List
// @Description Banner List
// @ID Banner List
// @Accept json
// @Produce json
// @Tags Banners
// @Param X-Correlation-ID header string true "for request tracking"
// @Param body body models.FindBannerDataRequestJSON true "get by  Segment"
// @Success 200 {object} models.FindBannerDataResponseJSON
// @Failure 400 {object} utils.errorResponse
// @Failure 500 {object} utils.errorResponse
// @Router /v1/banners [get]
func FindBanners() {}
