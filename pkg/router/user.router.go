package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/spitfireooo/form-constructor-server-v2/pkg/controller/v1"
)

func UserRouter(group fiber.Router, path string) {
	user := group.Group(path)
	{
		user.Get("/", v1.GetAllUsers)
		user.Get("/:id", v1.GetOneUser)
		user.Get("/:id", v1.UpdateUser)
		user.Get("/:id", v1.DeleteUser)

		tag := user.Group("/:userId/tag")
		{
			tag.Post("/", v1.CreateTag)
			tag.Get("/", v1.GetAllTags)
			tag.Get("/:id", v1.GetOneTag)
			tag.Patch("/:id", v1.UpdateTag)
			tag.Delete("/:id", v1.DeleteTag)
		}

		permission := user.Group("/:userId/permission")
		{
			permission.Post("/", v1.CreatePermission)
			permission.Get("/", v1.GetAllPermissions)
			permission.Get("/:id", v1.GetOnePermission)
			permission.Patch("/:id", v1.UpdatePermission)
			permission.Delete("/:id", v1.DeletePermission)
		}
	}
}
