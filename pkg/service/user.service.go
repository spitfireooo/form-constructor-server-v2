package service

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/spitfireooo/form-constructor-server-v2/pkg/database"
	"github.com/spitfireooo/form-constructor-server-v2/pkg/model/response"
	"github.com/spitfireooo/form-constructor-server-v2/pkg/utils"
	"io"
	"log"
	"net/http"
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

func GetOneUser(id string) (response.User, error) {
	res := new(response.User)

	query := fmt.Sprintf(`
		SELECT id, email, phone, address, nickname, logo, created_at, updated_at
		FROM %s`, database.UsersTable,
	)
	err := database.Connect.Get(res, query)

	return *res, err
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
		}

		_, err = file.Seek(0, io.SeekStart)
		if err != nil {
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
