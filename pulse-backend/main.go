package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	// Create a GET route for our health check
	// Notice how '*fiber.Ctx' acts exactly like the (req, res) object in Express
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "ok",
		})
	})

	// Start the server on port 8080
	app.Listen(":8080")

}
