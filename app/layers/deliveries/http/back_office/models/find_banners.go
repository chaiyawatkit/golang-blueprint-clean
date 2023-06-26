package models

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"golang-blueprint-clean/app/constants"
	"golang-blueprint-clean/app/entities"
	"golang-blueprint-clean/app/errors"
)

const (
	Pb      = "pb"
	General = "general"
	Wisdom  = "wisdom"
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

	switch *model.Segment {
	case Pb, General, Wisdom:
		return model, nil
	}

	return nil, errors.ParameterError{Message: constants.SegmentType}

}

func (model *FindBannerDataRequestJSON) ToEntity() entities.SegmentTypes {

	segmentTypes := entities.SegmentTypes{}
	if model.Segment != nil {
		segmentTypes.SegmentType = model.Segment
	}
	return segmentTypes
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
