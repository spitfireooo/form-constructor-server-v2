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

		userTag := user.Group("/:userId/tag")
		{
			userTag.Post("/", controller_v1.CreateUserTag)
			userTag.Get("/", controller_v1.GetUserTags)
			userTag.Delete("/", controller_v1.DeleteUserTags)
			userTag.Get("/:id", controller_v1.GetUserTag)
			userTag.Patch("/:id", controller_v1.UpdateUserTag)
			userTag.Delete("/:id", controller_v1.DeleteUserTag)
		}

		userPermission := user.Group("/:userId/permission")
		{
			userPermission.Post("/", controller_v1.CreatePermission)
			userPermission.Get("/", controller_v1.GetUserPermissions)
			userPermission.Delete("/", controller_v1.DeleteUserPermissions)
		}
	}

	group.Get("/user-tag", controller_v1.GetAllUserTags)
	tag := group.Group("/tag")
	{
		tag.Post("/", controller_v1.CreateTag)
		tag.Get("/", controller_v1.GetAllTags)
		tag.Get("/:id", controller_v1.GetOneTag)
		tag.Patch("/:id", controller_v1.UpdateTag)
		tag.Delete("/:id", controller_v1.DeleteTag)
	}

	permission := group.Group("/permission")
	{
		permission.Get("/", controller_v1.GetAllPermissions)
		permission.Get("/:id", controller_v1.GetOnePermission)
		permission.Patch("/:id", controller_v1.UpdatePermission)
		permission.Delete("/:id", controller_v1.DeletePermission)
	}
}
