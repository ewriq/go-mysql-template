package Routers

import (
	"template/Handler"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Initalize(app *fiber.App){
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173",
		AllowHeaders:  "Origin, Content-Type, Accept",
		AllowCredentials: true,
	}))
	app.Use(logger.New(logger.Config{
		Format: "${ip}:${port} ${status} - ${method} ${path}\n",
	}))

    app.Get("/", Handlers.Home)
	
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(404).JSON(fiber.Map{
			"code":    404,
			"message": "404: Not Found",
		})
	})

}