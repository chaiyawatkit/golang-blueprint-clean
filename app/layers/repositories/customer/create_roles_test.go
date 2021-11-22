package customer_test

import (
	gomocket "github.com/Selvatico/go-mocket"
	"golang-blueprint-clean/app/layers/repositories/customer"
	testhelper "golang-blueprint-clean/app/test_helpers"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPostgresRepository_CreateRole(t *testing.T) {
	var (
		queryPattern        = `INSERT INTO "roles"`
		expectedCode        = "CodeTest"
		expectedDisplayName = "DisplayNameTest"
	)

	roleEntity := testhelper.GetMockRoleEntity()

	t.Run("Success", func(t *testing.T) {
		DB := testhelper.SetupMockDB()
		defer DB.Close()

		commonReply := []map[string]interface{}{{
			"code":         expectedCode,
			"display_name": expectedDisplayName,
		}}

		gomocket.Catcher.Reset().NewMock().WithQuery(queryPattern).WithReply(commonReply)

		CustomerRepo := customer.InitRepo(DB)

		role, err := CustomerRepo.CreateRoles(&roleEntity)

		assert.NoError(t, err)
		assert.NotNil(t, role)
		assert.Equal(t, expectedCode, role.Code)
		assert.Equal(t, expectedDisplayName, role.DisplayName)
	})

	t.Run("Failure", func(t *testing.T) {
		DB := testhelper.SetupMockDB()
		defer DB.Close()

		gomocket.Catcher.Reset().NewMock().WithError(errors.New("error"))

		CustomerRepo := customer.InitRepo(DB)

		user, err := CustomerRepo.CreateRoles(nil)
		expErr := "error"

		assert.Error(t, err)
		assert.Nil(t, user)
		assert.Equal(t, expErr, err.Error())
	})
}
