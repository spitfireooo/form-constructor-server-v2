package service

import (
	"fmt"
	"github.com/spitfireooo/form-constructor-server-v2/pkg/database"
	"github.com/spitfireooo/form-constructor-server-v2/pkg/model/response"
)

func GetAllUsers() ([]response.User, error) {
	res := new([]response.User)

	query := fmt.Sprintf(`SELECT id, email, phone, address, nickname, logo, created_at, updated_at FROM %s`, database.UsersTable)
	err := database.Connect.Select(res, query)

	return *res, err
}

func GetOneUser(id string) (response.User, error) {
	res := new(response.User)

	query := fmt.Sprintf(`SELECT id, email, phone, address, nickname, logo, created_at, updated_at FROM %s`, database.UsersTable)
	err := database.Connect.Get(res, query)

	return *res, err
}
