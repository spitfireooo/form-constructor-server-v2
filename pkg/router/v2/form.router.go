package router_v2

import (
	"github.com/gofiber/fiber/v2"
	controller_v2 "github.com/spitfireooo/form-constructor-server-v2/pkg/controller/v2"
)

func FormRouter(group fiber.Router, path string) {
	form := group.Group(path)
	{
		form.Post("/:userId", controller_v2.CreateForm)
		form.Get("/:id", controller_v2.GetForm)
		form.Patch("/:id", controller_v2.UpdateForm)
		form.Post("/:userId/upload/:slug", controller_v2.UploadForm)

		field := form.Group("/:formId/field")
		{
			field.Post("/", controller_v2.CreateField)
			field.Get("/:id", controller_v2.GetOneField)
			field.Patch("/:id", controller_v2.UpdateField)
		}
	}
}
