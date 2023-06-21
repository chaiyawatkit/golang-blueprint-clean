package models

import (
	"github.com/jinzhu/copier"
	"golang-blueprint-clean/app/entities"
)

type Banners struct {
	ID        uint
	Title     string
	Thumbnail string
	Status    string
	Type      string
	Redirect  string
	CreatedAt int
	EndAt     int
	Segment   string
}
type BannersArray []Banners

func (model Banners) Parse() (*entities.Banners, error) {
	entity := new(entities.Banners)
	err := copier.Copy(entity, model)
	return entity, err
}

func (model BannersArray) BannerEntity() []entities.Banners {
	var results []entities.Banners
	for _, v := range model {
		entity, _ := v.Parse()
		results = append(results, *entity)
	}
	return results
}
