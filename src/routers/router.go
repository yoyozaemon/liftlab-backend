package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetUpRouter() *fiber.App {
	app := fiber.New()

	app.Use(logger.New())
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("liftlab is working")
	})

	app.Use(cors.New())

	SetupRoutes(app)

	return app
}
