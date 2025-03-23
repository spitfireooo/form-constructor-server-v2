package controller_v1

import (
	"github.com/gofiber/fiber/v2"
)

// @Summary	CreateField
// @Tags Field
// @Description Create Field
// @ID create-field
// @Accept json
// @Produce	json
// @Param input	body request.Field true "body info"
// @Success 200 {object} response.Field
// @Router /api/v1/form/:formId/field [post]
func CreateField(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "CreateField",
	})
}

// @Summary	GetAllFields
// @Tags Field
// @Description Get All Fields
// @ID get-all-fields
// @Accept json
// @Produce	json
// @Success 200 {array} response.Field
// @Router /api/v1/form/:formId/field [get]
func GetAllFields(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "GetAllFields",
	})
}

// @Summary	GetOneField
// @Tags Field
// @Description Get One Field
// @ID get-one-field
// @Accept json
// @Produce	json
// @Success 200 {object} response.Field
// @Router /api/v1/form/:formId/field/:id [get]
func GetOneField(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "GetOneField",
	})
}

// @Summary	UpdateField
// @Tags Field
// @Description Update Field
// @ID update-field
// @Accept json
// @Produce	json
// @Param input	body request.Field true "body info"
// @Success 200 {object} response.Field
// @Router /api/v1/form/:formId/field/:id [patch]
func UpdateField(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "UpdateField",
	})
}

// @Summary DeleteField
// @Tags Field
// @Description Delete Field
// @ID delete-field
// @Accept json
// @Produce	json
// @Success 200 {string} string
// @Router /api/v1/form/:formId/field/:id [delete]
func DeleteField(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "DeleteField",
	})
}
