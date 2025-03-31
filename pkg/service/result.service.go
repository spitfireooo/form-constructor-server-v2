package service

import (
	"errors"
	"fmt"
	"github.com/spitfireooo/form-constructor-server-v2/pkg/database"
	"github.com/spitfireooo/form-constructor-server-v2/pkg/model/request"
	"github.com/spitfireooo/form-constructor-server-v2/pkg/model/response"
)

func CreateResult(result request.Result) (response.Result, error) {
	field, err := GetOneField(int(result.FieldID))
	if err != nil {
		return response.Result{}, errors.New("field not exist")
	} else {
		if _, err = GetOneForm(int(field.FormID)); err != nil {
			return response.Result{}, errors.New("form not exist")
		}
	}

	if _, err = GetOneUser(int(result.UserID)); err != nil {
		return response.Result{}, errors.New("user not exist")
	}

	res := new(response.Result)

	query := fmt.Sprintf(
		`INSERT INTO %s (
            field_id, value, type, form_id, user_id
		) VALUES ($1, $2, $3, $4, $5) RETURNING *`,
		database.ResultsTable,
	)
	err = database.Connect.
		QueryRowx(query, result.FieldID, result.Value, field.Type, field.FormID, result.UserID).
		Scan(&res.ID, &res.FieldID, &res.Value, &res.Type, &res.FormID, &res.UserID)

	return *res, err
}

func GetAllResults() ([]response.Result, error) {
	res := new([]response.Result)

	query := fmt.Sprintf(`SELECT * FROM %s`, database.ResultsTable)
	err := database.Connect.Select(res, query)

	return *res, err
}

func GetOneResult(id int) (response.Result, error) {
	res := new(response.Result)

	query := fmt.Sprintf(`SELECT * FROM %s WHERE id = $1`, database.ResultsTable)
	err := database.Connect.Get(res, query, id)

	return *res, err
}

func GetFormResults(id int) ([]response.Result, error) {
	res := new([]response.Result)

	query := fmt.Sprintf(`SELECT * FROM %s WHERE form_id = $1`, database.ResultsTable)
	err := database.Connect.Select(res, query, id)

	return *res, err
}

func GetFieldResults(id int) ([]response.Result, error) {
	res := new([]response.Result)

	query := fmt.Sprintf(`SELECT * FROM %s WHERE field_id = $1`, database.ResultsTable)
	err := database.Connect.Select(res, query, id)

	return *res, err
}

func UpdateResult(result request.ResultUpdate, id int) (response.Result, error) {
	field := new(request.FieldUpdate)

	if _, err := GetOneResult(id); err != nil {
		return response.Result{}, errors.New("result not exist")
	}

	if result.FieldID != nil {
		if res, err := GetOneField(int(*result.FieldID)); err != nil {
			return response.Result{}, errors.New("field not exist")
		} else {
			field.Type = &res.Type
		}
	}

	res := new(response.Result)
	query := fmt.Sprintf(
		`UPDATE %s SET
			field_id = COALESCE($1, field_id),
			value = COALESCE($2, value),
			type = COALESCE($3, type)
       	WHERE id = $4 RETURNING *`,
		database.ResultsTable,
	)
	err := database.Connect.
		QueryRowx(query, result.FieldID, result.Value, field.Type, id).
		Scan(&res.ID, &res.FieldID, &res.Value, &res.Type, &res.FormID, &res.UserID)

	return *res, err
}

func DeleteResult(id int) error {
	if _, err := GetOneResult(id); err != nil {
		return errors.New("result not exist")
	}

	query := fmt.Sprintf(`DELETE FROM %s WHERE id = $1`, database.ResultsTable)
	_, err := database.Connect.Exec(query, id)

	return err
}

func DeleteFormResults(formId int) error {
	if res, err := GetFormResults(formId); err != nil && res == nil {
		return errors.New("result not exist")
	}

	query := fmt.Sprintf(`DELETE FROM %s WHERE form_id = $1`, database.ResultsTable)
	_, err := database.Connect.Exec(query, formId)

	return err
}

func DeleteFieldResults(fieldId int) error {
	if res, err := GetFieldResults(fieldId); err != nil && res == nil {
		return errors.New("result not exist")
	}

	query := fmt.Sprintf(`DELETE FROM %s WHERE field_id = $1`, database.ResultsTable)
	_, err := database.Connect.Exec(query, fieldId)

	return err
}
