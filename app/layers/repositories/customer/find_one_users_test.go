package customer_test

import (
	gomocket "github.com/Selvatico/go-mocket"
	"golang-blueprint-clean/app/entities"
	"golang-blueprint-clean/app/layers/repositories/customer"
	testhelper "golang-blueprint-clean/app/test_helpers"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestRepo_FindOneUsers(t *testing.T) {
	testhelper.InitEnv()

	var (
		id    = uint(1)
		email = "tech@mail.co"

		queryPattern = `SELECT * FROM "users"`
		queryStub    = []map[string]interface{}{{
			"id":         id,
			"email":      email,
			"created_at": time.Time{},
		}}

		filters = entities.UsersFilter{
			ID:    &id,
			Email: &email,
		}
	)

	t.Run("Success", func(t *testing.T) {
		DB := testhelper.SetupMockDB()
		defer DB.Close()

		gomocket.Catcher.Reset().NewMock().WithQuery(queryPattern).WithReply(queryStub)

		CustomerRepo := customer.InitRepo(DB)

		results, err := CustomerRepo.FindOneUser(&filters)

		assert.NoError(t, err)
		assert.Equal(t, id, results.ID)
		assert.Equal(t, email, results.Email)
	})

	t.Run("Failure : got an internal error", func(t *testing.T) {
		DB := testhelper.SetupMockDB()
		defer DB.Close()

		gomocket.Catcher.Reset().NewMock().WithQuery(queryPattern).WithReply(queryStub).WithError(errors.New("error"))
		CustomerRepo := customer.InitRepo(DB)

		result, err := CustomerRepo.FindOneUser(&filters)

		assert.Error(t, err)
		assert.Nil(t, result)
	})
}
