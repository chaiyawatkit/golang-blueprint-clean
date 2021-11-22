package customer

import (
	"golang-blueprint-clean/app/entities"
	"golang-blueprint-clean/app/layers/repositories/customer"
)

type useCase struct {
	CustomerRepo customer.Repo
}

func InitUseCase(customerRepository customer.Repo) UseCase {
	return &useCase{
		CustomerRepo: customerRepository,
	}
}

// InitUseCase init auth use case
type UseCase interface {
	CreateUser(input *entities.Users) (*entities.Users, error)
	CreateRole(input *entities.Roles) (*entities.Roles, error)

	FindOneUserData(input *entities.UsersFilter) (*entities.Users, error)
}
