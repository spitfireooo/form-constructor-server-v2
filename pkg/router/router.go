package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	router_v2 "github.com/spitfireooo/form-constructor-server-v2/pkg/router/v2"
	"time"
)

func Router(r *fiber.App) {
	api := r.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			UserRouter(v1, "/user")
			FormRouter(v1, "/form")
			FieldRouter(v1, "/field")
			ResultRouter(v1, "/result")
		}

		v2 := api.Group("/v2")
		{
			router_v2.FormRouter(v2, "/form")
		}
	}

	utils := r.Group("/utils")
	{
		utils.Get("/metrics", monitor.New(monitor.Config{
			Title:   "Metrics Of Form Constructor Server",
			Refresh: time.Second,
		}))
	}
}
