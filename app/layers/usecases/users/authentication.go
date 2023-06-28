package users

import (
	"context"
)

func (u *useCase) Authentication(ctx context.Context, jwtAccessToken string) (*string, error, error) {
	response, err, errCode := u.UsersRepo.Authentication(ctx, jwtAccessToken)

	if err != nil {
		return nil, err, errCode
	}

	return response, nil, nil
}
