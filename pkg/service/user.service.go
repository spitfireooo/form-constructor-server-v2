package service

import (
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

	query := fmt.Sprintf(`
		SELECT id, email, phone, address, nickname, logo, created_at, updated_at
		FROM %s WHERE id=$1`, database.UsersTable,
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
