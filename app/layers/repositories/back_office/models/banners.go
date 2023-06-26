package models

import (
	"github.com/jinzhu/copier"
	"golang-blueprint-clean/app/entities"
)

type Banners struct {
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

type ResponseData struct {
	Status  string    `json:"status"`
	Message string    `json:"message"`
	Data    []Banners `json:"data"`
}

type BannersArray []Banners

func (model Banners) Parse() (*entities.Banners, error) {
	entity := new(entities.Banners)
	err := copier.Copy(entity, model)
	return entity, err
}

func (model ResponseData) BannerEntity() []entities.Banners {
	var results []entities.Banners
	for _, v := range model.Data {
		entity, _ := v.Parse()
		results = append(results, *entity)
	}
	return results
}
