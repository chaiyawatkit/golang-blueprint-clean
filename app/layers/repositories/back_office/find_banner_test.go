package back_office_test

import (
	gomocket "github.com/Selvatico/go-mocket"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"golang-blueprint-clean/app/layers/repositories/back_office"
	testhelper "golang-blueprint-clean/app/test_helpers"
	"testing"
)

func TestRepo_FindBanners(t *testing.T) {
	testhelper.InitEnv()

	var (
		id        = uint(1)
		title     = "test"
		thumbnail = "test"
		createdAt = 1687350715
		endAt     = 1687350715
		status    = "y"
		redirect  = "www.gg.cc"
		typeMock  = "internal"

		segment = "pb"

		queryPattern = `SELECT ID,TITLE,THUMBNAIL,CREATED_AT,STATUS,TYPE,REDIRECT,END_AT,SEGMENT FROM BANNERS WHERE SEGMENT='pb'`
		queryStub    = []map[string]interface{}{
			{
				"ID":         id,
				"TITLE":      title,
				"THUMBNAIL":  thumbnail,
				"CREATED_AT": createdAt,
				"STATUS":     status,
				"TYPE":       typeMock,
				"REDIRECT":   redirect,
				"END_AT":     endAt,
				"SEGMENT":    segment,
			},
		}

		queryFailure = []map[string]interface{}{
			{
				"Title": title,
			},
		}
	)

	t.Run("Success", func(t *testing.T) {
		DB := testhelper.SetupMockDB()
		defer DB.Close()

		gomocket.Catcher.Reset().NewMock().WithQuery(queryPattern).WithReply(queryStub)
		BackOfficeRepo := back_office.InitRepo(DB)
		results, err := BackOfficeRepo.FindBanners(segment)
		result := results[0]
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, id, result.ID)
		assert.Equal(t, title, result.Title)
		assert.Equal(t, thumbnail, result.Thumbnail)
		assert.Equal(t, createdAt, result.CreatedAt)
		assert.Equal(t, status, result.Status)
		assert.Equal(t, typeMock, result.Type)
		assert.Equal(t, redirect, result.Redirect)
		assert.Equal(t, endAt, result.EndAt)
		assert.Equal(t, segment, result.Segment)

	})

	t.Run("Failure Scan ", func(t *testing.T) {
		DB := testhelper.SetupMockDB()
		defer DB.Close()

		gomocket.Catcher.Reset().NewMock().WithQuery(queryPattern).WithReply(queryFailure)
		BackOfficeRepo := back_office.InitRepo(DB)
		results, err := BackOfficeRepo.FindBanners(segment)
		assert.Error(t, err)
		assert.Nil(t, results)

	})

	t.Run("Failure : got an internal error", func(t *testing.T) {
		DB := testhelper.SetupMockDB()
		defer DB.Close()

		gomocket.Catcher.Reset().NewMock().WithQuery(queryPattern).WithReply(queryStub).WithError(errors.New("error"))
		BackOfficeRepo := back_office.InitRepo(DB)
		result, err := BackOfficeRepo.FindBanners(segment)
		assert.Error(t, err)
		assert.Nil(t, result)
	})
}
