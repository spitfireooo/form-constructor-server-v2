package controller_v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/spitfireooo/form-constructor-server-v2/pkg/model/request"
	"github.com/spitfireooo/form-constructor-server-v2/pkg/service"
	"log"
	"net/http"
)

// @Summary	CreatePermission
// @Tags Permission
// @Description Create Permission
// @ID create-permission
// @Accept json
// @Produce	json
// @Param input	body request.UserPermission true "body info"
// @Success 200 {object} response.UserPermission
// @Router /api/v1/user/permission/:userId [post]
func CreatePermission(ctx *fiber.Ctx) error {
	body := new(request.UserPermission)
	id, _ := ctx.ParamsInt("userId")

	if err := ctx.BodyParser(body); err != nil {
		log.Println("Error in parsing request", err)
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Error in parsing request",
		})
	}

	if permission, err := service.CreatePermission(*body, id); err != nil {
		log.Println("Error in permission service", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error in user service",
		})
	} else {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"data": permission,
		})
	}
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
	if permissions, err := service.GetAllPermissions(); err != nil {
		log.Println("Error in permission service", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error in user service",
		})
	} else {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"data": permissions,
		})
	}
}

// @Summary	GetAllPermissions
// @Tags Permission
// @Description Get All Permissions
// @ID get-all-permissions
// @Accept json
// @Produce	json
// @Success 200 {array} response.UserPermission
// @Router /api/v1/user/:userId/permission [get]
func GetUserPermissions(ctx *fiber.Ctx) error {
	id, _ := ctx.ParamsInt("userId")

	if permissions, err := service.GetUserPermissions(id); err != nil {
		log.Println("Error in permission service", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error in user service",
		})
	} else {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"data": permissions,
		})
	}
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
	id, _ := ctx.ParamsInt("id")

	if permission, err := service.GetOnePermission(id); err != nil {
		log.Println("Error in permission service", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error in user service",
		})
	} else {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"data": permission,
		})
	}
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
	body := new(request.UserPermission)
	id, _ := ctx.ParamsInt("id")

	if err := ctx.BodyParser(body); err != nil {
		log.Println("Error in parsing request", err)
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Error in parsing request",
		})
	}

	if permission, err := service.UpdatePermission(*body, id); err != nil {
		log.Println("Error in permission service", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error in user service",
		})
	} else {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"data": permission,
		})
	}
}

// @Summary DeletePermission
// @Tags Permission
// @Description Delete Permission
// @ID delete-permission
// @Accept json
// @Produce	json
// @Success 200 {string} string
// @Router /api/v1/user/:userId/permission [delete]
func DeleteUserPermissions(ctx *fiber.Ctx) error {
	id, _ := ctx.ParamsInt("userId")

	if err := service.DeleteUserPermissions(id); err != nil {
		log.Println("Error in permission service", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error in user service",
		})
	} else {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Permission deleted",
		})
	}
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
	id, _ := ctx.ParamsInt("id")

	if err := service.DeletePermission(id); err != nil {
		log.Println("Error in permission service", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error in user service",
		})
	} else {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Permission deleted",
		})
	}
}
