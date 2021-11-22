package models

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"golang-blueprint-clean/app/constants"
	"golang-blueprint-clean/app/entities"
	"golang-blueprint-clean/app/errors"
)

// FindOneUserDataRequestJSON receive request data for get one record
type FindOneUserDataRequestJSON struct {
	ID    *uint   `form:"id" example:"1"`
	Email *string `form:"email" example:"tech@mail.co"`
}

type FindOneUserDataResponseJSON struct {
	Email       string       `json:"email" validate:"email" example:"tech@mail.co.th"`
	FirstName   string       `json:"firstName" example:"tech@mail.co"`
	LastName    string       `json:"lastName" example:"tech@mail.co"`
	Age         *string      `json:"age" example:"20"`
	BirthDate   *string      `json:"birthDate" example:"1999/09/09"`
	Address     *string      `json:"address" example:"92/7 Parkred"`
	PhoneNumber *string      `json:"phoneNumber" example:"+66999999999"`
	Provider    string       `json:"provider" example:"OWN"`
	StatusID    uint         `json:"statusId" example:"1"`
	Role        roleResponse `json:"role"`
	RoleTypeID  uint         `json:"roleTypeId" example:"1"`
}

type roleResponse struct {
	ID          uint   `json:"id" example:"1"`
	Code        string `json:"code" example:"ADMIN"`
	DisplayName string `json:"displayName" example:"Admin"`
}

// Parse parse data to create new record
func (model *FindOneUserDataRequestJSON) Parse(c *gin.Context) (*FindOneUserDataRequestJSON, error) {
	emptyRequest := FindOneUserDataRequestJSON{}
	c.ShouldBindQuery(model)
	if *model == emptyRequest {
		return nil, errors.ParameterError{Message: constants.EmptyParameter}
	}
	return model, nil
}

// Entity convert FindOneUserDataRequestJSON to Roles entity
func (model *FindOneUserDataRequestJSON) Entity() *entities.UsersFilter {
	filterEntity := entities.UsersFilter{
		ID:    model.ID,
		Email: model.Email,
	}
	return &filterEntity
}

// FindOneUserDataResponseJSON Parse to JSON
func (model *FindOneUserDataResponseJSON) Parse(data interface{}) (*FindOneUserDataResponseJSON, error) {
	err := copier.Copy(model, data)
	return model, err
}

// FindOneUserDataResponseSwagger Find user date response for swagger
type FindOneUserDataResponseSwagger struct {
	Status  string                      `json:"status" example:"success"`
	Message string                      `json:"message" example:"OK"`
	Data    FindOneUserDataResponseJSON `json:"data,omitempty"`
}
