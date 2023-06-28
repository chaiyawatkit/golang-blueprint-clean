package entities

import (
	"time"
)

type UsersSignIn struct {
	SessionID string
}

type UsersFilter struct {
	Uuid *int
}

type Users struct {
	Uuid        int
	UserStatus  string
	Segment     string
	AccessToken *string
	LastLogin   *time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}
