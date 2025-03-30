package service

import (
	"errors"
	"fmt"
	"github.com/spitfireooo/form-constructor-server-v2/pkg/database"
	"github.com/spitfireooo/form-constructor-server-v2/pkg/model/request"
	"github.com/spitfireooo/form-constructor-server-v2/pkg/model/response"
)

func CreateFieldMultiply(multiply request.FieldMultiply, fieldId int) (response.FieldMultiply, error) {
	if _, err := GetOneField(fieldId); err != nil {
		return response.FieldMultiply{}, errors.New("field not exist")
	}

	if _, err := GetFieldMultiply(fieldId); err == nil {
		return response.FieldMultiply{}, errors.New("multiply exist now")
	}

	res := new(response.FieldMultiply)

	query := fmt.Sprintf(
		`INSERT INTO %s (field_id, is_multiply) VALUES ($1, $2) RETURNING *`,
		database.FieldMultiplyTable,
	)
	err := database.Connect.
		QueryRowx(query, fieldId, multiply.IsMultiply).
		Scan(&res.ID, &res.FieldId, &res.IsMultiply)

	return *res, err
}

func GetFieldMultiply(fieldId int) (response.FieldMultiply, error) {
	res := new(response.FieldMultiply)

	query := fmt.Sprintf(`SELECT * FROM %s WHERE field_id = $1`, database.FieldMultiplyTable)
	err := database.Connect.Get(res, query, fieldId)

	return *res, err
}

func UpdateFieldMultiply(multiply request.FieldMultiply, fieldId int) (response.FieldMultiply, error) {
	if _, err := GetFieldMultiply(fieldId); err != nil {
		return response.FieldMultiply{}, errors.New("multiply not exist")
	}

	res := new(response.FieldMultiply)

	query := fmt.Sprintf(
		`UPDATE %s SET is_multiply = $1 WHERE field_id = $2 RETURNING *`,
		database.FieldMultiplyTable,
	)
	err := database.Connect.
		QueryRowx(query, multiply.IsMultiply, fieldId).
		Scan(&res.ID, &res.FieldId, &res.IsMultiply)

	return *res, err
}

func DeleteFieldMultiply(fieldId int) error {
	if _, err := GetOneField(fieldId); err != nil {
		return errors.New("field not exist")
	}

	query := fmt.Sprintf(`DELETE FROM %s WHERE field_id = $1`, database.FieldMultiplyTable)
	_, err := database.Connect.Exec(query, fieldId)

	return err
}
