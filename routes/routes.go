package routes

import (
	v1_routes "main/routes/v1"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func StartService() {
	app := fiber.New()

	api := app.Group("/api")

	app.Use(func(c *fiber.Ctx) error {
		c.Set("Access-Control-Allow-Origin", "http://localhost:5173")
		c.Set("Access-Control-Allow-Headers", "Content-Type")
		c.Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		return c.Next()
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:5173",
		AllowCredentials: true,
	}))

	v1_routes.StartService(api)

	app.Listen(":3000")
}
