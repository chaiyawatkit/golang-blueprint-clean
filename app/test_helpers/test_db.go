package testhelper

import (
	"bytes"
	"database/sql"
	gomocket "github.com/Selvatico/go-mocket"
	"github.com/gin-gonic/gin"
	"golang-blueprint-clean/app/entities"
	"log"
	"net/http"
	"net/http/httptest"
)

func GetMockBannerList() []entities.Banners {
	mockBannerEntity := []entities.Banners{
		{
			ID:        1,
			Title:     "kkkk",
			Thumbnail: "https://www.google.com/",
			Status:    "y",
			Type:      "internal",
			Redirect:  "https://www.google.com/",
			CreatedAt: 1687313433,
			EndAt:     1687313433,
			Segment:   "pb",
		},
	}

	return mockBannerEntity
}

func SetupMockDB() *sql.DB {
	gomocket.Catcher.Register()

	db, err := sql.Open(gomocket.DriverName, "")
	if err != nil {
		log.Fatalf("error mocking up the database %s", err)
	}

	//db.LogMode(true)

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
