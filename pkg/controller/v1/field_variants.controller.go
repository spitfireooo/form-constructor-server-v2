package controller_v1

import (
	"github.com/gofiber/fiber/v2"
)

// @Summary	CreateFieldVariants
// @Tags FieldVariants
// @Description Create Field Variants
// @ID create-field-variants
// @Accept json
// @Produce	json
// @Param input	body request.FieldVariants true "body info"
// @Success 200 {object} response.FieldVariants
// @Router /api/v1/field/:fieldId/variants [post]
func CreateFieldVariants(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "CreateFieldVariants",
	})
}

// @Summary	GetFieldVariants
// @Tags FieldVariants
// @Description Get Field Variants
// @ID get-field-variants
// @Accept json
// @Produce	json
// @Success 200 {object} response.FieldVariants
// @Router /api/v1/field/:fieldId/variants [get]
func GetFieldVariants(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "GetFieldVariants",
	})
}

// @Summary	UpdateFieldVariants
// @Tags FieldVariants
// @Description Update Field Variant
// @ID update-field-variants
// @Accept json
// @Produce	json
// @Param input	body request.FieldVariants true "body info"
// @Success 200 {object} response.FieldVariants
// @Router /api/v1/field/:fieldId/variants [patch]
func UpdateFieldVariants(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "UpdateFieldVariants",
	})
}

// @Summary DeleteFieldVariants
// @Tags FieldVariants
// @Description Delete Field Variants
// @ID delete-field-variants
// @Accept json
// @Produce	json
// @Success 200 {string} string
// @Router /api/v1/field/:fieldId/variants [delete]
func DeleteFieldVariants(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "DeleteFieldVariants",
	})
}
