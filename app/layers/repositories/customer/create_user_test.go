package customer_test

import (
	gomocket "github.com/Selvatico/go-mocket"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"golang-blueprint-clean/app/layers/repositories/customer"
	testhelper "golang-blueprint-clean/app/test_helpers"
	"testing"
)

func TestPostgresRepository_CreateUser(t *testing.T) {
	var (
		queryPattern  = `INSERT INTO "users"`
		expectedEmail = "tech@mail.co"
	)

	userEntity := testhelper.GetMockUserEntity()

	t.Run("Success", func(t *testing.T) {
		DB := testhelper.SetupMockDB()
		defer DB.Close()

		commonReply := []map[string]interface{}{{"email": expectedEmail}}
		gomocket.Catcher.Reset().NewMock().WithQuery(queryPattern).WithReply(commonReply)
		CustomerRepo := customer.InitRepo(DB)

		user, err := CustomerRepo.CreateUser(&userEntity)

		assert.NoError(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, expectedEmail, user.Email)
	})

	t.Run("Failure", func(t *testing.T) {
		DB := testhelper.SetupMockDB()
		defer DB.Close()

		gomocket.Catcher.Reset().NewMock().WithError(errors.New("error"))

		CustomerRepo := customer.InitRepo(DB)

		user, err := CustomerRepo.CreateUser(nil)

		assert.Error(t, err)
		assert.Nil(t, user)
	})
}
