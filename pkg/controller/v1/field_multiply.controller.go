package controller_v1

import (
	"github.com/gofiber/fiber/v2"
)

// @Summary	CreateFieldMultiply
// @Tags FieldMultiply
// @Description Create Field Multiply
// @ID create-field-multiply
// @Accept json
// @Produce	json
// @Param input	body request.FieldMultiply true "body info"
// @Success 200 {object} response.FieldMultiply
// @Router /api/v1/field/:fieldId/multiply [post]
func CreateFieldMultiply(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "CreateFieldMultiply",
	})
}

// @Summary	GetFieldMultiply
// @Tags FieldMultiply
// @Description Get Field Multiply
// @ID get-field-multiply
// @Accept json
// @Produce	json
// @Success 200 {object} response.FieldMultiply
// @Router /api/v1/field/:fieldId/multiply [get]
func GetFieldMultiply(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "GetFieldMultiply",
	})
}

// @Summary	UpdateFieldMultiply
// @Tags FieldMultiply
// @Description Update Field Multiply
// @ID update-field-multiply
// @Accept json
// @Produce	json
// @Param input	body request.FieldMultiply true "body info"
// @Success 200 {object} response.FieldMultiply
// @Router /api/v1/field/:fieldId/multiply [patch]
func UpdateFieldMultiply(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "UpdateFieldMultiply",
	})
}

// @Summary DeleteFieldMultiply
// @Tags FieldMultiply
// @Description Delete Field Multiply
// @ID delete-field-multiply
// @Accept json
// @Produce	json
// @Success 200 {string} string
// @Router /api/v1/field/:fieldId/multiply [delete]
func DeleteFieldMultiply(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "DeleteFieldMultiply",
	})
}
