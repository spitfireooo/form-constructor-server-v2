package controller_v1

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/spitfireooo/form-constructor-server-v2/pkg/model/request"
	"github.com/spitfireooo/form-constructor-server-v2/pkg/service"
	"log"
	"net/http"
)

// @Summary	CreateTag
// @Tags Tag
// @Description Create Tag
// @ID create-tag
// @Accept json
// @Produce	json
// @Param input	body request.Tag true "body info"
// @Success 200 {object} response.Tag
// @Router /api/v1/tag [post]
func CreateTag(ctx *fiber.Ctx) error {
	body := new(request.Tag)

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

	if tag, err := service.CreateTag(*body); err != nil {
		log.Println("Error in permission service", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error in user service",
		})
	} else {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"data": tag,
		})
	}
}

// @Summary	GetAllUserTags
// @Tags Tag
// @Description Get All User Tags
// @ID get-all-user-tags
// @Accept json
// @Produce	json
// @Success 200 {array} response.UserTag
// @Router /api/v1/tag [get]
func GetAllUserTags(ctx *fiber.Ctx) error {
	if tags, err := service.GetAllUserTags(); err != nil {
		log.Println("Error in permission service", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error in user service",
		})
	} else {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"data": tags,
		})
	}
}

// @Summary	GetAllTags
// @Tags Tag
// @Description Get All Tags
// @ID get-all-tags
// @Accept json
// @Produce	json
// @Success 200 {array} response.Tag
// @Router /api/v1/tag [get]
func GetAllTags(ctx *fiber.Ctx) error {
	if tags, err := service.GetAllTags(); err != nil {
		log.Println("Error in permission service", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error in user service",
		})
	} else {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"data": tags,
		})
	}
}

// @Summary	GetOneTag
// @Tags Tag
// @Description Get One Tag
// @ID get-one-tag
// @Accept json
// @Produce	json
// @Success 200 {object} response.Tag
// @Router /api/v1/tag/:id [get]
func GetOneTag(ctx *fiber.Ctx) error {
	id, _ := ctx.ParamsInt("id")

	if tags, err := service.GetOneTag(id); err != nil {
		log.Println("Error in permission service", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error in user service",
		})
	} else {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"data": tags,
		})
	}
}

// @Summary	UpdateTag
// @Tags Tag
// @Description Update Tag
// @ID update-tag
// @Accept json
// @Produce	json
// @Param input	body request.TagUpdate true "body info"
// @Success 200 {object} response.Tag
// @Router /api/v1/tag/:id [patch]
func UpdateTag(ctx *fiber.Ctx) error {
	body := new(request.TagUpdate)
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

	if tag, err := service.UpdateTag(*body, id); err != nil {
		log.Println("Error in permission service", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error in user service",
		})
	} else {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"data": tag,
		})
	}
}

// @Summary DeleteTag
// @Tags Tag
// @Description Delete Tag
// @ID delete-tag
// @Accept json
// @Produce	json
// @Success 200 {string} string
// @Router /api/v1/tag/:id [delete]
func DeleteTag(ctx *fiber.Ctx) error {
	id, _ := ctx.ParamsInt("id")

	if err := service.DeleteTag(id); err != nil {
		log.Println("Error in permission service", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error in user service",
		})
	} else {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"data": "Tag deleted",
		})
	}
}
