package service

import (
	"errors"
	"fmt"
	"github.com/spitfireooo/form-constructor-server-v2/pkg/database"
	"github.com/spitfireooo/form-constructor-server-v2/pkg/model/request"
	"github.com/spitfireooo/form-constructor-server-v2/pkg/model/response"
)

func CreateFieldVariant(variant request.FieldVariants, fieldId int) (response.FieldVariants, error) {
	if _, err := GetOneField(fieldId); err != nil {
		return response.FieldVariants{}, errors.New("field not exist")
	}

	res := new(response.FieldVariants)

	query := fmt.Sprintf(
		`INSERT INTO %s (field_id, variant, name) VALUES ($1, $2, $3) RETURNING *`,
		database.FieldVariantsTable,
	)
	err := database.Connect.
		QueryRowx(query, fieldId, variant.Variant, variant.Name).
		Scan(&res.ID, &res.FieldId, &res.Variant, &res.Name)

	return *res, err
}

func GetFieldVariants(fieldId int) ([]response.FieldVariants, error) {
	res := new([]response.FieldVariants)

	query := fmt.Sprintf(`SELECT * FROM %s WHERE field_id = $1`, database.FieldVariantsTable)
	err := database.Connect.Select(res, query, fieldId)

	return *res, err
}

func GetAllFieldVariants() ([]response.FieldVariants, error) {
	res := new([]response.FieldVariants)

	query := fmt.Sprintf(`SELECT * FROM %s`, database.FieldVariantsTable)
	err := database.Connect.Select(res, query)

	return *res, err
}

func GetOneFieldVariants(id int) (response.FieldVariants, error) {
	res := new(response.FieldVariants)

	query := fmt.Sprintf(`SELECT * FROM %s WHERE id = $1`, database.FieldVariantsTable)
	err := database.Connect.Get(res, query, id)

	return *res, err
}

func GetOneFieldVariantsByNameAndFieldID(name string, id int) (response.FieldVariants, error) {
	res := new(response.FieldVariants)

	query := fmt.Sprintf(`SELECT * FROM %s WHERE name = $1 AND field_id = $2`, database.FieldVariantsTable)
	err := database.Connect.Get(res, query, name, id)

	return *res, err
}

func UpdateFieldVariants(variant request.FieldVariantsUpdate, id int) (response.FieldVariants, error) {
	if _, err := GetOneFieldVariants(id); err != nil {
		return response.FieldVariants{}, errors.New("variant not exist")
	}

	res := new(response.FieldVariants)

	query := fmt.Sprintf(
		`UPDATE %s SET 
			variant = COALESCE($1, variant), 
            name= COALESCE($2, name) 
        WHERE id = $3 RETURNING *`,
		database.FieldVariantsTable,
	)
	err := database.Connect.
		QueryRowx(query, variant.Variant, variant.Name, id).
		Scan(&res.ID, &res.FieldId, &res.Variant, &res.Name)

	return *res, err
}

func DeleteOneFieldVariant(id int) error {
	if _, err := GetOneFieldVariants(id); err != nil {
		return errors.New("variant not exist")
	}

	query := fmt.Sprintf(`DELETE FROM %s WHERE id = $1`, database.FieldVariantsTable)
	_, err := database.Connect.Exec(query, id)

	return err
}

func DeleteFieldVariants(fieldId int) error {
	if res, err := GetFieldVariants(fieldId); err != nil && res != nil {
		return errors.New("variants not exist")
	}

	query := fmt.Sprintf(`DELETE FROM %s WHERE field_id = $1`, database.FieldVariantsTable)
	_, err := database.Connect.Exec(query, fieldId)

	return err
}
