package controller_v1

import (
	"github.com/gofiber/fiber/v2"
)

// @Summary	CreateFieldRange
// @Tags FieldRange
// @Description Create Field Range
// @ID create-field-range
// @Accept json
// @Produce	json
// @Param input	body request.FieldRange true "body info"
// @Success 200 {object} response.FieldRange
// @Router /api/v1/field/:fieldId/range [post]
func CreateFieldRange(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "CreateFieldRange",
	})
}

// @Summary	GetFieldRange
// @Tags FieldRange
// @Description Get Field Range
// @ID get-field-range
// @Accept json
// @Produce	json
// @Success 200 {object} response.FieldRange
// @Router /api/v1/field/:fieldId/range [get]
func GetFieldRange(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "GetFieldRange",
	})
}

// @Summary	UpdateFieldRange
// @Tags FieldRange
// @Description Update Field Range
// @ID update-field-range
// @Accept json
// @Produce	json
// @Param input	body request.FieldRange true "body info"
// @Success 200 {object} response.FieldRange
// @Router /api/v1/field/:fieldId/range [patch]
func UpdateFieldRange(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "UpdateFieldRange",
	})
}

// @Summary DeleteFieldRange
// @Tags FieldRange
// @Description Delete Field Range
// @ID delete-field-range
// @Accept json
// @Produce	json
// @Success 200 {string} string
// @Router /api/v1/field/:fieldId/range [delete]
func DeleteFieldRange(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "DeleteFieldRange",
	})
}
