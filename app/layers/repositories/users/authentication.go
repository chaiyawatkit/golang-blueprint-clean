package users

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"golang-blueprint-clean/app/constants"
	"golang-blueprint-clean/app/env"
	"golang-blueprint-clean/app/utils"
	"io/ioutil"
	"net/http"
)

func (r *repo) Authentication(ctx context.Context, jwtAccessToken string) (*string, error, error) {
	var (
		errContextMsg = fmt.Sprintf("%s users from %s", constants.FailToAuthenticate, env.AppFeatureServiceUrl)
		requestURL    = fmt.Sprintf("%s/v1/users.auth", env.AppFeatureServiceUrl)
	)

	client := &http.Client{}
	request, err := http.NewRequest(http.MethodPost, requestURL, nil)
	if err != nil {
		return nil, errors.Wrap(err, errContextMsg), err
	}

	if ginContext, _ := utils.FromContextToGinContext(ctx); ginContext != nil {
		request.Header.Set(constants.CorrelationIdHeaderKey, ginContext.GetHeader(constants.CorrelationIdHeaderKey))
	}

	request.Header.Set(constants.XFinPlusAuth, jwtAccessToken)

	res, err := client.Do(request)
	if err != nil {
		return nil, errors.Wrap(err, errContextMsg), err
	}

	if res.StatusCode != http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(res.Body)
		defer res.Body.Close()
		if err != nil {
			return nil, errors.Wrap(err, errContextMsg), err
		}

		var responseBody map[string]interface{}
		err = json.Unmarshal(bodyBytes, &responseBody)
		if err != nil {
			return nil, errors.Wrap(err, errContextMsg), err
		}

		err = errors.New(responseBody["message"].(string))
		if res.StatusCode == http.StatusInternalServerError {
			return nil, errors.New("InternalServerError"), errors.New("ApiGatewayUnauthorized")
		} else if res.StatusCode == http.StatusUnauthorized {
			return nil, errors.New("ApiGatewayUnauthorized"), errors.New("ApiGatewayUnauthorized")
		}

		return nil, errors.Wrap(err, errContextMsg), err
	}

	accessToken := res.Header.Get(constants.XFinPlusAuth)

	return &accessToken, nil, nil
}
