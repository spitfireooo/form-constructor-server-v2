package service

import (
	"errors"
	"fmt"
	"github.com/spitfireooo/form-constructor-server-v2/pkg/database"
	"github.com/spitfireooo/form-constructor-server-v2/pkg/model/request"
	"github.com/spitfireooo/form-constructor-server-v2/pkg/model/response"
)

func CreateFieldRange(ranges request.FieldRange, fieldId int) (response.FieldRange, error) {
	if _, err := GetOneField(fieldId); err != nil {
		return response.FieldRange{}, errors.New("field not exist")
	}

	if _, err := GetFieldRange(fieldId); err == nil {
		return response.FieldRange{}, errors.New("range exist now")
	}

	res := new(response.FieldRange)

	query := fmt.Sprintf(
		`INSERT INTO %s (field_id, min, max) VALUES ($1, $2, $3) RETURNING *`,
		database.FieldRangeTable,
	)
	err := database.Connect.
		QueryRowx(query, fieldId, ranges.Min, ranges.Max).
		Scan(&res.ID, &res.FieldId, &res.Min, &res.Max)

	return *res, err
}

func GetFieldRange(fieldId int) (response.FieldRange, error) {
	res := new(response.FieldRange)

	query := fmt.Sprintf(`SELECT * FROM %s WHERE field_id = $1`, database.FieldRangeTable)
	err := database.Connect.Get(res, query, fieldId)

	return *res, err
}

func UpdateFieldRange(ranges request.FieldRangeUpdate, fieldId int) (response.FieldRange, error) {
	if _, err := GetFieldRange(fieldId); err != nil {
		return response.FieldRange{}, errors.New("range not exist")
	}

	res := new(response.FieldRange)

	query := fmt.Sprintf(
		`UPDATE %s SET 
              min = COALESCE($1, min), 
              max = COALESCE($2, max) 
        WHERE field_id = $3 RETURNING *`,
		database.FieldRangeTable,
	)
	err := database.Connect.
		QueryRowx(query, ranges.Min, ranges.Max, fieldId).
		Scan(&res.ID, &res.FieldId, &res.Min, &res.Max)

	return *res, err
}

func DeleteFieldRange(fieldId int) error {
	if _, err := GetOneField(fieldId); err != nil {
		return errors.New("field not exist")
	}

	query := fmt.Sprintf(`DELETE FROM %s WHERE field_id = $1`, database.FieldRangeTable)
	_, err := database.Connect.Exec(query, fieldId)

	return err
}
