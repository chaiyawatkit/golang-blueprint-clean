package back_office

import (
	"golang-blueprint-clean/app/entities"
)

type repo struct {
}

// InitRepo CRUD
func InitRepo() Repo {
	return &repo{}
}

type Repo interface {
	FindBanners(filter *entities.SegmentTypes) ([]entities.Banners, error, error)
}
