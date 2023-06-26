package back_office

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
	"golang-blueprint-clean/app/constants"
	"golang-blueprint-clean/app/entities"
	"golang-blueprint-clean/app/env"
	"golang-blueprint-clean/app/layers/repositories/back_office/models"
	"net/url"
)

func (r *repo) FindBanners(filter *entities.SegmentTypes) ([]entities.Banners, error, error) {
	client := resty.New()
	errContextMsg := fmt.Sprintf("%s app-feature from %s", constants.FailToGet, env.AppFeatureServiceUrl)
	queryUrl := findBannerQueryString(filter)
	requestUrl := fmt.Sprintf("%s/v1/banners?%s", env.AppFeatureServiceUrl, queryUrl)
	var responseModel models.ResponseData
	resp, err := client.R().SetHeader("Accept", "application/json").SetResult(responseModel).Get(requestUrl)
	if err != nil {
		return nil, err, errors.Wrap(err, errContextMsg)
	}

	result := resp.Result().(*models.ResponseData)
	output := result.BannerEntity()
	return output, nil, nil
}

func findBannerQueryString(input *entities.SegmentTypes) string {
	queryString := url.Values{}
	if input.SegmentType != nil {
		queryString.Set("segment", fmt.Sprintf("%s", *input.SegmentType))
	}

	return queryString.Encode()
}
