package users

import (
	"golang-blueprint-clean/app/entities"
	"golang-blueprint-clean/app/layers/repositories/users"
)

type useCase struct {
	UsersRepo users.Repo
}

func InitUseCase(users users.Repo) UseCase {
	return &useCase{
		UsersRepo: users,
	}
}

// InitUseCase init auth use case
type UseCase interface {
	LoginBySession(input *entities.UsersSignIn) (*string, error)
	FindUserByAccessToken(accessToken string) (*entities.Users, error)
}
