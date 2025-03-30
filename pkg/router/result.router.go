package router

import (
	"github.com/gofiber/fiber/v2"
	controller_v1 "github.com/spitfireooo/form-constructor-server-v2/pkg/controller/v1"
)

func ResultRouter(group fiber.Router, path string) {
	form := group.Group(path)
	{
		form.Post("/", controller_v1.CreateResult)
		form.Get("/", controller_v1.GetAllResults)
		form.Get("/:id", controller_v1.GetOneResult)
		//form.Patch("/:id", controller_v1.UpdateResult)
		form.Delete("/:id", controller_v1.DeleteResult)

		form.Get("/form/:formId", controller_v1.GetFormResults)
		form.Get("/field/:fieldId", controller_v1.GetFieldResults)
		form.Delete("/form/:formId", controller_v1.DeleteFormResults)
		form.Delete("/field/:fieldId", controller_v1.DeleteFieldResults)
	}
}
