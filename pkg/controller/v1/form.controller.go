package controller_v1

import (
	"github.com/gofiber/fiber/v2"
)

// @Summary	CreateForm
// @Tags Form
// @Description Create Form
// @ID create-form
// @Accept json
// @Produce	json
// @Param input	body request.Form true "body info"
// @Success 200 {object} response.Form
// @Router /api/v1/form [post]
func CreateForm(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "CreateForm",
	})
}

// @Summary	GetAllForms
// @Tags Form
// @Description Get All Forms
// @ID get-all-forms
// @Accept json
// @Produce	json
// @Success 200 {array} response.Form
// @Router /api/v1/form [get]
func GetAllForms(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "GetAllForms",
	})
}

// @Summary	GetOneForm
// @Tags Form
// @Description Get One Form
// @ID get-one-form
// @Accept json
// @Produce	json
// @Success 200 {object} response.Form
// @Router /api/v1/form/:id [get]
func GetOneForm(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "GetOneForm",
	})
}

// @Summary	UpdateForm
// @Tags Form
// @Description Update Form
// @ID update-form
// @Accept json
// @Produce	json
// @Param input	body request.Form true "body info"
// @Success 200 {object} response.Form
// @Router /api/v1/form/:id [patch]
func UpdateForm(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "UpdateForm",
	})
}

// @Summary DeleteForm
// @Tags Form
// @Description Delete Form
// @ID delete-form
// @Accept json
// @Produce	json
// @Success 200 {string} string
// @Router /api/v1/form/:id [delete]
func DeleteForm(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "DeleteForm",
	})
}
