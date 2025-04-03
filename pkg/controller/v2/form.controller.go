package controller_v2

import (
	"github.com/gofiber/fiber/v2"
	"github.com/spitfireooo/form-constructor-server-v2/pkg/model/request"
	service_v2 "github.com/spitfireooo/form-constructor-server-v2/pkg/service/v2"
	"log"
	"net/http"
)

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

func GetForm(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "GetOneForm",
	})
}
