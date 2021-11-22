package models

import (
	"github.com/jinzhu/copier"
	"golang-blueprint-clean/app/entities"
	"time"
)

type Roles struct {
	ID          uint       `gorm:"column:id;primary_key"`
	Code        string     `gorm:"column:code"`
	DisplayName string     `gorm:"column:display_name"`
	CreatedAt   time.Time  `gorm:"column:created_at"`
	UpdatedAt   time.Time  `gorm:"column:updated_at"`
	DeletedAt   *time.Time `gorm:"column:deleted_at"`
}

func (model *Roles) ParseRoleToDB(data interface{}) (*Roles, error) {
	err := copier.Copy(model, data)
	return model, err
}

func (model *Roles) RoleEntity() *entities.Roles {
	var entity entities.Roles
	_ = copier.Copy(&entity, model)

	return &entity
}
