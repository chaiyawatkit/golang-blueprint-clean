package entities

import "time"

type Roles struct {
	ID          uint
	Code        string
	DisplayName string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}

type RolesFilter struct {
	ID *uint
}
