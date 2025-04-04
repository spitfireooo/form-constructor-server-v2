package service

import (
	"errors"
	"fmt"
	"github.com/spitfireooo/form-constructor-server-v2/pkg/database"
	"github.com/spitfireooo/form-constructor-server-v2/pkg/model/request"
	"github.com/spitfireooo/form-constructor-server-v2/pkg/model/response"
)

func CreateField(field request.Field, formId int) (response.Field, error) {
	if _, err := GetOneForm(formId); err != nil {
		return response.Field{}, errors.New("form not exist")
	}

	res := new(response.Field)

	query := fmt.Sprintf(`
		INSERT INTO %s (
    		form_id, type, label, name, order_of, required	            
		) VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING *`,
		database.FieldTable,
	)
	err := database.Connect.
		QueryRowx(query, formId, field.Type, field.Label, field.Name, field.OrderOf, field.Required).
		Scan(&res.ID, &res.FormID, &res.Type, &res.Label, &res.Name, &res.OrderOf, &res.Required)

	return *res, err
}

func GetAllFormFields(formId int) ([]response.Field, error) {
	res := new([]response.Field)

	query := fmt.Sprintf(`SELECT * FROM %s WHERE form_id = $1`, database.FieldTable)
	err := database.Connect.Select(res, query, formId)

	return *res, err
}

func GetAllFields() ([]response.Field, error) {
	res := new([]response.Field)

	query := fmt.Sprintf(`SELECT * FROM %s`, database.FieldTable)
	err := database.Connect.Select(res, query)

	return *res, err
}

func GetOneField(id int) (response.Field, error) {
	res := new(response.Field)

	query := fmt.Sprintf(`SELECT * FROM %s WHERE id = $1`, database.FieldTable)
	err := database.Connect.Get(res, query, id)

	return *res, err
}

func GetFieldByNameAndFormID(name string, id int) (response.Field, error) {
	res := new(response.Field)

	query := fmt.Sprintf(`SELECT * FROM %s WHERE name = $1 form_id = $2`, database.FieldTable)
	err := database.Connect.Get(res, query, name, id)

	return *res, err
}

func UpdateField(field request.FieldUpdate, id int) (response.Field, error) {
	if field.FormID != nil {
		if _, err := GetOneForm(int(*field.FormID)); err != nil {
			return response.Field{}, errors.New("form not exist")
		}
	}

	res := new(response.Field)

	query := fmt.Sprintf(
		`UPDATE %s SET
    		form_id = COALESCE($1, form_id),
    		type = COALESCE($2, type),
    		label = COALESCE($3, label),
    		name = COALESCE($4, name),
    		order_of = COALESCE($5, order_of),
    		required = COALESCE($6, required)
	  	WHERE id = $7
	  	RETURNING *`,
		database.FieldTable,
	)
	if err := database.Connect.QueryRowx(
		query,
		field.FormID,
		field.Type,
		field.Label,
		field.Name,
		field.OrderOf,
		field.Required,
		id,
	).Scan(&res.ID, &res.FormID, &res.Type, &res.Label, &res.Name, &res.OrderOf, &res.Required); err != nil {
		return response.Field{}, err
	}

	return *res, nil
}

func DeleteField(id int) error {
	if _, err := GetOneField(id); err != nil {
		return errors.New("field not exist")
	}

	query := fmt.Sprintf(`DELETE FROM %s WHERE id = $1`, database.FieldTable)
	_, err := database.Connect.Exec(query, id)

	return err
}

func DeleteFormFields(formId int) error {
	if res, err := GetAllFormFields(formId); err != nil || res == nil {
		return errors.New("field not exist")
	}

	query := fmt.Sprintf(`DELETE FROM %s WHERE form_id = $1`, database.FieldTable)
	_, err := database.Connect.Exec(query, formId)

	return err
}
