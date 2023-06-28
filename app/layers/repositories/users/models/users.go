package models

import (
	"golang-blueprint-clean/app/entities"
	"log"
	"time"
)

type Users struct {
	Uuid                     int
	TwilioSid                *string
	UuidFirebasePhoneNumber  *string
	UuidFirebaseAppleAccount *string
	UuidFirebaseGmail        *string
	NotificationToken        *string
	Slug                     *string
	PhoneNumber              *string
	Gmail                    *string
	AppleUid                 *string
	Email                    *string
	Password                 *string
	Code                     *string
	ExpCode                  *time.Time
	IsEmailVerify            *bool
	UserStatus               string
	AccessToken              *string
	LastLogin                *time.Time
	CreatedAt                time.Time
	UpdatedAt                time.Time
	DeletedAt                *time.Time
	CreateEmailAt            *time.Time
	ForgotPasswordCode       *string
	OfficerUpdateId          *int
	OfficerUpdateAt          *time.Time
}

type UsersListModel []Users

// ToEntity convert model to entity
func (m Users) ToEntity() entities.Users {
	log.Printf(">>>>m%+v", m)
	return entities.Users{
		Uuid:        m.Uuid,
		UserStatus:  m.UserStatus,
		AccessToken: m.AccessToken,
		LastLogin:   m.LastLogin,
		CreatedAt:   m.CreatedAt,
		UpdatedAt:   m.UpdatedAt,
	}
}

// ToEntities convert model to entities
func (m UsersListModel) ToEntities() []entities.Users {
	var results []entities.Users
	for _, v := range m {
		results = append(results, v.ToEntity())
	}
	return results
}
