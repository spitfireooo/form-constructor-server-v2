package controller_v1

import (
	"github.com/gofiber/fiber/v2"
)

// @Summary	CreateTag
// @Tags tag
// @Description Create Tag
// @ID create-tag
// @Accept json
// @Produce	json
// @Param input	body request.Tag true "body info"
// @Success 200 {object} response.Tag
// @Router /api/v1/user/:userId/tag [post]
func CreateTag(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "CreateTag",
	})
}

// @Summary	GetAllTags
// @Tags Tag
// @Description Get All Tags
// @ID get-all-tags
// @Accept json
// @Produce	json
// @Success 200 {array} response.Tag
// @Router /api/v1/user/:userId/tag [get]
func GetAllTags(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "GetAllTags",
	})
}

// @Summary	GetOneTag
// @Tags Tag
// @Description Get One Tag
// @ID get-one-tag
// @Accept json
// @Produce	json
// @Success 200 {object} response.Tag
// @Router /api/v1/user/:userId/tag/:id [get]
func GetOneTag(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "GetOneTag",
	})
}

// @Summary	UpdateTag
// @Tags Tag
// @Description Update Tag
// @ID update-tag
// @Accept json
// @Produce	json
// @Param input	body request.Tag true "body info"
// @Success 200 {object} response.Tag
// @Router /api/v1/user/:userId/tag/:id [patch]
func UpdateTag(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "UpdateTag",
	})
}

// @Summary DeleteTag
// @Tags Tag
// @Description Delete Tag
// @ID delete-tag
// @Accept json
// @Produce	json
// @Success 200 {string} string
// @Router /api/v1/user/:userId/tag/:id [delete]
func DeleteTag(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "DeleteTag",
	})
}
