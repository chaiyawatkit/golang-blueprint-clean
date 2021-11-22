package models

import (
	"github.com/jinzhu/copier"
	"golang-blueprint-clean/app/entities"
	"time"
)

type Users struct {
	ID          uint
	Email       string     `gorm:"column:email"`
	Password    string     `gorm:"column:password"`
	FirstName   string     `gorm:"column:first_name"`
	LastName    string     `gorm:"column:last_name"`
	Age         *string    `gorm:"column:age"`
	BirthDate   *string    `gorm:"column:birth_date"`
	Address     *string    `gorm:"column:address"`
	PhoneNumber *string    `gorm:"column:phone_number"`
	AccessToken *string    `gorm:"column:access_token"`
	Provider    string     `gorm:"column:provider"`
	LasLogin    *time.Time `gorm:"column:last_login"`
	StatusID    uint       `gorm:"column:status_id"`
	RoleID      uint       `gorm:"column:role_id"`
	RoleTypeID  uint       `gorm:"column:role_type_id"`
	CreatedAt   time.Time  `gorm:""`
	CreatedBy   string     `gorm:"column:created_by"`
	UpdatedAt   time.Time  `gorm:""`
	UpdatedBy   string     `gorm:"column:updated_by"`
	DeletedAt   *time.Time `gorm:""`
	DeletedBy   *string    `gorm:""`
	number      int
}

func (model *Users) ParseUserToDB(data interface{}) (*Users, error) {
	err := copier.Copy(model, data)
	return model, err
}

func (model *Users) UserEntity() *entities.Users {
	var entity entities.Users
	_ = copier.Copy(&entity, model)

	return &entity
}
