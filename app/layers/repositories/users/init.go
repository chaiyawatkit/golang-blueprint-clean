package users

import (
	"context"
	"golang-blueprint-clean/app/entities"
)

type repo struct {
}

func InitRepo() Repo {
	return &repo{}
}

type Repo interface {
	Authentication(ctx context.Context, jwtAccessToken string) (*string, error, error)
	LoginBySession(ctx context.Context, input entities.UsersSignIn) (*string, error, error)
}
