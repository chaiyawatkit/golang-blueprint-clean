package back_office_test

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"golang-blueprint-clean/app/entities"
	"golang-blueprint-clean/app/errors"
	backOfficeUseCase "golang-blueprint-clean/app/layers/usecases/back_office"
	backOfficeMock "golang-blueprint-clean/app/mocks/back_office"
	testhelper "golang-blueprint-clean/app/test_helpers"
	"golang-blueprint-clean/app/utils"
	"testing"
)

func TestUseCase_FindBanners(t *testing.T) {
	testhelper.InitEnv()
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()
	backOfficeMockRepo := backOfficeMock.NewMockRepo(ctrl)

	var (
		segment = utils.ToString("pb")
	)

	t.Run("Happy", func(t *testing.T) {
		bannerEntity := testhelper.GetMockBannerList()
		segmentType := &entities.SegmentTypes{
			SegmentType: segment,
		}
		backOfficeMockRepo.EXPECT().FindBanners(segmentType).Return(bannerEntity, nil, nil)
		useCase := backOfficeUseCase.InitUseCase(backOfficeMockRepo)
		banners, err, unWrap := useCase.FindBanners(*segmentType)
		assert.Nil(t, err)
		assert.Nil(t, unWrap)
		assert.NotNil(t, banners)

	})

	t.Run("Fail: repo FindBanner return error", func(t *testing.T) {
		segmentType := &entities.SegmentTypes{
			SegmentType: segment,
		}
		backOfficeMockRepo.EXPECT().
			FindBanners(segmentType).
			Return(nil, errors.InternalError{Message: "error here"}, errors.InternalError{Message: "error here"})

		useCase := backOfficeUseCase.InitUseCase(backOfficeMockRepo)
		_, err, unWrap := useCase.FindBanners(*segmentType)
		assert.Error(t, err)
		assert.Error(t, unWrap)
	})

}
