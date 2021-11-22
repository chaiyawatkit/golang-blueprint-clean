package customer

import (
	"golang-blueprint-clean/app/constants"
	"golang-blueprint-clean/app/entities"
	"golang-blueprint-clean/app/errors"
)

func (useCase *useCase) FindOneUserData(input *entities.UsersFilter) (*entities.Users, error) {
	// Find User
	user, err := useCase.CustomerRepo.FindOneUser(input)
	if err != nil {
		return nil, errors.InternalError{Message: constants.FailToGetDataFromDB}
	}

	// Find Role By RoleID
	filterRole := entities.RolesFilter{
		ID: &user.RoleID,
	}
	role, err := useCase.CustomerRepo.FindOneRole(&filterRole)
	if err != nil {
		return nil, errors.InternalError{Message: constants.FailToGetDataFromDB}
	}

	user.Roles = role

	return user, nil
}
