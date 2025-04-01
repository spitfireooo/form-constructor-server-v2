package router_v2

import (
	"github.com/gofiber/fiber/v2"
	controller_v2 "github.com/spitfireooo/form-constructor-server-v2/pkg/controller/v2"
)

func FormRouter(group fiber.Router, path string) {
	form := group.Group(path)
	{
		form.Post("/", controller_v2.CreateForm)
		form.Get("/:id", controller_v2.GetOneForm)
		form.Patch("/:id", controller_v2.UpdateForm)
		form.Delete("/:id", controller_v2.DeleteForm)

		field := form.Group("/:formId/field")
		{
			field.Post("/", controller_v2.CreateField)
			field.Get("/:id", controller_v2.GetOneField)
			//field.Patch("/:id", controller_v2.UpdateField)
		}
	}
}
