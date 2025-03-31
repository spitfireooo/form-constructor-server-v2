package router

import (
	"github.com/gofiber/fiber/v2"
	controller_v1 "github.com/spitfireooo/form-constructor-server-v2/pkg/controller/v1"
)

func ResultRouter(group fiber.Router, path string) {
	result := group.Group(path)
	{
		result.Post("/", controller_v1.CreateResults)
		result.Get("/", controller_v1.GetAllResults)
		result.Get("/:id", controller_v1.GetOneResult)
		result.Patch("/:id", controller_v1.UpdateResult)
		result.Delete("/:id", controller_v1.DeleteResult)

		result.Post("/field/:fieldId", controller_v1.CreateResult)
		result.Get("/form/:formId", controller_v1.GetFormResults)
		result.Get("/field/:fieldId", controller_v1.GetFieldResults)
		result.Delete("/form/:formId", controller_v1.DeleteFormResults)
		result.Delete("/field/:fieldId", controller_v1.DeleteFieldResults)
	}
}
