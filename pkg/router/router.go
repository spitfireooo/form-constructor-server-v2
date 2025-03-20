package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"time"
)

func Router(r *fiber.App) {
	api := r.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			v1.Get("/test", func(ctx *fiber.Ctx) error {
				return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
					"message": "OK",
				})
			})
		}

		v2 := api.Group("/v2")
		{
			v2.Get("/test2", func(ctx *fiber.Ctx) error {
				return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
					"message": "OK",
				})
			})
		}
	}

	utils := r.Group("/utils")
	{
		utils.Get("metrics", monitor.New(monitor.Config{
			Title:   "Metrics Of Form Constructor Server",
			Refresh: time.Second,
		}))
	}
}
