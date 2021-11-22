package customer_test

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"golang-blueprint-clean/app/entities"
	"golang-blueprint-clean/app/errors"
	customerUseCase "golang-blueprint-clean/app/layers/usecases/customer"
	customerMock "golang-blueprint-clean/app/mocks/customer"
	testhelper "golang-blueprint-clean/app/test_helpers"
	"testing"
)

func TestUseCase_CreateUser(t *testing.T) {
	testhelper.InitEnv()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	customerMockRepo := customerMock.NewMockRepo(ctrl)
	expectedUserSlug := "fd9b22a9-30bf-4cfe-8aee-65581ec88a9b"

	t.Run("Happy", func(t *testing.T) {
		userEntity := testhelper.GetMockUserEntity()

		customerMockRepo.EXPECT().
			CreateUser(&userEntity).
			Return(&userEntity, nil)

		useCase := customerUseCase.InitUseCase(customerMockRepo)
		user, err := useCase.CreateUser(&userEntity)
		assert.Nil(t, err)
		assert.Equal(t, user.UserSlug, expectedUserSlug)
	})

	t.Run("Fail: repo return error", func(t *testing.T) {
		userEntity := testhelper.GetMockUserEntity()

		customerMockRepo.EXPECT().
			CreateUser(&userEntity).
			Return(nil, errors.InternalError{Message: "error here"})

		useCase := customerUseCase.InitUseCase(customerMockRepo)
		_, err := useCase.CreateUser(&userEntity)
		assert.Error(t, err)
	})

	t.Run("Fail: can't create users", func(t *testing.T) {
		userEntity := entities.Users{}

		customerMockRepo.EXPECT().
			CreateUser(&userEntity).
			Return(nil, errors.InternalError{Message: "error here"})

		useCase := customerUseCase.InitUseCase(customerMockRepo)
		_, err := useCase.CreateUser(&userEntity)
		assert.Error(t, err)
	})
}
