package models

import (
	"github.com/gin-gonic/gin"
	"golang-blueprint-clean/app/entities"
	"golang-blueprint-clean/app/errors"
)

// CreateUserRequestJSON receive request data to create new record
type CreateUserRequestJSON struct {
	Email       string  `json:"email" binding:"required" validate:"email" example:"tech@mail.co"`
	Password    string  `json:"password" binding:"required" example:"tech@mail.co"`
	FirstName   string  `json:"firstName" binding:"required" example:"tech@mail.co"`
	LastName    string  `json:"lastName" binding:"required" example:"tech@mail.co"`
	Age         *string `json:"age" example:"20"`
	BirthDate   *string `json:"birthDate" example:"1999/09/09"`
	Address     *string `json:"address" example:"9999/99 Parkred"`
	PhoneNumber *string `json:"phoneNumber" example:"+66999999999"`
	Provider    string  `json:"provider" binding:"required" example:"OWN"`
	StatusID    uint    `json:"statusId" binding:"required" example:"1"`
	RoleID      uint    `json:"roleId" binding:"required" example:"1"`
	RoleTypeID  uint    `json:"roleTypeId" binding:"required" example:"1"`
	CreatedBy   string  `json:"createdBy" binding:"required" example:"N API"`
}

type CreateUserResponseJSON struct {
	Email       string  `json:"email" binding:"required" validate:"email" example:"tech@mail.co.th"`
	Password    string  `json:"password" binding:"required" example:"tech@mail.co"`
	FirstName   string  `json:"firstName" binding:"required" example:"tech@mail.co"`
	LastName    string  `json:"lastName" binding:"required" example:"tech@mail.co"`
	Age         *string `json:"age" example:"20"`
	BirthDate   *string `json:"birthDate" example:"1999/09/09"`
	Address     *string `json:"address" example:"92/7 Parkred"`
	PhoneNumber *string `json:"phoneNumber" example:"+66999999999"`
	Provider    string  `json:"provider" binding:"required" example:"OWN"`
	StatusID    uint    `json:"statusId" binding:"required" example:"1"`
	RoleID      uint    `json:"roleId" binding:"required" example:"1"`
	RoleTypeID  uint    `json:"roleTypeId" binding:"required" example:"1"`
	CreatedBy   string  `json:"createdBy" binding:"required" example:"N API"`
}

// Parse parse data to create new record
func (model *CreateUserRequestJSON) Parse(c *gin.Context) (*CreateUserRequestJSON, error) {
	err := c.ShouldBindJSON(model)
	if err != nil {
		return nil, errors.ParameterError{Message: err.Error()}
	}

	return model, nil
}

// Entity convert CreateUserRequestJSON to Users entity
func (model *CreateUserRequestJSON) Entity() *entities.Users {
	userEntity := entities.Users{
		Email:       model.Email,
		Password:    model.Password,
		FirstName:   model.FirstName,
		LastName:    model.LastName,
		Age:         model.Age,
		BirthDate:   model.BirthDate,
		Address:     model.Address,
		PhoneNumber: model.PhoneNumber,
		Provider:    model.Provider,
		StatusID:    model.StatusID,
		RoleID:      model.RoleID,
		RoleTypeID:  model.RoleTypeID,
		CreatedBy:   model.CreatedBy,
	}
	return &userEntity
}

// CreateUserResponse create auth response
type CreateUserResponse struct {
	Status  string                 `json:"status" example:"success"`
	Message string                 `json:"message" example:"OK"`
	Data    CreateUserResponseJSON `json:"data,omitempty"`
}
