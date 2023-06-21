package back_office

import (
	"golang-blueprint-clean/app/constants"
	"golang-blueprint-clean/app/entities"
	"golang-blueprint-clean/app/errors"
)

func (useCase *useCase) FindBanners(segment string) ([]entities.Banners, error) {
	banners, err := useCase.BackOfficeRepo.FindBanners(segment)
	if err != nil {
		return nil, errors.InternalError{Message: constants.FailToGetDataFromDB}
	}

	return banners, nil
}
