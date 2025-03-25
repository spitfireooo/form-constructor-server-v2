package service

import (
	"fmt"
	"github.com/spitfireooo/form-constructor-server-v2/pkg/database"
	"github.com/spitfireooo/form-constructor-server-v2/pkg/model/entity"
	"github.com/spitfireooo/form-constructor-server-v2/pkg/model/request"
	"github.com/spitfireooo/form-constructor-server-v2/pkg/model/response"
)

func CreatePermission(permission request.UserPermission, id int) (response.UserPermission, error) {
	res := new(response.UserPermission)

	query := fmt.Sprintf(
		`INSERT INTO %s (user_id, permission) VALUES ($1, $2) RETURNING *`,
		database.UserPermissionsTable,
	)
	err := database.Connect.
		QueryRowx(query, id, permission.Permission).
		Scan(&res.ID, &res.UserID, &res.Permission)

	return *res, err
}

func GetAllPermissions() ([]response.UserPermission, error) {
	res := new([]response.UserPermission)

	query := fmt.Sprintf(`SELECT * FROM %s`, database.UserPermissionsTable)
	err := database.Connect.Select(res, query)

	return *res, err
}

func GetUserPermissions(id int) ([]response.UserPermission, error) {
	res := new([]response.UserPermission)

	query := fmt.Sprintf(`SELECT * FROM %s WHERE user_id = $1`, database.UserPermissionsTable)
	err := database.Connect.Select(res, query, id)

	return *res, err
}

func GetOnePermission(id int) (response.UserPermission, error) {
	res := new(response.UserPermission)

	query := fmt.Sprintf(`SELECT * FROM %s WHERE id = $1`, database.UserPermissionsTable)
	err := database.Connect.Get(res, query, id)

	return *res, err
}

func UpdatePermission(permission request.UserPermission, id int) (response.UserPermission, error) {
	permissionExist := new(entity.UserPermission)
	query := fmt.Sprintf(`SELECT * FROM %s WHERE id = $1`, database.UserPermissionsTable)
	if err := database.Connect.Get(permissionExist, query, id); err != nil {
		return response.UserPermission{}, err
	}

	res := new(response.UserPermission)

	query = fmt.Sprintf(
		`UPDATE %s SET permission = $1 WHERE id = $2 RETURNING *`,
		database.UserPermissionsTable,
	)
	if err := database.Connect.QueryRowx(query, permission.Permission, id).
		Scan(&res.ID, &res.UserID, &res.Permission); err != nil {
		return response.UserPermission{}, err
	}

	return *res, nil
}

func DeletePermission(id int) error {
	permissionExist := new(entity.UserPermission)
	query := fmt.Sprintf(`SELECT * FROM %s WHERE id = $1`, database.UserPermissionsTable)
	if err := database.Connect.Get(permissionExist, query, id); err != nil {
		return err
	}

	query = fmt.Sprintf(`DELETE FROM %s WHERE id = $1`, database.UserPermissionsTable)
	if _, err := database.Connect.Exec(query, id); err != nil {
		return err
	}

	return nil
}

func DeleteUserPermissions(id int) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE user_id = $1`, database.UserPermissionsTable)
	if _, err := database.Connect.Exec(query, id); err != nil {
		return err
	}

	return nil
}
