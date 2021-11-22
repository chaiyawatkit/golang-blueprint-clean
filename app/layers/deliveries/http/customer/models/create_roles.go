package models

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"golang-blueprint-clean/app/entities"
	"golang-blueprint-clean/app/errors"
	"time"
)

// CreateRoleRequestJSON receive request data to create new record
type CreateRoleRequestJSON struct {
	Code        string `json:"code" binding:"required" example:"BOOM"`
	DisplayName string `json:"displayName" binding:"required" example:"BoomDisplay"`
}

type CreateRoleResponseJSON struct {
	ID          uint      `json:"id" example:"1"`
	Code        string    `json:"code" example:"BOOM"`
	DisplayName string    `json:"displayName" example:"BoomDisplay"`
	CreatedAt   time.Time `json:"createAt" example:"2019-02-14T02:35:31.2296459Z"`
}

// Parse parse data to create new record
func (model *CreateRoleRequestJSON) Parse(c *gin.Context) (*CreateRoleRequestJSON, error) {
	err := c.ShouldBindJSON(model)
	if err != nil {
		return nil, errors.ParameterError{Message: err.Error()}
	}

	return model, nil
}

// Entity convert CreateRoleRequestJSON to Roles entity
func (model *CreateRoleRequestJSON) Entity() *entities.Roles {
	roleEntity := entities.Roles{
		Code:        model.Code,
		DisplayName: model.DisplayName,
	}
	return &roleEntity
}

// CreateRoleResponseJSON Parse to JSON
func (model *CreateRoleResponseJSON) Parse(data interface{}) (*CreateRoleResponseJSON, error) {
	err := copier.Copy(model, data)
	return model, err
}

// CreateRoleResponseSwagger create role response for swagger
type CreateRoleResponseSwagger struct {
	Status  string                 `json:"status" example:"success"`
	Message string                 `json:"message" example:"OK"`
	Data    CreateRoleResponseJSON `json:"data,omitempty"`
}
