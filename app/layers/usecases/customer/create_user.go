package customer

import (
	"golang-blueprint-clean/app/entities"
)

func (useCase *useCase) CreateUser(input *entities.Users) (*entities.Users, error) {
	user, err := useCase.CustomerRepo.CreateUser(input)
	if err != nil {
		return nil, err
	}

	return user, nil
}
