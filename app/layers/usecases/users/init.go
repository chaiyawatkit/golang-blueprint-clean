package users

import (
	"context"
	"golang-blueprint-clean/app/entities"
	users "golang-blueprint-clean/app/layers/repositories/users"
)

type useCase struct {
	UsersRepo users.Repo
}

func InitUseCase(usersRepo users.Repo) UseCase {
	return &useCase{
		UsersRepo: usersRepo,
	}
}

type UseCase interface {
	Authentication(ctx context.Context, jwtAccessToken string) (*string, error, error)
	SignInBySession(ctx context.Context, input *entities.UsersSignIn) (*entities.Accessibility, error, error)
}
