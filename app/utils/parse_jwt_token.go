package utils

import (
	"crypto/rsa"
	"github.com/dgrijalva/jwt-go"
	"golang-blueprint-clean/app/constants"
	"golang-blueprint-clean/app/errors"
)

func ParseJwtToken(accessToken string, publicKey *rsa.PublicKey) (*jwt.Token, error) {
	jwtToken, _ := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodRSA)
		if !ok {
			return nil, errors.InternalError{Message: constants.FailUnExpectedSigningMethod}
		}
		return publicKey, nil
	})
	return jwtToken, nil
}
