package customer_test

import (
	gomocket "github.com/Selvatico/go-mocket"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"golang-blueprint-clean/app/entities"
	"golang-blueprint-clean/app/layers/repositories/customer"
	testhelper "golang-blueprint-clean/app/test_helpers"
	"testing"
	"time"
)

func TestRepo_FindOneRole(t *testing.T) {
	testhelper.InitEnv()

	var (
		id   = uint(1)
		code = "ADMIN"

		queryPattern = `SELECT * FROM "roles"`
		queryStub    = []map[string]interface{}{{
			"id":         id,
			"code":       code,
			"created_at": time.Time{},
		}}

		filters = entities.RolesFilter{
			ID: &id,
		}
	)

	t.Run("Success", func(t *testing.T) {
		DB := testhelper.SetupMockDB()
		defer DB.Close()

		gomocket.Catcher.Reset().NewMock().WithQuery(queryPattern).WithReply(queryStub)

		CustomerRepo := customer.InitRepo(DB)

		results, err := CustomerRepo.FindOneRole(&filters)

		assert.NoError(t, err)
		assert.Equal(t, id, results.ID)
		assert.Equal(t, code, results.Code)
	})

	t.Run("Failure : got an internal error", func(t *testing.T) {
		DB := testhelper.SetupMockDB()
		defer DB.Close()

		gomocket.Catcher.Reset().NewMock().WithQuery(queryPattern).WithReply(queryStub).WithError(errors.New("error"))
		CustomerRepo := customer.InitRepo(DB)

		result, err := CustomerRepo.FindOneRole(&filters)

		assert.Error(t, err)
		assert.Nil(t, result)
	})
}
