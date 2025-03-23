package controller_v1

import (
	"github.com/gofiber/fiber/v2"
)

// @Summary	CreateFieldPlaceholder
// @Tags FieldPlaceholder
// @Description Create Field Placeholder
// @ID create-field-placeholder
// @Accept json
// @Produce	json
// @Param input	body request.FieldPlaceholder true "body info"
// @Success 200 {object} response.FieldPlaceholder
// @Router /api/v1/field/:fieldId/placeholder [post]
func CreateFieldPlaceholder(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "CreateFieldPlaceholder",
	})
}

// @Summary	GetFieldPlaceholder
// @Tags FieldPlaceholder
// @Description Get Field Placeholder
// @ID get-field-placeholder
// @Accept json
// @Produce	json
// @Success 200 {object} response.FieldPlaceholder
// @Router /api/v1/field/:fieldId/placeholder [get]
func GetFieldPlaceholder(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "GetFieldPlaceholder",
	})
}

// @Summary	UpdateFieldPlaceholder
// @Tags FieldPlaceholder
// @Description Update Field Placeholder
// @ID update-field-placeholder
// @Accept json
// @Produce	json
// @Param input	body request.FieldPlaceholder true "body info"
// @Success 200 {object} response.FieldPlaceholder
// @Router /api/v1/field/:fieldId/placeholder [patch]
func UpdateFieldPlaceholder(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "UpdateFieldPlaceholder",
	})
}

// @Summary DeleteFieldPlaceholder
// @Tags FieldPlaceholder
// @Description Delete Field Placeholder
// @ID delete-field-placeholder
// @Accept json
// @Produce	json
// @Success 200 {string} string
// @Router /api/v1/field/:fieldId/placehodler [delete]
func DeleteFieldPlaceholder(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "DeleteFieldPlaceholder",
	})
}
