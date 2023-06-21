package back_office

import (
	"golang-blueprint-clean/app/entities"
	backOffice "golang-blueprint-clean/app/layers/repositories/back_office"
)

type useCase struct {
	BackOfficeRepo backOffice.Repo
}

func InitUseCase(backOffice backOffice.Repo) UseCase {
	return &useCase{
		BackOfficeRepo: backOffice,
	}
}

// InitUseCase init auth use case
type UseCase interface {
	FindBanners(segment string) ([]entities.Banners, error)
}
