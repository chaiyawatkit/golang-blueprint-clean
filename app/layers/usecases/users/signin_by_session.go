package users

import (
	"context"
	"golang-blueprint-clean/app/entities"
)

func (u *useCase) SignInBySession(ctx context.Context, input *entities.UsersSignIn) (*entities.Accessibility, error, error) {
	sessionLogin := entities.UsersSignIn{
		SessionID: input.SessionID,
	}

	token, err, repoErr := u.UsersRepo.LoginBySession(ctx, sessionLogin)
	if err != nil {
		return nil, err, repoErr
	}

	access := &entities.Accessibility{
		AccessToken: *token,
	}
	return access, nil, nil
}
