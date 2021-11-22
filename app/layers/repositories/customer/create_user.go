package customer

import (
	"golang-blueprint-clean/app/entities"
	"golang-blueprint-clean/app/errors"
	"golang-blueprint-clean/app/layers/repositories/customer/models"
)

func (r *repo) CreateUser(input *entities.Users) (*entities.Users, error) {
	model, _ := new(models.Users).ParseUserToDB(input)

	if err := r.Conn.Save(&model).Error; err != nil {
		return nil, errors.InternalError{Message: err.Error()}
	}

	return model.UserEntity(), nil
}
