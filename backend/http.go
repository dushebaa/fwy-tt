package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func listenHttp(sc *Scanner) {
	app := fiber.New()
	app.Use(cors.New())

	baseGroup := app.Group("/api")

	baseGroup.Get("/created-collections", func(c *fiber.Ctx) error {
		return c.JSON(sc.collectionCreatedLogs)
	})

	app.Listen(":8080")
}
