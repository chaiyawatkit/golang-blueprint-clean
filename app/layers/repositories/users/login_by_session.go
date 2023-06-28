package users

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
	"golang-blueprint-clean/app/constants"
	"golang-blueprint-clean/app/entities"
	"golang-blueprint-clean/app/env"
	"golang-blueprint-clean/app/layers/repositories/users/models"
	"io/ioutil"
	"net/http"
)

func (r *repo) LoginBySession(ctx context.Context, input entities.UsersSignIn) (*string, error, error) {
	client := resty.New()

	errContextMsg := fmt.Sprintf("%s users from %s", "FailToSignIn", env.AppFeatureServiceUrl)
	requestBody, err := json.Marshal(models.LoginBySession{
		SessionID: input.SessionID,
	})
	if err != nil {
		return nil, errors.Wrap(err, errContextMsg), err
	}

	requestURL := fmt.Sprintf("%s/v1/users.session.login", env.AppFeatureServiceUrl)
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(bytes.NewBuffer(requestBody)).
		Post(requestURL)
	if err != nil {
		return nil, errors.Wrap(err, errContextMsg), err
	}

	if resp.StatusCode() != http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.RawBody())
		defer resp.RawBody().Close()
		if err != nil {
			return nil, errors.Wrap(err, errContextMsg), err
		}

		var responseBody map[string]interface{}
		err = json.Unmarshal(bodyBytes, &responseBody)
		if err != nil {
			return nil, errors.Wrap(err, errContextMsg), err
		}

		err = errors.New(responseBody["message"].(string))

		if err.Error() == constants.YourAccountWasDeleted {
			return nil, errors.Wrap(err, errContextMsg), err
		}

		//if err.Error() == annoyer.InvalidLoginCredential || err.Error() == "fail to get data from database" {
		//	return nil, errors.New(annoyer.InvalidLoginCredential), err
		//}
		//
		//if resp.StatusCode() == http.StatusInternalServerError {
		//	return nil, errors.New(annoyer.InternalServerError), errors.New(annoyer.InternalServerError)
		//} else if resp.StatusCode() == http.StatusUnauthorized {
		//	return nil, errors.New(annoyer.InvalidLoginCredential), errors.New(annoyer.InvalidLoginCredential)
		//}

		return nil, errors.Wrap(err, errContextMsg), err
	}

	accessToken := resp.Header().Get(constants.XFinPlusAuth)
	if accessToken == "" {
		return nil, errors.New(constants.InvalidResponseFormatError), errors.New(constants.InvalidResponseFormatError)
	}

	return &accessToken, nil, nil
}
