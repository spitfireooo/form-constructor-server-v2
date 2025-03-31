package router

import (
	"github.com/gofiber/fiber/v2"
	middleware "github.com/spitfireooo/form-constructor-auth/pkg/middlewares"
	"github.com/spitfireooo/form-constructor-server-v2/pkg/controller/v1"
)

func FormRouter(group fiber.Router, path string) {
	form := group.Group(path)
	{
		form.Post("/", middleware.IsAuthorized, controller_v1.CreateFormAuth)
		form.Get("/", controller_v1.GetAllForms)
		form.Get("/:id", controller_v1.GetOneForm)
		form.Post("/:id", controller_v1.CreateForm)
		form.Patch("/:id", controller_v1.UpdateForm)
		form.Delete("/:id", controller_v1.DeleteForm)

		field := form.Group("/:formId/field")
		{
			field.Get("/", controller_v1.GetAllFormFields)
			field.Post("/", controller_v1.CreateField)
			field.Delete("/", controller_v1.DeleteFormFields)
		}
	}
}

func FieldRouter(group fiber.Router, path string) {
	field := group.Group(path)
	{
		field.Get("/", controller_v1.GetAllFields)
		field.Get("/:id", controller_v1.GetOneField)
		field.Patch("/:id", controller_v1.UpdateField)
		field.Delete("/:id", controller_v1.DeleteField)

		multiply := field.Group("/:fieldId/multiply")
		{
			multiply.Post("/", controller_v1.CreateFieldMultiply)
			multiply.Get("/", controller_v1.GetFieldMultiply)
			multiply.Patch("/", controller_v1.UpdateFieldMultiply)
			multiply.Delete("/", controller_v1.DeleteFieldMultiply)
		}

		placeholder := field.Group("/:fieldId/placeholder")
		{
			placeholder.Post("/", controller_v1.CreateFieldPlaceholder)
			placeholder.Get("/", controller_v1.GetFieldPlaceholder)
			placeholder.Patch("/", controller_v1.UpdateFieldPlaceholder)
			placeholder.Delete("/", controller_v1.DeleteFieldPlaceholder)
		}

		fieldRange := field.Group("/:fieldId/range")
		{
			fieldRange.Post("/", controller_v1.CreateFieldRange)
			fieldRange.Get("/", controller_v1.GetFieldRange)
			fieldRange.Patch("/", controller_v1.UpdateFieldRange)
			fieldRange.Delete("/", controller_v1.DeleteFieldRange)
		}

		variants := field.Group("/:fieldId/variants")
		{
			variants.Post("/", controller_v1.CreateFieldVariant)
			variants.Get("/", controller_v1.GetFieldVariants)
			variants.Delete("/", controller_v1.DeleteFieldVariants)
		}
	}

	variants := group.Group("/variants")
	{
		variants.Get("/", controller_v1.GetAllFieldVariants)
		variants.Get("/:id", controller_v1.GetOneFieldVariants)
		variants.Patch("/:id", controller_v1.UpdateFieldVariants)
		variants.Delete("/:id", controller_v1.DeleteOneFieldVariants)
	}
}
