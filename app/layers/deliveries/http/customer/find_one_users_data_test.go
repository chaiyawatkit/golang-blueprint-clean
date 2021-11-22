package customer_test

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"golang-blueprint-clean/app/entities"
	"golang-blueprint-clean/app/errors"
	customerHandler "golang-blueprint-clean/app/layers/deliveries/http/customer"
	customerMock "golang-blueprint-clean/app/mocks/customer"
	testhelper "golang-blueprint-clean/app/test_helpers"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestHandler_FindOneUserData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	customerMockUseCase := customerMock.NewMockUseCase(ctrl)

	var (
		baseURL = "/v1/users.data"
		id      = uint(1)
		email   = "tech@mail.co"
	)

	executeWithRequest := func(mockUseCase *customerMock.MockUseCase,
		uid uint,
		email string) *httptest.ResponseRecorder {

		response := httptest.NewRecorder()
		_, ginEngine := gin.CreateTestContext(response)

		var requestURL *url.URL

		requestURL, _ = url.Parse(fmt.Sprintf("/v1/users.data?id=%d&email=%s",
			uid,
			email,
		))

		httpRequest, _ := http.NewRequest("GET", requestURL.String(), nil)
		customerHandler.NewEndpointHttpHandler(ginEngine, mockUseCase)
		ginEngine.ServeHTTP(response, httpRequest)
		return response
	}

	executeWithRequest2 := func(mockUseCase *customerMock.MockUseCase, requestURL *url.URL) *httptest.ResponseRecorder {
		response := httptest.NewRecorder()
		_, ginEngine := gin.CreateTestContext(response)
		httpRequest, _ := http.NewRequest("GET", requestURL.String(), nil)
		customerHandler.NewEndpointHttpHandler(ginEngine, mockUseCase)
		ginEngine.ServeHTTP(response, httpRequest)
		return response
	}

	t.Run("Success : find one user", func(t *testing.T) {
		userMock := testhelper.GetMockUserEntity()

		filter := entities.UsersFilter{
			ID:    &id,
			Email: &email,
		}

		customerMockUseCase.EXPECT().FindOneUserData(&filter).Return(&userMock, nil)

		res := executeWithRequest(customerMockUseCase, id, email)
		assert.Equal(t, http.StatusOK, res.Code)
	})

	t.Run("Failure : find one user error", func(t *testing.T) {
		filter := entities.UsersFilter{
			ID:    &id,
			Email: &email,
		}

		customerMockUseCase.EXPECT().
			FindOneUserData(gomock.AssignableToTypeOf(&filter)).
			Return(nil, errors.UnprocessableEntity{Message: "err"})

		res := executeWithRequest(customerMockUseCase, id, email)

		assert.Equal(t, http.StatusBadRequest, res.Code)
	})

	t.Run("Failure : Parameter empty", func(t *testing.T) {
		var requestURL *url.URL
		requestURL, _ = url.Parse(fmt.Sprintf("%s", baseURL))

		res := executeWithRequest2(nil, requestURL)

		assert.Equal(t, http.StatusBadRequest, res.Code)
	})
}
