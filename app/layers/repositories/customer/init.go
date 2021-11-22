package customer

import (
	"github.com/jinzhu/gorm"
	"golang-blueprint-clean/app/entities"
)

type repo struct {
	Conn *gorm.DB
}

// InitRepo CRUD
func InitRepo(Conn *gorm.DB) Repo {
	return &repo{Conn: Conn}
}

type Repo interface {
	CreateUser(input *entities.Users) (*entities.Users, error)
	CreateRoles(input *entities.Roles) (*entities.Roles, error)
	FindOneUser(filter *entities.UsersFilter) (*entities.Users, error)
	FindOneRole(filter *entities.RolesFilter) (*entities.Roles, error)
}
