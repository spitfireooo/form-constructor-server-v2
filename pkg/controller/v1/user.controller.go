package controller_v1

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/spitfireooo/form-constructor-server-v2/pkg/model/request"
	"github.com/spitfireooo/form-constructor-server-v2/pkg/service"
	"github.com/spitfireooo/form-constructor-server-v2/pkg/utils"
	"log"
	"net/http"
)

// @Summary	GetAllUsers
// @Tags User
// @Description Get All Users
// @ID get-all-users
// @Accept json
// @Produce	json
// @Success 200 {array} response.User
// @Router /api/v1/user [get]
func GetAllUsers(ctx *fiber.Ctx) error {
	if res, err := service.GetAllUsers(); err != nil {
		log.Println("Error in user service", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error in user service",
		})
	} else {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"data": res,
		})
	}
}

// @Summary	GetOneUser
// @Tags User
// @Description Get One User
// @ID get-one-user
// @Accept json
// @Produce	json
// @Success 200 {object} response.User
// @Router /api/v1/user/:id [get]
func GetOneUser(ctx *fiber.Ctx) error {
	id, _ := ctx.ParamsInt("id")

	if res, err := service.GetOneUser(id); err != nil {
		log.Println("Error in user service", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error in user service",
		})
	} else {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"data": res,
		})
	}
}

// @Summary	UpdateUser
// @Tags User
// @Description Update User
// @ID update-user
// @Accept json
// @Produce	json
// @Param input	body request.User true "body info"
// @Success 200 {object} response.User
// @Router /api/v1/user/:id [patch]
func UpdateUser(ctx *fiber.Ctx) error {
	body := new(request.User)
	id := ctx.Params("id")

	if err := ctx.BodyParser(body); err != nil {
		log.Println("Error in parsing request", err)
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Error in parsing request",
		})
	}

	if err := validator.New().Struct(body); err != nil {
		log.Println("Validation errors", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Validation errors",
		})
	}

	if file, err := ctx.FormFile("logo"); err != nil {
		log.Println("Error in file upload", err)
	} else if file.Size > 0 {
		fmt.Println("file", file.Header.Get("Content-Type"))
		if err = utils.CheckContentType(
			file.Header.Get("Content-Type"),
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

		filename := fmt.Sprintf("./static/uploads/%s_%s", uuid.New(), file.Filename)
		if err = ctx.SaveFile(file, filename); err != nil {
			log.Println("Error in save file", err)
		}
		body.Logo = filename
	}

	if user, err := service.UpdateUser(*body, id); err != nil {
		log.Println("Error in user service", err)
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error in user service",
		})
	} else {
		return ctx.Status(http.StatusOK).JSON(fiber.Map{
			"data": user,
		})
	}
}

// @Summary DeleteUser
// @Tags User
// @Description Delete User
// @ID delete-user
// @Accept json
// @Produce	json
// @Success 200 {string} string
// @Router /api/v1/user/:id [delete]
func DeleteUser(ctx *fiber.Ctx) error {
	id, _ := ctx.ParamsInt("id")

	if err := service.DeleteUser(id); err != nil {
		log.Println("Error in user service", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error in user service",
		})
	} else {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "User deleted!",
		})
	}
}
