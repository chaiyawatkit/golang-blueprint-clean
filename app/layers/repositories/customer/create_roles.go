package customer

import (
	"golang-blueprint-clean/app/entities"
	"golang-blueprint-clean/app/errors"
	"golang-blueprint-clean/app/layers/repositories/customer/models"
)

func (r *repo) CreateRoles(input *entities.Roles) (*entities.Roles, error) {
	model, _ := new(models.Roles).ParseRoleToDB(input)

	if err := r.Conn.Save(&model).Error; err != nil {
		return nil, errors.InternalError{Message: err.Error()}
	}

	return model.RoleEntity(), nil
}
