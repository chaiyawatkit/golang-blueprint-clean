package customer

import "golang-blueprint-clean/app/entities"

func (useCase *useCase) CreateRole(input *entities.Roles) (*entities.Roles, error) {
	role, err := useCase.CustomerRepo.CreateRoles(input)
	if err != nil {
		return nil, err
	}

	return role, nil
}
