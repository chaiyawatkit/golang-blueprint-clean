package back_office

import (
	"golang-blueprint-clean/app/entities"
)

func (useCase *useCase) FindBanners(input entities.SegmentTypes) (*[]entities.Banners, error, error) {

	banners, err, errMsg := useCase.BackOfficeRepo.FindBanners(&input)
	if err != nil {
		return nil, err, errMsg
	}

	return &banners, nil, nil
}
