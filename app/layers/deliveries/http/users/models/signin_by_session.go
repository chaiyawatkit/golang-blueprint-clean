package models

import (
	"github.com/gin-gonic/gin"
	"golang-blueprint-clean/app/entities"
	"golang-blueprint-clean/app/errors"
	"strings"
)

type SignInBySessionRequestJSON struct {
	SessionID string `json:"session_id" binding:"required"`
}

type SignInBySessionResponseJSON struct {
	AccessToken string `json:"accessToken"`
	Method      string `json:"method"`
}

func (model *SignInBySessionRequestJSON) Parse(c *gin.Context) (*SignInBySessionRequestJSON, error) {
	err := c.ShouldBindJSON(model)
	if err != nil {
		return nil, errors.ParameterError{Message: err.Error()}
	}

	return model, nil
}

func (model *SignInBySessionRequestJSON) SessionIdIsValid() error {
	if model.SessionID == "" {
		return errors.ParameterError{Message: ""}
	}
	model.SessionID = strings.ToLower(model.SessionID)
	return nil
}

func (model *SignInBySessionResponseJSON) Parse(accessToken string) *SignInBySessionResponseJSON {
	model.AccessToken = accessToken
	model.Method = "signin"
	return model
}

func (model *SignInBySessionRequestJSON) ToEntity() *entities.UsersSignIn {
	usersEntity := entities.UsersSignIn{
		SessionID: model.SessionID,
	}

	return &usersEntity
}
