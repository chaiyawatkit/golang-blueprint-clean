package back_office

import (
	"fmt"
	"golang-blueprint-clean/app/entities"
	"golang-blueprint-clean/app/errors"
	"golang-blueprint-clean/app/layers/repositories/back_office/models"
	"log"
)

func (r *repo) FindBanners(segment string) ([]entities.Banners, error) {
	result := models.BannersArray{}
	query := fmt.Sprintf("SELECT ID,TITLE,THUMBNAIL,CREATED_AT,STATUS,TYPE,REDIRECT,END_AT,SEGMENT FROM BANNERS WHERE SEGMENT='%s'", segment)
	rows, err := r.Conn.Query(query)
	if err != nil {
		log.Printf(">>>err%+v", err)
		return nil, errors.InternalError{Message: err.Error()}
	}
	defer rows.Close()

	for rows.Next() {

		banner := models.Banners{}
		err := rows.Scan(&banner.ID, &banner.Title, &banner.Thumbnail, &banner.CreatedAt, &banner.Status, &banner.Type, &banner.Redirect, &banner.EndAt, &banner.Segment)
		if err != nil {

			return nil, errors.InternalError{Message: err.Error()}
		}

		result = append(result, banner)
	}

	output := result.BannerEntity()

	return output, nil
}
