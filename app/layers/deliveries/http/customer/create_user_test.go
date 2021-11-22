package customer_test

import (
	"github.com/gin-gonic/gin"
	"golang-blueprint-clean/app/errors"
	customerHandler "golang-blueprint-clean/app/layers/deliveries/http/customer"
	customerMock "golang-blueprint-clean/app/mocks/customer"
	testhelper "golang-blueprint-clean/app/test_helpers"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandler_CreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	customerMockUseCase := customerMock.NewMockUseCase(ctrl)

	executeWithRequest := func(mockUseCase *customerMock.MockUseCase, jsonRequestBody []byte) *httptest.ResponseRecorder {
		response := httptest.NewRecorder()
		_, ginEngine := gin.CreateTestContext(response)
		requestURL := "/v1/users"
		httpRequest, _ := http.NewRequest("POST", requestURL, strings.NewReader(string(jsonRequestBody)))

		customerHandler.NewEndpointHttpHandler(ginEngine, mockUseCase)
		ginEngine.ServeHTTP(response, httpRequest)
		return response
	}

	t.Run("Success", func(t *testing.T) {
		jsonRequestBody := []byte(`{
									"email": "tech@mail.co",
									"password": "P@ssw0rd",
									"firstName": "n",
									"lastName": "digital",
									"age": "20",
									"birthDate": "1999/09/09",
									"address": "92/7",
									"phoneNumber": "+66899999999",
									"provider": "OWN",
									"statusID": 1,
									"roleID": 1,
									"roleTypeID": 1,
									"createdBy": "N System"
								}`)

		userMock := testhelper.GetMockUserEntity()

		customerMockUseCase.EXPECT().
			CreateUser(gomock.Any()).
			Return(&userMock, nil)

		res := executeWithRequest(customerMockUseCase, jsonRequestBody)
		assert.Equal(t, http.StatusOK, res.Code)
	})

	t.Run("Failure: CreateUser Throw Error", func(t *testing.T) {
		jsonRequestBody := []byte(`{
									"email": "tech@mail.co",
									"password": "P@ssw0rd",
									"firstName": "n",
									"lastName": "digital",
									"age": "20",
									"birthDate": "1999/09/09",
									"address": "92/7",
									"phoneNumber": "+66899999999",
									"provider": "OWN",
									"statusID": 1,
									"roleID": 1,
									"roleTypeID": 1,
									"createdBy": "N System"
								}`)

		customerMockUseCase.EXPECT().
			CreateUser(gomock.Any()).
			Return(nil, errors.UnprocessableEntity{Message: "error here"})

		res := executeWithRequest(customerMockUseCase, jsonRequestBody)
		assert.Equal(t, http.StatusBadRequest, res.Code)
	})

	t.Run("Failure : Parse Error", func(t *testing.T) {
		jsonRequestBody := []byte(`{
									"email": "fake-email"
								}`)

		res := executeWithRequest(customerMockUseCase, jsonRequestBody)
		assert.Equal(t, http.StatusBadRequest, res.Code)
	})
}
