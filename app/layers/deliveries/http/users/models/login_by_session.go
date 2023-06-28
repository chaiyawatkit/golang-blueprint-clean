package models

import (
	"github.com/gin-gonic/gin"
	"golang-blueprint-clean/app/entities"
	"golang-blueprint-clean/app/errors"
)

type SignInBySessionRequestJSON struct {
	SessionID string `json:"session_id" binding:"required"`
}

type LoginBySessionResponseJSON struct {
	AccessToken string `json:"accessToken" example:"xxxxxxx.xxxxxxx.xxxxxxx"`
}

// Parse parse data to login
func (model *SignInBySessionRequestJSON) Parse(c *gin.Context) (*SignInBySessionRequestJSON, error) {
	err := c.ShouldBindJSON(model)
	if err != nil {
		return nil, errors.ParameterError{Message: err.Error()}
	}

	return model, nil
}

// IsValid data to loginByEmail
func (model *SignInBySessionRequestJSON) IsValid() (*SignInBySessionRequestJSON, error) {

	if model.SessionID == "" {
		return nil, errors.ParameterError{Message: ""}
	}

	return model, nil
}

// Entity convert LoginByEmailRequestJSON to Users entity
func (model *SignInBySessionRequestJSON) Entity() *entities.UsersSignIn {
	user := entities.UsersSignIn{
		SessionID: model.SessionID,
	}

	return &user
}

// LoginByEmailResponseJSON Parse to JSON
func (model *LoginBySessionResponseJSON) Parse(accessToken *string) *LoginBySessionResponseJSON {
	model.AccessToken = *accessToken
	return model
}

//// LoginByEmailResponseSwagger Login users response for swagger
//type LoginByEmailResponseSwagger struct {
//	Status  string                        `json:"status" example:"success"`
//	Message string                        `json:"message" example:"OK"`
//	Data    CreateUserByEmailResponseJSON `json:"data,omitempty"`
//}
