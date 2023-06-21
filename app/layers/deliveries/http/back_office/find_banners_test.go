package back_office_test

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"golang-blueprint-clean/app/errors"
	backOfficeHandler "golang-blueprint-clean/app/layers/deliveries/http/back_office"
	backOfficeMock "golang-blueprint-clean/app/mocks/back_office"
	testhelper "golang-blueprint-clean/app/test_helpers"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestHandler_FindBanners(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	backOfficeMockUseCase := backOfficeMock.NewMockUseCase(ctrl)

	var (
		segment = "pb"
		baseURL = "/v1/banners"
	)

	executeWithRequest := func(mockUseCase *backOfficeMock.MockUseCase, segment string) *httptest.ResponseRecorder {
		response := httptest.NewRecorder()
		_, ginEngine := gin.CreateTestContext(response)

		var requestURL *url.URL

		requestURL, _ = url.Parse(fmt.Sprintf("/v1/banners?segment=%s", segment))

		httpRequest, _ := http.NewRequest("GET", requestURL.String(), nil)
		backOfficeHandler.NewEndpointHttpHandler(ginEngine, mockUseCase)
		ginEngine.ServeHTTP(response, httpRequest)
		return response
	}

	executeWithRequest2 := func(mockUseCase *backOfficeMock.MockUseCase, requestURL *url.URL) *httptest.ResponseRecorder {
		response := httptest.NewRecorder()
		_, ginEngine := gin.CreateTestContext(response)
		httpRequest, _ := http.NewRequest("GET", requestURL.String(), nil)
		backOfficeHandler.NewEndpointHttpHandler(ginEngine, mockUseCase)
		ginEngine.ServeHTTP(response, httpRequest)
		return response
	}

	t.Run("Success : find banner list", func(t *testing.T) {
		userMock := testhelper.GetMockBannerList()
		backOfficeMockUseCase.EXPECT().FindBanners(segment).Return(userMock, nil)
		res := executeWithRequest(backOfficeMockUseCase, segment)
		assert.Equal(t, http.StatusOK, res.Code)
	})

	t.Run("Failure : find banner list error", func(t *testing.T) {

		backOfficeMockUseCase.EXPECT().FindBanners(segment).Return(nil, errors.UnprocessableEntity{Message: "err"})

		res := executeWithRequest(backOfficeMockUseCase, segment)

		assert.Equal(t, http.StatusBadRequest, res.Code)
	})

	t.Run("Failure : Parameter empty", func(t *testing.T) {
		var requestURL *url.URL
		requestURL, _ = url.Parse(fmt.Sprintf("%s", baseURL))

		res := executeWithRequest2(nil, requestURL)

		assert.Equal(t, http.StatusBadRequest, res.Code)
	})

	t.Run("Failure : invalid sagment", func(t *testing.T) {
		var requestURL *url.URL

		requestURL, _ = url.Parse(fmt.Sprintf("%s?segment=", baseURL))
		res := executeWithRequest2(nil, requestURL)
		assert.Equal(t, http.StatusBadRequest, res.Code)
	})
}