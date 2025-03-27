package service

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/spitfireooo/form-constructor-server-v2/pkg/database"
	"github.com/spitfireooo/form-constructor-server-v2/pkg/model/entity"
	"github.com/spitfireooo/form-constructor-server-v2/pkg/model/request"
	"github.com/spitfireooo/form-constructor-server-v2/pkg/model/response"
	"github.com/spitfireooo/form-constructor-server-v2/pkg/utils"
	"io"
	"log"
	"net/http"
	"os"
)

func GetAllUsers() ([]response.User, error) {
	res := new([]response.User)

	query := fmt.Sprintf(`
		SELECT id, email, phone, address, nickname, logo, created_at, updated_at
		FROM %s`, database.UsersTable,
	)
	err := database.Connect.Select(res, query)

	return *res, err
}

func GetOneUser(id int) (response.User, error) {
	res := new(response.User)

	fmt.Println(id)
	query := fmt.Sprintf(`
		SELECT id, email, phone, address, nickname, logo, created_at, updated_at
		FROM %s WHERE id = $1`, database.UsersTable,
	)
	err := database.Connect.Get(res, query, id)

	return *res, err
}

func UpdateUser(user request.User, id string) (response.User, error) {
	userExist := new(entity.User)
	query := fmt.Sprintf(`SELECT * FROM %s WHERE id = $1`, database.UsersTable)
	if err := database.Connect.Get(userExist, query, id); err != nil {
		return response.User{}, err
	}

	if user.Logo != "" {
		old_path := userExist.Logo
		if err := os.Remove(old_path); err != nil {
			log.Println("Error in deleting image")
		}
	}

	res := new(response.User)

	query = fmt.Sprintf(
		`UPDATE %s SET
    		email = COALESCE($1, email),
    		phone = COALESCE($2, phone),
    		address = COALESCE($3, address),
    		password = COALESCE($4, password),
    		nickname = COALESCE($5, nickname),
    		logo = COALESCE($6, logo),
          	updated_at = NOW()
	  	WHERE id = $7
	  	RETURNING id, email, phone, address, nickname, logo, created_at, updated_at`,
		database.UsersTable,
	)
	if err := database.Connect.QueryRowx(
		query,
		user.Email,
		user.Phone,
		user.Address,
		user.Password,
		user.Nickname,
		user.Logo,
		id,
	).Scan(&res.ID, &res.Email, &res.Phone, &res.Address, &res.Nickname, &res.Logo, &res.CreatedAt, &res.UpdatedAt); err != nil {
		return response.User{}, err
	}

	return *res, nil
}

func DeleteUser(id int) error {
	userExist := new(entity.User)
	query := fmt.Sprintf(`SELECT * FROM %s WHERE id = $1`, database.UsersTable)
	if err := database.Connect.Get(userExist, query, id); err != nil {
		return err
	}

	query = fmt.Sprintf(`DELETE FROM %s WHERE id = $1`, database.UsersTable)
	if _, err := database.Connect.Exec(query, id); err != nil {
		return err
	}

	if err := os.Remove(userExist.Logo); err != nil {
		return err
	}

	return nil
}

func FileUpload(ctx *fiber.Ctx) error {
	var filename string

	if fileHeader, err := ctx.FormFile("logo"); err != nil {
		log.Println("Error in file upload", err)
	} else if fileHeader.Size > 0 {
		file, err := fileHeader.Open()

		buffer := make([]byte, 512)
		n, err := file.Read(buffer)
		if err != nil && err != io.EOF {
			log.Println("Error in file reading", err)
		}

		_, err = file.Seek(0, io.SeekStart)
		if err != nil {
			log.Println("Error in file upload", err)
		}

		contentTypeFromFile := http.DetectContentType(buffer[:n])

		if err = utils.CheckContentType(
			contentTypeFromFile,
			"image/jpg",
			"image/png",
			"image/gif",
			"image/jpeg",
		); err != nil {
			log.Println("Bad ext of file", err)
			return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
				"message": "Bad ext of file",
			})
		}

		filename = utils.GenerateFilename(fileHeader.Filename)
		if err = ctx.SaveFile(fileHeader, filename); err != nil {
			log.Println("Error in save file", err)
		}
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"filename": filename,
	})
}

