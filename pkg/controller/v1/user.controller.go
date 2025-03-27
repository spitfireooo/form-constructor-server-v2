package controller_v1

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/spitfireooo/form-constructor-server-v2/pkg/model/request"
	"github.com/spitfireooo/form-constructor-server-v2/pkg/service"
	"github.com/spitfireooo/form-constructor-server-v2/pkg/utils"
	"log"
	"net/http"
)

// @Summary	GetAllUsers
// @Tags User
// @Description Get All Users
// @ID get-all-users
// @Accept json
// @Produce	json
// @Success 200 {array} response.User
// @Router /api/v1/user [get]
func GetAllUsers(ctx *fiber.Ctx) error {
	if res, err := service.GetAllUsers(); err != nil {
		log.Println("Error in user service", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error in user service",
		})
	} else {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"data": res,
		})
	}
}

// @Summary	GetOneUser
// @Tags User
// @Description Get One User
// @ID get-one-user
// @Accept json
// @Produce	json
// @Success 200 {object} response.User
// @Router /api/v1/user/:id [get]
func GetOneUser(ctx *fiber.Ctx) error {
	id, _ := ctx.ParamsInt("id")

	if res, err := service.GetOneUser(id); err != nil {
		log.Println("Error in user service", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error in user service",
		})
	} else {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"data": res,
		})
	}
}

// @Summary	UpdateUser
// @Tags User
// @Description Update User
// @ID update-user
// @Accept json
// @Produce	json
// @Param input	body request.User true "body info"
// @Success 200 {object} response.User
// @Router /api/v1/user/:id [patch]
func UpdateUser(ctx *fiber.Ctx) error {
	body := new(request.User)
	id := ctx.Params("id")

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

		filename := fmt.Sprintf("./static/uploads/%s_%s", uuid.New(), file.Filename)
		if err = ctx.SaveFile(file, filename); err != nil {
			log.Println("Error in save file", err)
		}
		body.Logo = filename
	}

	if user, err := service.UpdateUser(*body, id); err != nil {
		log.Println("Error in user service", err)
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error in user service",
		})
	} else {
		return ctx.Status(http.StatusOK).JSON(fiber.Map{
			"data": user,
		})
	}
}

// @Summary DeleteUser
// @Tags User
// @Description Delete User
// @ID delete-user
// @Accept json
// @Produce	json
// @Success 200 {string} string
// @Router /api/v1/user/:id [delete]
func DeleteUser(ctx *fiber.Ctx) error {
	id, _ := ctx.ParamsInt("id")

	if err := service.DeleteUser(id); err != nil {
		log.Println("Error in user service", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error in user service",
		})
	} else {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "User deleted!",
		})
	}
}

// @Summary	CreateUserTag
// @Tags User
// @Description Create User Tag
// @ID create-user-tag
// @Accept json
// @Produce	json
// @Param input	body request.UserTag true "body info"
// @Success 200 {object} response.UserTag
// @Router /api/v1/user/:userId/tag [post]
func CreateUserTag(ctx *fiber.Ctx) error {
	body := new(request.UserTag)
	userId, _ := ctx.ParamsInt("userId")

	if err := ctx.BodyParser(body); err != nil {
		log.Println("Error in parsing request", err)
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Error in parsing request",
		})
	}

	if tag, err := service.CreateUserTag(*body, userId); err != nil {
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

// @Summary	GetUserTags
// @Tags User
// @Description Get User Tags
// @ID get-user-tags
// @Accept json
// @Produce	json
// @Success 200 {array} response.UserTag
// @Router /api/v1/user/:userId/tag [get]
func GetUserTags(ctx *fiber.Ctx) error {
	userId, _ := ctx.ParamsInt("userId")

	if tags, err := service.GetUserTags(userId); err != nil {
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

// @Summary	GetUserTag
// @Tags User
// @Description Get User Tag
// @ID get-user-tag
// @Accept json
// @Produce	json
// @Success 200 {object} response.UserTag
// @Router /api/v1/user/:userId/tag/:id [get]
func GetUserTag(ctx *fiber.Ctx) error {
	userId, _ := ctx.ParamsInt("userId")
	id, _ := ctx.ParamsInt("id")

	if tags, err := service.GetUserTag(userId, id); err != nil {
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

// @Summary	UpdateUserTag
// @Tags User
// @Description Update User Tag
// @ID update-user-tag
// @Accept json
// @Produce	json
// @Param input	body request.UserTagUpdate true "body info"
// @Success 200 {object} response.UserTag
// @Router /api/v1/user/:userId/tag/:id [patch]
func UpdateUserTag(ctx *fiber.Ctx) error {
	body := new(request.UserTagUpdate)
	userId, _ := ctx.ParamsInt("userId")
	id, _ := ctx.ParamsInt("id")

	if err := ctx.BodyParser(body); err != nil {
		log.Println("Error in parsing request", err)
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Error in parsing request",
		})
	}

	if tag, err := service.UpdateUserTag(*body, userId, id); err != nil {
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

// @Summary DeleteUserTags
// @Tags User
// @Description Delete User Tags
// @ID delete-user-tags
// @Accept json
// @Produce	json
// @Success 200 {string} string
// @Router /api/v1/user/:userId/tag [delete]
func DeleteUserTags(ctx *fiber.Ctx) error {
	userId, _ := ctx.ParamsInt("userId")

	if err := service.DeleteUserTags(userId); err != nil {
		log.Println("Error in permission service", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error in user service",
		})
	} else {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "User tags deleted",
		})
	}
}

// @Summary DeleteUserTag
// @Tags User
// @Description Delete User Tag
// @ID delete-user-tag
// @Accept json
// @Produce	json
// @Success 200 {string} string
// @Router /api/v1/user/:userId/tag/:id [delete]
func DeleteUserTag(ctx *fiber.Ctx) error {
	userId, _ := ctx.ParamsInt("userId")
	id, _ := ctx.ParamsInt("id")

	if err := service.DeleteUserTag(userId, id); err != nil {
		log.Println("Error in permission service", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error in user service",
		})
	} else {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "User tag deleted",
		})
	}
}
