package entities

import "time"

type RoleTypes struct {
	ID          uint
	Code        string
	DisplayName string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}
