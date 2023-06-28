package users

import (
	"database/sql"
	"golang-blueprint-clean/app/entities"
)

type repo struct {
	Conn *sql.DB
}

// InitRepo CRUD
func InitRepo(Conn *sql.DB) Repo {
	return &repo{Conn: Conn}
}

type Repo interface {
	FindUserByAccessToken(accessToken string) (*entities.Users, error)
}
