package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/spitfireooo/form-constructor-server-v2/pkg/controller/v1"
)

func FormRouter(group fiber.Router, path string) {
	form := group.Group(path)
	{
		form.Post("/", v1.CreateForm)
		form.Get("/", v1.GetAllForms)
		form.Get("/:id", v1.GetOneForm)
		form.Patch("/:id", v1.UpdateForm)
		form.Delete("/:id", v1.DeleteForm)

		field := form.Group("/:formId/field")
		{
			field.Post("/", v1.CreateField)
			field.Get("/", v1.GetAllFields)
			field.Get("/:id", v1.GetOneField)
			field.Patch("/:id", v1.UpdateField)
			field.Delete("/:id", v1.DeleteField)
		}
	}
}

func FieldRouter(group fiber.Router, path string) {
	field := group.Group(path)
	{
		multiply := field.Group("/multiply")
		{
			multiply.Post("/", v1.CreateFieldMultiply)
			multiply.Get("/", v1.GetFieldMultiply)
			multiply.Patch("/", v1.UpdateFieldMultiply)
			multiply.Delete("/", v1.DeleteFieldMultiply)
		}

		placeholder := field.Group("/placeholder")
		{
			placeholder.Post("/", v1.CreateFieldPlaceholder)
			placeholder.Get("/", v1.GetFieldPlaceholder)
			placeholder.Patch("/", v1.UpdateFieldPlaceholder)
			placeholder.Delete("/", v1.DeleteFieldPlaceholder)
		}

		fieldRange := field.Group("/range")
		{
			fieldRange.Post("/", v1.CreateFieldRange)
			fieldRange.Get("/", v1.GetFieldRange)
			fieldRange.Patch("/", v1.UpdateFieldRange)
			fieldRange.Delete("/", v1.DeleteFieldRange)
		}

		variants := field.Group("/variants")
		{
			variants.Post("/", v1.CreateFieldVariants)
			variants.Get("/", v1.GetFieldVariants)
			variants.Patch("/", v1.UpdateFieldVariants)
			variants.Delete("/", v1.DeleteFieldVariants)
		}
	}
}
