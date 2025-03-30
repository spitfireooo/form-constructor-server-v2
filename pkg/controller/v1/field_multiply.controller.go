package controller_v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/spitfireooo/form-constructor-server-v2/pkg/model/request"
	"github.com/spitfireooo/form-constructor-server-v2/pkg/service"
	"log"
	"net/http"
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
	body := new(request.FieldMultiply)
	id, _ := ctx.ParamsInt("fieldId")

	if err := ctx.BodyParser(body); err != nil {
		log.Println("Error in parsing request", err)
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Error in parsing request",
		})
	}

	if res, err := service.CreateFieldMultiply(*body, id); err != nil {
		log.Println("Error in multiply service", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error in multiply service",
		})
	} else {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"data": res,
		})
	}
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
	id, _ := ctx.ParamsInt("fieldId")

	if res, err := service.GetFieldMultiply(id); err != nil {
		log.Println("Error in multiply service", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error in multiply service",
		})
	} else {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"data": res,
		})
	}
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
	body := new(request.FieldMultiply)
	id, _ := ctx.ParamsInt("fieldId")

	if err := ctx.BodyParser(body); err != nil {
		log.Println("Error in parsing request", err)
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Error in parsing request",
		})
	}

	if res, err := service.UpdateFieldMultiply(*body, id); err != nil {
		log.Println("Error in multiply service", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error in multiply service",
		})
	} else {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"data": res,
		})
	}
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
	id, _ := ctx.ParamsInt("fieldId")

	if err := service.DeleteFieldMultiply(id); err != nil {
		log.Println("Error in multiply service", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error in multiply service",
		})
	} else {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Multiply deleted",
		})
	}
}
