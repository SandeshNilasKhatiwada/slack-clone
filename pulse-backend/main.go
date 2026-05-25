package main

import (
	"github.com/SandeshNilasKhatiwada/slack-clone/internal/database"
	"github.com/SandeshNilasKhatiwada/slack-clone/internal/middleware"

	"github.com/SandeshNilasKhatiwada/slack-clone/internal/auth"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	//connect to the database
	data, err := database.Connect()
	// Create a GET route for our health check
	// Notice how '*fiber.Ctx' acts exactly like the (req, res) object in Express
	app.Get("/api/health", func(c *fiber.Ctx) error {

		// databse connection check
		if err != nil {
			return c.JSON(fiber.Map{
				"status":  "error",
				"message": "Database connection failed",
				"error":   err.Error(),
			})
		}
		_ = data // this means ignore this variable, we will use it later when we implement the database operations
		return c.JSON(fiber.Map{
			"status":  "ok",
			"message": "Database connection successful",
		})
	})
	
	app.Post("/api/user/register", auth.Register)
	app.Post("/api/user/login", auth.Login)
	
	user := app.Group("/api/user", middleware.RequireAuth)
	user.Get("/me", auth.GetMe)


	// Start the server on port 8080
	app.Listen(":8080")

}
