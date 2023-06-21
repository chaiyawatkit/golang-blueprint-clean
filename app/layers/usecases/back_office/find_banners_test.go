package back_office_test

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"golang-blueprint-clean/app/errors"
	backOfficeUseCase "golang-blueprint-clean/app/layers/usecases/back_office"
	backOfficeMock "golang-blueprint-clean/app/mocks/back_office"
	testhelper "golang-blueprint-clean/app/test_helpers"
	"testing"
)

func TestUseCase_FindBanners(t *testing.T) {
	testhelper.InitEnv()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	backOfficeMockRepo := backOfficeMock.NewMockRepo(ctrl)

	var (
		segment        = "pb"
		expectedUserId = uint(1)
		expectedBanner = "kkkk"
	)

	t.Run("Happy", func(t *testing.T) {
		bannerEntity := testhelper.GetMockBannerList()

		backOfficeMockRepo.EXPECT().FindBanners(segment).Return(bannerEntity, nil)

		useCase := backOfficeUseCase.InitUseCase(backOfficeMockRepo)
		banners, err := useCase.FindBanners(segment)
		banner := banners[0]
		assert.Nil(t, err)
		assert.Equal(t, expectedUserId, banner.ID)
		assert.Equal(t, expectedBanner, banner.Title)

	})

	t.Run("Fail: repo FindBanner return error", func(t *testing.T) {

		backOfficeMockRepo.EXPECT().
			FindBanners(segment).
			Return(nil, errors.InternalError{Message: "error here"})

		useCase := backOfficeUseCase.InitUseCase(backOfficeMockRepo)
		_, err := useCase.FindBanners(segment)
		assert.Error(t, err)
	})

}
