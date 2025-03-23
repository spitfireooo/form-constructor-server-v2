package controller_v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/spitfireooo/form-constructor-server-v2/pkg/service"
	"log"
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
	id := ctx.Params("id")

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
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "GetAllUsers",
	})
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
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "GetAllUsers",
	})
}
