package service

import (
	"errors"
	"fmt"
	"github.com/spitfireooo/form-constructor-server-v2/pkg/database"
	"github.com/spitfireooo/form-constructor-server-v2/pkg/model/request"
	"github.com/spitfireooo/form-constructor-server-v2/pkg/model/response"
)

func CreateFieldPlaceholder(placeholder request.FieldPlaceholder, fieldId int) (response.FieldPlaceholder, error) {
	if _, err := GetOneField(fieldId); err != nil {
		return response.FieldPlaceholder{}, errors.New("field not exist")
	}

	if _, err := GetFieldPlaceholder(fieldId); err == nil {
		return response.FieldPlaceholder{}, errors.New("placeholder exist now")
	}

	res := new(response.FieldPlaceholder)

	query := fmt.Sprintf(
		`INSERT INTO %s (field_id, placeholder) VALUES ($1, $2) RETURNING *`,
		database.FieldPlaceholderTable,
	)
	err := database.Connect.
		QueryRowx(query, fieldId, placeholder.Placeholder).
		Scan(&res.ID, &res.FieldId, &res.Placeholder)

	return *res, err
}

func GetFieldPlaceholder(fieldId int) (response.FieldPlaceholder, error) {
	res := new(response.FieldPlaceholder)

	query := fmt.Sprintf(`SELECT * FROM %s WHERE field_id = $1`, database.FieldPlaceholderTable)
	err := database.Connect.Get(res, query, fieldId)

	return *res, err
}

func UpdateFieldPlaceholder(placeholder request.FieldPlaceholder, fieldId int) (response.FieldPlaceholder, error) {
	if _, err := GetFieldPlaceholder(fieldId); err != nil {
		return response.FieldPlaceholder{}, errors.New("placeholder not exist")
	}

	res := new(response.FieldPlaceholder)

	query := fmt.Sprintf(
		`UPDATE %s SET placeholder = $1 WHERE field_id = $2 RETURNING *`,
		database.FieldPlaceholderTable,
	)
	err := database.Connect.
		QueryRowx(query, placeholder.Placeholder, fieldId).
		Scan(&res.ID, &res.FieldId, &res.Placeholder)

	return *res, err
}

func DeleteFieldPlaceholder(fieldId int) error {
	if _, err := GetOneField(fieldId); err != nil {
		return errors.New("field not exist")
	}

	query := fmt.Sprintf(`DELETE FROM %s WHERE field_id = $1`, database.FieldPlaceholderTable)
	_, err := database.Connect.Exec(query, fieldId)

	return err
}
