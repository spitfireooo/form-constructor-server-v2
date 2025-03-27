package service

import (
	"fmt"
	"github.com/spitfireooo/form-constructor-server-v2/pkg/database"
	"github.com/spitfireooo/form-constructor-server-v2/pkg/model/entity"
	"github.com/spitfireooo/form-constructor-server-v2/pkg/model/request"
	"github.com/spitfireooo/form-constructor-server-v2/pkg/model/response"
)

func CreateTag(tag request.Tag) (response.Tag, error) {
	res := new(response.Tag)

	query := fmt.Sprintf(
		`INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING *`,
		database.TagsTable,
	)
	err := database.Connect.
		QueryRowx(query, tag.Title, tag.Description).
		Scan(&res.ID, &res.Title, &res.Description)

	return *res, err
}

func GetAllUserTags() ([]response.UserTag, error) {
	res := new([]response.UserTag)

	query := fmt.Sprintf(`SELECT * FROM %s`, database.UserTagsTable)
	err := database.Connect.Select(res, query)

	return *res, err
}

func GetAllTags() ([]response.Tag, error) {
	res := new([]response.Tag)

	query := fmt.Sprintf(`SELECT * FROM %s`, database.TagsTable)
	err := database.Connect.Select(res, query)

	return *res, err
}

func GetOneTag(id int) (response.Tag, error) {
	res := new(response.Tag)

	query := fmt.Sprintf(`SELECT * FROM %s WHERE id = $1`, database.TagsTable)
	err := database.Connect.Get(res, query, id)

	return *res, err
}

func UpdateTag(tag request.TagUpdate, id int) (response.Tag, error) {
	tagExist := new(entity.Tag)
	query := fmt.Sprintf(`SELECT * FROM %s WHERE id = $1`, database.TagsTable)
	if err := database.Connect.Get(tagExist, query, id); err != nil {
		return response.Tag{}, err
	}

	res := new(response.Tag)

	query = fmt.Sprintf(
		`UPDATE %s SET 
	    	title = COALESCE($1, title), 
            description = COALESCE($2, description)
	  	WHERE id = $3 RETURNING *`,
		database.TagsTable,
	)
	if err := database.Connect.
		QueryRowx(query, tag.Title, tag.Description, id).
		Scan(&res.ID, &res.Title, &res.Description); err != nil {
		return response.Tag{}, err
	}

	return *res, nil
}

func DeleteTag(id int) error {
	tagExist := new(entity.Tag)
	query := fmt.Sprintf(`SELECT * FROM %s WHERE id = $1`, database.TagsTable)
	if err := database.Connect.Get(tagExist, query, id); err != nil {
		return err
	}

	query = fmt.Sprintf(`DELETE FROM %s WHERE id = $1`, database.TagsTable)
	_, err := database.Connect.Exec(query, id)

	return err
}
