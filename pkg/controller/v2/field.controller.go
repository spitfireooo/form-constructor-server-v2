package controller_v2

import (
	"github.com/gofiber/fiber/v2"
	"github.com/spitfireooo/form-constructor-server-v2/pkg/model/response"
	service_v2 "github.com/spitfireooo/form-constructor-server-v2/pkg/service/v2"
	"log"
	"net/http"
)

type Field interface {
	response.FieldString |
		response.FieldNumber |
		response.FieldEmail |
		response.FieldText |
		response.FieldDate |
		response.FieldRadio |
		response.FieldSelect
}

func CreateField(ctx *fiber.Ctx) error {
	formId, _ := ctx.ParamsInt("formId")

	var body map[string]interface{}

	if err := ctx.BodyParser(&body); err != nil {
		log.Println("Error in parsing request", err)
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Error in parsing request",
		})
	}

	if res, err := service_v2.CreateField(body, formId); err != nil {
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

func GetOneField(ctx *fiber.Ctx) error {
	id, _ := ctx.ParamsInt("id")

	if res, err := service_v2.GetOneField(id); err != nil {
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
