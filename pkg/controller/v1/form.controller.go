package controller_v1

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/spitfireooo/form-constructor-server-v2/pkg/model/request"
	"github.com/spitfireooo/form-constructor-server-v2/pkg/service"
	"github.com/spitfireooo/form-constructor-server-v2/pkg/utils"
	"log"
	"net/http"
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
func CreateFormAuth(ctx *fiber.Ctx) error {
	body := new(request.Form)
	userId := int(ctx.Locals("user_id").(int64))

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

	if file, err := ctx.FormFile("logo"); err != nil {
		log.Println("Error in file upload", err)
	} else if file.Size > 0 {
		fmt.Println("file", file.Header.Get("Content-Type"))
		if err = utils.CheckContentType(
			file.Header.Get("Content-Type"),
			"image/jpg",
			"image/png",
			"image/gif",
			"image/jpeg",
		); err != nil {
			log.Println("Bad ext of file", err)
			return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
				"message": "Bad ext of file",
			})
		}

		filename := utils.GenerateFilename(file.Filename)
		if err = ctx.SaveFile(file, filename); err != nil {
			log.Println("Error in save file", err)
		}
		body.Logo = &filename
	}

	if res, err := service.CreateForm(*body, userId); err != nil {
		log.Println("Error in form service", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error in form service",
		})
	} else {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"data": res,
		})
	}
}

// @Summary	CreateForm
// @Tags Form
// @Description Create Form
// @ID create-form
// @Accept json
// @Produce	json
// @Param input	body request.Form true "body info"
// @Success 200 {object} response.Form
// @Router /api/v1/form/:id [post]
func CreateForm(ctx *fiber.Ctx) error {
	body := new(request.Form)
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

	if file, err := ctx.FormFile("logo"); err != nil {
		log.Println("Error in file upload", err)
	} else if file.Size > 0 {
		if err = utils.CheckContentType(
			file.Header.Get("Content-Type"),
			"image/jpg",
			"image/png",
			"image/gif",
			"image/jpeg",
		); err != nil {
			log.Println("Bad ext of file", err)
			return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
				"message": "Bad ext of file",
			})
		}

		filename := utils.GenerateFilename(file.Filename)
		if err = ctx.SaveFile(file, filename); err != nil {
			log.Println("Error in save file", err)
		}
		body.Logo = &filename
	}

	if res, err := service.CreateForm(*body, id); err != nil {
		log.Println("Error in form service", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error in form service",
		})
	} else {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"data": res,
		})
	}
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
	if res, err := service.GetAllForms(); err != nil {
		log.Println("Error in form service", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error in form service",
		})
	} else {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"data": res,
		})
	}
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
	id, _ := ctx.ParamsInt("id")

	if res, err := service.GetOneForm(id); err != nil {
		log.Println("Error in form service", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error in form service",
		})
	} else {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"data": res,
		})
	}
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
	body := new(request.FormUpdate)
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

	if file, err := ctx.FormFile("logo"); err != nil {
		log.Println("Error in file upload", err)
	} else if file.Size > 0 {
		fmt.Println("file", file.Header.Get("Content-Type"))
		if err = utils.CheckContentType(
			file.Header.Get("Content-Type"),
			"image/jpg",
			"image/png",
			"image/gif",
			"image/jpeg",
		); err != nil {
			log.Println("Bad ext of file", err)
			return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
				"message": "Bad ext of file",
			})
		}

		filename := utils.GenerateFilename(file.Filename)
		if err = ctx.SaveFile(file, filename); err != nil {
			log.Println("Error in save file", err)
		}
		body.Logo = &filename
	}

	if res, err := service.UpdateForm(*body, id); err != nil {
		log.Println("Error in form service", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error in form service",
		})
	} else {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"data": res,
		})
	}
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
	id, _ := ctx.ParamsInt("id")

	if err := service.DeleteForm(id); err != nil {
		log.Println("Error in form service", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error in form service",
		})
	} else {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Form deleted!",
		})
	}
}
