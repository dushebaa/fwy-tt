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

	baseGroup.Get("/tokens-minted/:address", func(c *fiber.Ctx) error {
		address := c.Params("address")

		if sc.tokensMintedLogs[address] == nil {
			return c.JSON([]LogTokenMinted{})
		}

		return c.JSON(sc.tokensMintedLogs[address])
	})

	app.Listen(":8080")
}
