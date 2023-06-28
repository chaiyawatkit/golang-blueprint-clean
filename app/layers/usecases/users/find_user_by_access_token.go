package users

import (
	"github.com/dgrijalva/jwt-go"
	"golang-blueprint-clean/app/constants"
	"golang-blueprint-clean/app/entities"
	"golang-blueprint-clean/app/errors"
	"golang-blueprint-clean/app/utils"
)

func (u *useCase) FindUserByAccessToken(accessToken string) (*entities.Users, error) {
	if accessToken == "" {
		return nil, errors.ParameterError{Message: "AccessTokenMissing"}
	}

	publicKey, err := utils.GetPublicKey()
	if err != nil {
		return nil, errors.InternalError{Message: constants.FailToGrabPublicKey}
	}

	jwtToken, _ := utils.ParseJwtToken(accessToken, publicKey)

	if jwtToken == nil {
		return nil, errors.Unauthorized{Message: constants.FailJwtTokenCannotBeClaimed}
	}

	_, ok := jwtToken.Claims.(jwt.MapClaims)
	if !ok || !jwtToken.Valid {
		return nil, errors.Unauthorized{Message: constants.FailJwtTokenCannotBeClaimed}
	}

	user, err := u.UsersRepo.FindUserByAccessToken(accessToken)
	if err != nil {
		return nil, errors.InternalError{Message: constants.FailToGetDataFromDB}
	}

	if user.Uuid == 0 {
		return nil, errors.Unauthorized{Message: constants.NotFoundUserWithGivenAccessToken}
	}

	//เทียบว่าตรงกันไหม
	//userUuid := user.Uuid
	//jwtUserUuid := claims["uuid"]
	//
	//if jwtUserUuid != userUuid {
	//	return nil, errors.Unauthorized{Message: constants.GivenUserUuidIsUnmatchedWithJwtUserUid}
	//}

	return user, nil
}
