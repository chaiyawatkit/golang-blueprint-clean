package back_office

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
	FindBanners(segment string) ([]entities.Banners, error)
}
