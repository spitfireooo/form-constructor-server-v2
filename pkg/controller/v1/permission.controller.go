package controller_v1

import (
	"github.com/gofiber/fiber/v2"
)

// @Summary	CreatePermission
// @Tags Permission
// @Description Create Permission
// @ID create-permission
// @Accept json
// @Produce	json
// @Param input	body request.UserPermission true "body info"
// @Success 200 {object} response.UserPermission
// @Router /api/v1/permission [post]
func CreatePermission(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "CreatePermission",
	})
}

// @Summary	GetAllPermissions
// @Tags Permission
// @Description Get All Permissions
// @ID get-all-permissions
// @Accept json
// @Produce	json
// @Success 200 {array} response.UserPermission
// @Router /api/v1/permission [get]
func GetAllPermissions(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "GetAllPermissions",
	})
}

// @Summary	GetOnePermission
// @Tags Permission
// @Description Get One Permission
// @ID get-one-permission
// @Accept json
// @Produce	json
// @Success 200 {object} response.UserPermission
// @Router /api/v1/permission/:id [get]
func GetOnePermission(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "GetOnePermission",
	})
}

// @Summary	UpdatePermission
// @Tags Permission
// @Description Update Permission
// @ID update-permission
// @Accept json
// @Produce	json
// @Param input	body request.UserPermission true "body info"
// @Success 200 {object} response.UserPermission
// @Router /api/v1/permission/:id [patch]
func UpdatePermission(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "UpdatePermission",
	})
}

// @Summary DeletePermission
// @Tags Permission
// @Description Delete Permission
// @ID delete-permission
// @Accept json
// @Produce	json
// @Success 200 {string} string
// @Router /api/v1/permission/:id [delete]
func DeletePermission(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "DeletePermission",
	})
}