func CreateUserTag(tag request.UserTag, userId int) (response.UserTag, error) {
	tagExist := new(entity.UserTag)
	query := fmt.Sprintf(`SELECT * FROM %s WHERE user_id = $1 AND tag_id =$2`, database.UserTagsTable)
	if err := database.Connect.Get(tagExist, query, userId, tag.TagID); err == nil {
		return response.UserTag{}, errors.New("user exist")
	}

	res := new(response.UserTag)

	query = fmt.Sprintf(
		`INSERT INTO %s (user_id, tag_id) VALUES ($1, $2) RETURNING *`,
		database.UserTagsTable,
	)
	err := database.Connect.
		QueryRowx(query, userId, tag.TagID).
		Scan(&res.ID, &res.UserID, &res.TagID)

	return *res, err
}

func GetUserTags(userId int) ([]response.UserTag, error) {
	res := new([]response.UserTag)

	query := fmt.Sprintf(
		`SELECT * FROM %s WHERE user_id = $1`,
		database.UserTagsTable,
	)
	err := database.Connect.Select(res, query, userId)

	return *res, err
}

func GetUserTag(userId, id int) (response.UserTag, error) {
	res := new(response.UserTag)

	query := fmt.Sprintf(
		`SELECT * FROM %s WHERE user_id = $1 AND tag_id = $2`,
		database.UserTagsTable,
	)
	err := database.Connect.Get(res, query, userId, id)

	return *res, err
}

func UpdateUserTag(tag request.UserTagUpdate, userId, id int) (response.UserTag, error) {
	userTagExist := new(entity.UserTag)
	query := fmt.Sprintf(
		`SELECT * FROM %s WHERE user_id = $1 AND tag_id = $2`,
		database.UserTagsTable,
	)
	if err := database.Connect.Get(userTagExist, query, userId, id); err != nil {
		return response.UserTag{}, err
	}

	if _, err := GetOneUser(userId); err != nil {
		return response.UserTag{}, err
	}
	if _, err := GetOneTag(id); err != nil {
		return response.UserTag{}, err
	}

	res := new(response.UserTag)

	query = fmt.Sprintf(
		`UPDATE %s SET 
	    	user_id = COALESCE($1, user_id), 
            tag_id = COALESCE($2, tag_id)
	  	WHERE user_id = $3 AND tag_id = $4 RETURNING *`,
		database.UserTagsTable,
	)
	if err := database.Connect.
		QueryRowx(query, tag.UserId, tag.TagID, userId, id).
		Scan(&res.ID, &res.UserID, &res.TagID); err != nil {
		return response.UserTag{}, err
	}

	return *res, nil
}

func DeleteUserTags(userId int) error {
	tagExist := new(entity.UserTag)
	query := fmt.Sprintf(`SELECT * FROM %s WHERE user_id = $1`, database.UserTagsTable)
	if err := database.Connect.Get(tagExist, query, userId); err != nil {
		return err
	}

	query = fmt.Sprintf(`DELETE FROM %s WHERE user_id = $1`, database.UserTagsTable)
	_, err := database.Connect.Exec(query, userId)

	return err
}

func DeleteUserTag(userId, id int) error {
	tagExist := new(entity.UserTag)
	query := fmt.Sprintf(`SELECT * FROM %s WHERE user_id = $1 AND tag_id = $2`, database.UserTagsTable)
	if err := database.Connect.Get(tagExist, query, userId, id); err != nil {
		return err
	}

	query = fmt.Sprintf(`DELETE FROM %s WHERE user_id = $1 AND tag_id = $2`, database.UserTagsTable)
	_, err := database.Connect.Exec(query, userId, id)

	return err
}
