package users

import (
	"fmt"
	"golang-blueprint-clean/app/entities"
	"golang-blueprint-clean/app/errors"
	"golang-blueprint-clean/app/layers/repositories/users/models"
)

func (r *repo) FindUserByAccessToken(accessToken string) (*entities.Users, error) {
	var result models.Users
	query := fmt.Sprintf("SELECT UUID,USERSTATUS,ACCESSTOKEN FROM USERS WHERE ACCESSTOKEN='%s'", accessToken)
	rows, err := r.Conn.Query(query)

	if err != nil {
		return nil, errors.InternalError{Message: err.Error()}
	}
	defer rows.Close()

	for rows.Next() {

		err := rows.Scan(&result.Uuid, &result.UserStatus, &result.AccessToken)
		if err != nil {
			return nil, errors.InternalError{Message: err.Error()}
		}
	}

	output := result.ToEntity()

	return &output, nil
}
