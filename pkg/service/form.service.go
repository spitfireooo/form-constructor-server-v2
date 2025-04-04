package service

import (
	"errors"
	"fmt"
	"github.com/spitfireooo/form-constructor-server-v2/pkg/database"
	"github.com/spitfireooo/form-constructor-server-v2/pkg/model/entity"
	"github.com/spitfireooo/form-constructor-server-v2/pkg/model/request"
	"github.com/spitfireooo/form-constructor-server-v2/pkg/model/response"
	"log"
	"os"
)

func CreateForm(form request.Form, userId int) (response.Form, error) {
	formExist := new(entity.Form)
	query := fmt.Sprintf(`SELECT * FROM %s WHERE slug = $1`, database.FormsTable)
	if err := database.Connect.Get(formExist, query, form.Slug); err == nil {
		return response.Form{}, errors.New("user with slug exist")
	}

	res := new(response.Form)

	query = fmt.Sprintf(`
		INSERT INTO %s (
		    title,
			slug, 
			description,
		    logo,
		    author_id
		) VALUES ($1, $2, $3, $4, $5) 
		RETURNING *`,
		database.FormsTable,
	)
	err := database.Connect.
		QueryRowx(query, form.Title, form.Slug, form.Description, form.Logo, userId).
		Scan(&res.ID, &res.Title, &res.Slug, &res.Description, &res.Logo, &res.AuthorId)

	return *res, err
}

func GetAllForms() ([]response.Form, error) {
	res := new([]response.Form)

	query := fmt.Sprintf(`SELECT * FROM %s`, database.FormsTable)
	err := database.Connect.Select(res, query)

	return *res, err
}

func GetOneForm(id int) (response.Form, error) {
	res := new(response.Form)

	query := fmt.Sprintf(`SELECT * FROM %s WHERE id = $1`, database.FormsTable)
	err := database.Connect.Get(res, query, id)

	return *res, err
}

func UpdateForm(form request.FormUpdate, id int) (response.Form, error) {
	formExist := new(entity.Form)
	query := fmt.Sprintf(`SELECT * FROM %s WHERE id = $1`, database.FormsTable)
	if err := database.Connect.Get(formExist, query, id); err != nil {
		return response.Form{}, err
	}

	if form.AuthorId != nil {
		if _, err := GetOneUser(int(*form.AuthorId)); err != nil {
			return response.Form{}, err
		}
	}

	if form.Logo != nil {
		old_path := formExist.Logo
		if err := os.Remove(*old_path); err != nil {
			log.Println("Error in deleting image")
		}
	}

	res := new(response.Form)

	query = fmt.Sprintf(
		`UPDATE %s SET
    		title = COALESCE($1, title),
    		slug = COALESCE($2, slug),
    		description = COALESCE($3, description),
    		logo = COALESCE($4, logo),
    		author_id = COALESCE($5, author_id)
	  	WHERE id = $6
	  	RETURNING *`,
		database.FormsTable,
	)
	if err := database.Connect.QueryRowx(
		query,
		form.Title,
		form.Slug,
		form.Description,
		form.Logo,
		form.AuthorId,
		id,
	).Scan(&res.ID, &res.Title, &res.Slug, &res.Description, &res.Logo, &res.AuthorId); err != nil {
		return response.Form{}, err
	}

	return *res, nil
}

func DeleteForm(id int) error {
	formExist := new(entity.Form)
	query := fmt.Sprintf(`SELECT * FROM %s WHERE id = $1`, database.FormsTable)
	if err := database.Connect.Get(formExist, query, id); err != nil {
		return err
	}

	query = fmt.Sprintf(`DELETE FROM %s WHERE id = $1`, database.FormsTable)
	_, err := database.Connect.Exec(query, id)

	return err
}
