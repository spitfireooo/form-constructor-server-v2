package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/spitfireooo/form-constructor-server-v2/pkg/controller/v1"
	"github.com/spitfireooo/form-constructor-server-v2/pkg/service"
)

func UserRouter(group fiber.Router, path string) {
	user := group.Group(path)
	{
		user.Get("/", controller_v1.GetAllUsers)
		user.Get("/:id", controller_v1.GetOneUser)
		user.Patch("/:id", controller_v1.UpdateUser)
		user.Delete("/:id", controller_v1.DeleteUser)
		user.Post("/upload", service.FileUpload)

		tag := user.Group("/:userId/tag")
		{
			tag.Post("/", controller_v1.CreateTag)
			tag.Get("/", controller_v1.GetAllTags)
			tag.Get("/:id", controller_v1.GetOneTag)
			tag.Patch("/:id", controller_v1.UpdateTag)
			tag.Delete("/:id", controller_v1.DeleteTag)
		}

		userPermission := user.Group("/:userId/permission")
		{
			userPermission.Post("/", controller_v1.CreatePermission)
			userPermission.Get("/", controller_v1.GetUserPermissions)
			userPermission.Delete("/", controller_v1.DeleteUserPermissions)
		}
	}

	permission := group.Group("/permission")
	{
		permission.Get("/", controller_v1.GetAllPermissions)
		permission.Get("/:id", controller_v1.GetOnePermission)
		permission.Patch("/:id", controller_v1.UpdatePermission)
		permission.Delete("/:id", controller_v1.DeletePermission)
	}
}
