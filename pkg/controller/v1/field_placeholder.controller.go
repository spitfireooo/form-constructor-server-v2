package controller_v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/spitfireooo/form-constructor-server-v2/pkg/model/request"
	"github.com/spitfireooo/form-constructor-server-v2/pkg/service"
	"log"
	"net/http"
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
	body := new(request.FieldPlaceholder)
	id, _ := ctx.ParamsInt("fieldId")

	if err := ctx.BodyParser(body); err != nil {
		log.Println("Error in parsing request", err)
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Error in parsing request",
		})
	}

	if res, err := service.CreateFieldPlaceholder(*body, id); err != nil {
		log.Println("Error in placeholder service", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error in placeholder service",
		})
	} else {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"data": res,
		})
	}
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
	id, _ := ctx.ParamsInt("fieldId")

	if res, err := service.GetFieldPlaceholder(id); err != nil {
		log.Println("Error in placeholder service", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error in placeholder service",
		})
	} else {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"data": res,
		})
	}
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
	body := new(request.FieldPlaceholder)
	id, _ := ctx.ParamsInt("fieldId")

	if err := ctx.BodyParser(body); err != nil {
		log.Println("Error in parsing request", err)
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Error in parsing request",
		})
	}

	if res, err := service.UpdateFieldPlaceholder(*body, id); err != nil {
		log.Println("Error in placeholder service", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error in placeholder service",
		})
	} else {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"data": res,
		})
	}
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
	id, _ := ctx.ParamsInt("fieldId")

	if err := service.DeleteFieldPlaceholder(id); err != nil {
		log.Println("Error in placeholder service", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error in placeholder service",
		})
	} else {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Placeholder deleted",
		})
	}
}
