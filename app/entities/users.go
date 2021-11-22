package entities

import (
	"time"
)

type Users struct {
	ID          uint
	UserSlug    string
	Email       string
	Password    string
	FirstName   string
	LastName    string
	Age         *string
	BirthDate   *string
	Address     *string
	PhoneNumber *string
	AccessToken *string
	Provider    string
	LasLogin    *time.Time
	StatusID    uint
	Statuses    *Statuses
	RoleID      uint
	Roles       *Roles
	RoleTypeID  uint
	RoleTypes   *RoleTypes
	CreatedAt   time.Time
	CreatedBy   string
	UpdatedAt   time.Time
	UpdatedBy   string
	DeletedAt   *time.Time
	DeletedBy   *string
}

// UsersFilter for filtering user
type UsersFilter struct {
	ID    *uint
	Email *string
}
