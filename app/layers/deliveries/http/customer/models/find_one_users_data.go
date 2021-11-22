package models

import (
	"github.com/gin-gonic/gin"
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
	err := c.ShouldBindQuery(model)
	if err != nil {
		return nil, errors.ParameterError{Message: err.Error()}
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
func (model *FindOneUserDataResponseJSON) Parse(data entities.Users) (*FindOneUserDataResponseJSON, error) {
	dataRoleResponse := roleResponse{
		ID:          data.Roles.ID,
		Code:        data.Roles.Code,
		DisplayName: data.Roles.DisplayName,
	}

	dataResponse := FindOneUserDataResponseJSON{
		Email:       data.Email,
		FirstName:   data.FirstName,
		LastName:    data.LastName,
		Age:         data.Age,
		BirthDate:   data.BirthDate,
		Address:     data.Address,
		PhoneNumber: data.PhoneNumber,
		Provider:    data.Provider,
		StatusID:    data.StatusID,
		Role:        dataRoleResponse,
		RoleTypeID:  data.RoleTypeID,
	}

	return &dataResponse, nil
}

// FindOneUserDataResponseSwagger Find user date response for swagger
type FindOneUserDataResponseSwagger struct {
	Status  string                      `json:"status" example:"success"`
	Message string                      `json:"message" example:"OK"`
	Data    FindOneUserDataResponseJSON `json:"data,omitempty"`
}
