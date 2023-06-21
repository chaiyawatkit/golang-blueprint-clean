package models

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"golang-blueprint-clean/app/constants"
	"golang-blueprint-clean/app/errors"
)

type FindBannerDataRequestJSON struct {
	Segment *string `form:"segment"`
}

type FindBannerDataResponseJSON struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	Thumbnail string `json:"thumbnail"`
	Status    string `json:"status"`
	Type      string `json:"type"`
	Redirect  string `json:"redirect"`
	CreatedAt int    `json:"created_at"`
	EndAt     int    `json:"end_at"`
	Segment   string `json:"segment"`
}

type FindBannerListResponseJSON []FindBannerDataResponseJSON

func (model *FindBannerDataRequestJSON) Parse(c *gin.Context) (*FindBannerDataRequestJSON, error) {
	emptyRequest := FindBannerDataRequestJSON{}
	c.ShouldBindQuery(model)
	if *model == emptyRequest {
		return nil, errors.ParameterError{Message: constants.EmptyParameter}
	}

	return model, nil
}

func (model *FindBannerDataRequestJSON) IsValid() (*FindBannerDataRequestJSON, error) {

	if model.Segment == nil || *model.Segment == "" {
		return nil, errors.ParameterError{Message: constants.InValidSegment}
	}

	return model, nil
}

func (model *FindBannerListResponseJSON) Parse(data interface{}) (*FindBannerListResponseJSON, error) {

	err := copier.Copy(model, data)
	return model, err
}

type FindBannerDataResponseSwagger struct {
	Status  string                     `json:"status" example:"success"`
	Message string                     `json:"message" example:"OK"`
	Data    FindBannerListResponseJSON `json:"data,omitempty"`
}
