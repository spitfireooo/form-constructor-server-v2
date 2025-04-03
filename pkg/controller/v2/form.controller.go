package controller_v2

import (
	"github.com/gofiber/fiber/v2"
	"github.com/spitfireooo/form-constructor-server-v2/pkg/model/request"
	service_v2 "github.com/spitfireooo/form-constructor-server-v2/pkg/service/v2"
	"log"
	"net/http"
)

// @Summary	CreateFieldWithField
// @Tags Form
// @Description Create Form With Field
// @ID create-form-v2
// @Accept json
// @Produce	json
// @Param input	body request.FormWithField true "body info"
// @Success 200 {object} response.FormWithField
// @Router /api/v2/form/:id [post]
func CreateForm(ctx *fiber.Ctx) error {
	var body request.FormWithField
	userId, _ := ctx.ParamsInt("id")

	if err := ctx.BodyParser(&body); err != nil {
		log.Println("Error in parsing request", err)
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Error in parsing request",
		})
	}

	if res, err := service_v2.CreateForm(body, userId); err != nil {
		log.Println("Error in field service", err)
		return err
		//return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
		//	"message": "Error in field service",
		//})
	} else {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"data": res,
		})
	}
}

// @Summary	GetFormWithField
// @Tags Form
// @Description Get Form With Field
// @ID get-form-v2
// @Accept json
// @Produce	json
// @Success 200 {array} response.FormWithField
// @Router /api/v2/form/:id [get]
func GetForm(ctx *fiber.Ctx) error {
	formId, _ := ctx.ParamsInt("id")

	if res, err := service_v2.GetForm(formId); err != nil {
		log.Println("Error in field service", err)
		return err
		//return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
		//	"message": "Error in field service",
		//})
	} else {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"data": res,
		})
	}
}
