package controller_v1

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/spitfireooo/form-constructor-server-v2/pkg/model/request"
	"github.com/spitfireooo/form-constructor-server-v2/pkg/service"
	"log"
	"net/http"
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
	body := new(request.Field)
	formId, _ := ctx.ParamsInt("formId")

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

	if res, err := service.CreateField(*body, formId); err != nil {
		log.Println("Error in field service", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error in field service",
		})
	} else {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"data": res,
		})
	}
}

// @Summary	GetAllFields
// @Tags Field
// @Description Get All Fields
// @ID get-all-fields
// @Accept json
// @Produce	json
// @Success 200 {array} response.Field
// @Router /api/v1/field [get]
func GetAllFields(ctx *fiber.Ctx) error {
	if res, err := service.GetAllFields(); err != nil {
		log.Println("Error in field service", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error in field service",
		})
	} else {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"data": res,
		})
	}
}

// @Summary	GetAllFields
// @Tags Field
// @Description Get All Fields
// @ID get-all-fields
// @Accept json
// @Produce	json
// @Success 200 {array} response.Field
// @Router /api/v1/form/:formId/field [get]
func GetAllFormFields(ctx *fiber.Ctx) error {
	formId, _ := ctx.ParamsInt("formId")

	if res, err := service.GetAllFormFields(formId); err != nil {
		log.Println("Error in field service", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error in field service",
		})
	} else {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"data": res,
		})
	}
}

// @Summary	GetOneField
// @Tags Field
// @Description Get One Field
// @ID get-one-field
// @Accept json
// @Produce	json
// @Success 200 {object} response.Field
// @Router /api/v1/field/:id [get]
func GetOneField(ctx *fiber.Ctx) error {
	id, _ := ctx.ParamsInt("id")

	if res, err := service.GetOneField(id); err != nil {
		log.Println("Error in field service", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error in field service",
		})
	} else {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"data": res,
		})
	}
}

// @Summary	UpdateField
// @Tags Field
// @Description Update Field
// @ID update-field
// @Accept json
// @Produce	json
// @Param input	body request.Field true "body info"
// @Success 200 {object} response.Field
// @Router /api/v1/field/:id [patch]
func UpdateField(ctx *fiber.Ctx) error {
	body := new(request.FieldUpdate)
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

	if res, err := service.UpdateField(*body, id); err != nil {
		log.Println("Error in field service", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error in field service",
		})
	} else {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"data": res,
		})
	}
}

// @Summary DeleteField
// @Tags Field
// @Description Delete Field
// @ID delete-field
// @Accept json
// @Produce	json
// @Success 200 {string} string
// @Router /api/v1/field/:id [delete]
func DeleteField(ctx *fiber.Ctx) error {
	id, _ := ctx.ParamsInt("id")

	if err := service.DeleteField(id); err != nil {
		log.Println("Error in field service", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error in field service",
		})
	} else {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Field deleted",
		})
	}
}

// @Summary DeleteField
// @Tags Field
// @Description Delete Field
// @ID delete-field
// @Accept json
// @Produce	json
// @Success 200 {string} string
// @Router /api/v1/form/:formId/field [delete]
func DeleteFormFields(ctx *fiber.Ctx) error {
	formId, _ := ctx.ParamsInt("formId")

	if err := service.DeleteFormFields(formId); err != nil {
		log.Println("Error in field service", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error in field service",
		})
	} else {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Form Fields deleted",
		})
	}
}
