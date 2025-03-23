package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/spitfireooo/form-constructor-server-v2/pkg/controller/v1"
)

func UserRouter(group fiber.Router, path string) {
	user := group.Group(path)
	{
		user.Get("/", controller_v1.GetAllUsers)
		user.Get("/:id", controller_v1.GetOneUser)
		user.Get("/:id", controller_v1.UpdateUser)
		user.Get("/:id", controller_v1.DeleteUser)

		tag := user.Group("/:userId/tag")
		{
			tag.Post("/", controller_v1.CreateTag)
			tag.Get("/", controller_v1.GetAllTags)
			tag.Get("/:id", controller_v1.GetOneTag)
			tag.Patch("/:id", controller_v1.UpdateTag)
			tag.Delete("/:id", controller_v1.DeleteTag)
		}

		permission := user.Group("/:userId/permission")
		{
			permission.Post("/", controller_v1.CreatePermission)
			permission.Get("/", controller_v1.GetAllPermissions)
			permission.Get("/:id", controller_v1.GetOnePermission)
			permission.Patch("/:id", controller_v1.UpdatePermission)
			permission.Delete("/:id", controller_v1.DeletePermission)
		}
	}
}
