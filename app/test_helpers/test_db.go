package testhelper

import (
	"bytes"
	gomocket "github.com/Selvatico/go-mocket"
	"github.com/gin-gonic/gin"
	"golang-blueprint-clean/app/entities"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
	"net/http/httptest"
	"time"
)

// GetMockUserEntity for unit test
func GetMockUserEntity() entities.Users {
	age := "20"
	birthDate := "1999/09/19"
	address := "9999/99"
	phoneNumber := "0999999999"
	mockUserEntity := entities.Users{
		ID:          1,
		UserSlug:    "fd9b22a9-30bf-4cfe-8aee-65581ec88a9b",
		Email:       "tech@mail.co",
		Password:    "P@ssw0rd",
		FirstName:   "dev",
		LastName:    "tech",
		Age:         &age,
		BirthDate:   &birthDate,
		Address:     &address,
		PhoneNumber: &phoneNumber,
		AccessToken: nil,
		Provider:    "OWN",
		LasLogin:    nil,
		StatusID:    1,
		RoleID:      1,
		RoleTypeID:  1,
		CreatedAt:   time.Time{},
		CreatedBy:   "SYSTEM",
		UpdatedAt:   time.Time{},
		UpdatedBy:   "SYSTEM",
		DeletedAt:   nil,
		DeletedBy:   nil,
	}
	return mockUserEntity
}

// GetMockRoleEntity for unit test
func GetMockRoleEntity() entities.Roles {
	mockRoleEntity := entities.Roles{
		ID:          1,
		Code:        "CodeTest",
		DisplayName: "DisplayNameTest",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	}
	return mockRoleEntity
}

func SetupMockDB() *gorm.DB {
	gomocket.Catcher.Register()

	db, err := gorm.Open(gomocket.DriverName, "")
	if err != nil {
		log.Fatalf("error mocking up the database %s", err)
	}

	db.LogMode(true)

	return db
}

func MakeStubContext(method string, url string, params string) (c *gin.Context) {
	const MIMEJSON = "application/json"

	body := bytes.NewBufferString(params)

	context, _ := gin.CreateTestContext(httptest.NewRecorder())
	context.Request, _ = http.NewRequest(method, url, body)
	context.Request.Header.Add("Content-Type", MIMEJSON)

	return context
}
