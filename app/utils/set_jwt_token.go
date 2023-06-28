package utils

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"golang-blueprint-clean/app/entities"
	"golang-blueprint-clean/app/env"
	"golang-blueprint-clean/app/errors"
	"strconv"
	"time"
)

func GetJwtToken(user *entities.Users) (*string, error) {
	stringToEncrypt := fmt.Sprint(user.Uuid)
	encrypted, err := EncryptAES(stringToEncrypt, env.EncryptKey)
	if err != nil {
		return nil, errors.InternalError{Message: err.Error()}
	}
	// Default 0 Jwt Token Life Is An Never Expired
	jwtTokenLife, err := strconv.Atoi("30")
	if err != nil {
		jwtTokenLife = 0
	}
	jwtMapClaims := jwt.MapClaims{}
	if jwtTokenLife > 0 {
		jwtMapClaims = jwt.MapClaims{
			"iat":     time.Now().Unix(),
			"exp":     time.Now().Add(time.Minute * time.Duration(jwtTokenLife)).Unix(),
			"uuid":    user.Uuid,
			"guid":    fmt.Sprintf("%x", encrypted),
			"status":  user.UserStatus,
			"segment": user.Segment,
		}
	} else {
		jwtMapClaims = jwt.MapClaims{
			"iat":     time.Now().Unix(),
			"guid":    fmt.Sprintf("%x", encrypted),
			"uuid":    user.Uuid,
			"status":  user.UserStatus,
			"segment": user.Segment,
		}
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodRS512, jwtMapClaims)

	privateKey, err := GetPrivateKey()
	if err != nil {
		return nil, errors.InternalError{Message: err.Error()}
	}
	jwtTokenString, err := jwtToken.SignedString(privateKey)
	if err != nil {
		return nil, errors.InternalError{Message: err.Error()}
	}

	return &jwtTokenString, nil
}
