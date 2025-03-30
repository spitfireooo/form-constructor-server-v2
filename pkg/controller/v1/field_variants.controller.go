package controller_v1

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/spitfireooo/form-constructor-server-v2/pkg/model/request"
	"github.com/spitfireooo/form-constructor-server-v2/pkg/service"
	"log"
	"net/http"
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
func CreateFieldVariant(ctx *fiber.Ctx) error {
	body := new(request.FieldVariants)
	id, _ := ctx.ParamsInt("fieldId")

	if err := ctx.BodyParser(body); err != nil {
		log.Println("Error in parsing request", err)
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Error in parsing request",
		})
	}

	if res, err := service.CreateFieldVariant(*body, id); err != nil {
		log.Println("Error in variant service", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error in variant service",
		})
	} else {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"data": res,
		})
	}
}

// @Summary	GetFieldVariants
// @Tags FieldVariants
// @Description Get Field Variants
// @ID get-field-variants
// @Accept json
// @Produce	json
// @Success 200 {array} response.FieldVariants
// @Router /api/v1/field/:fieldId/variants [get]
func GetFieldVariants(ctx *fiber.Ctx) error {
	id, _ := ctx.ParamsInt("fieldId")

	if res, err := service.GetFieldVariants(id); err != nil {
		log.Println("Error in variant service", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error in variant service",
		})
	} else {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"data": res,
		})
	}
}

// @Summary	GetAllFieldVariants
// @Tags FieldVariants
// @Description Get All Field Variants
// @ID get-all-field-variants
// @Accept json
// @Produce	json
// @Success 200 {array} response.FieldVariants
// @Router /api/v1/variants [get]
func GetAllFieldVariants(ctx *fiber.Ctx) error {
	if res, err := service.GetAllFieldVariants(); err != nil {
		log.Println("Error in variant service", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error in variant service",
		})
	} else {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"data": res,
		})
	}
}

// @Summary	GetOneFieldVariants
// @Tags FieldVariants
// @Description Get One Field Variants
// @ID get-one-field-variants
// @Accept json
// @Produce	json
// @Success 200 {object} response.FieldVariants
// @Router /api/v1/variants/:id [get]
func GetOneFieldVariants(ctx *fiber.Ctx) error {
	id, _ := ctx.ParamsInt("id")

	if res, err := service.GetOneFieldVariants(id); err != nil {
		log.Println("Error in variant service", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error in variant service",
		})
	} else {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"data": res,
		})
	}
}

// @Summary	UpdateFieldVariants
// @Tags FieldVariants
// @Description Update Field Variant
// @ID update-field-variants
// @Accept json
// @Produce	json
// @Param input	body request.FieldVariants true "body info"
// @Success 200 {object} response.FieldVariants
// @Router /api/v1/variants/:id [patch]
func UpdateFieldVariants(ctx *fiber.Ctx) error {
	body := new(request.FieldVariantsUpdate)
	id, _ := ctx.ParamsInt("id")

	if err := ctx.BodyParser(body); err != nil {
		log.Println("Error in parsing request", err)
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Error in parsing request",
		})
	}

	if err := validator.New().Struct(body); err != nil {
		log.Println("Validation errors", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Validation errors",
		})
	}

	if res, err := service.UpdateFieldVariants(*body, id); err != nil {
		log.Println("Error in variant service", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error in variant service",
		})
	} else {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"data": res,
		})
	}
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
	id, _ := ctx.ParamsInt("fieldId")

	if err := service.DeleteFieldVariants(id); err != nil {
		log.Println("Error in variant service", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error in variant service",
		})
	} else {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Range deleted",
		})
	}
}

// @Summary DeleteFieldVariants
// @Tags FieldVariants
// @Description Delete Field Variants
// @ID delete-field-variants
// @Accept json
// @Produce	json
// @Success 200 {string} string
// @Router /api/v1/variants/:id [delete]
func DeleteOneFieldVariants(ctx *fiber.Ctx) error {
	id, _ := ctx.ParamsInt("id")

	if err := service.DeleteOneFieldVariant(id); err != nil {
		log.Println("Error in variant service", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error in variant service",
		})
	} else {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Range deleted",
		})
	}
}
