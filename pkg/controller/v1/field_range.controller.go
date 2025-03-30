package controller_v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/spitfireooo/form-constructor-server-v2/pkg/model/request"
	"github.com/spitfireooo/form-constructor-server-v2/pkg/service"
	"log"
	"net/http"
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
	body := new(request.FieldRange)
	id, _ := ctx.ParamsInt("fieldId")

	if err := ctx.BodyParser(body); err != nil {
		log.Println("Error in parsing request", err)
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Error in parsing request",
		})
	}

	if res, err := service.CreateFieldRange(*body, id); err != nil {
		log.Println("Error in range service", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error in range service",
		})
	} else {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"data": res,
		})
	}
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
	id, _ := ctx.ParamsInt("fieldId")

	if res, err := service.GetFieldRange(id); err != nil {
		log.Println("Error in range service", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error in range service",
		})
	} else {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"data": res,
		})
	}
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
	body := new(request.FieldRangeUpdate)
	id, _ := ctx.ParamsInt("fieldId")

	if err := ctx.BodyParser(body); err != nil {
		log.Println("Error in parsing request", err)
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Error in parsing request",
		})
	}

	if res, err := service.UpdateFieldRange(*body, id); err != nil {
		log.Println("Error in range service", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error in range service",
		})
	} else {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"data": res,
		})
	}
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
	id, _ := ctx.ParamsInt("fieldId")

	if err := service.DeleteFieldRange(id); err != nil {
		log.Println("Error in range service", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error in range service",
		})
	} else {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Range deleted",
		})
	}
}
