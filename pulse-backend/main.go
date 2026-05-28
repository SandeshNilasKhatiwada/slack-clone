package main

import (
	"github.com/SandeshNilasKhatiwada/slack-clone/internal/auth"
	"github.com/SandeshNilasKhatiwada/slack-clone/internal/chat"
	"github.com/SandeshNilasKhatiwada/slack-clone/internal/database"
	"github.com/SandeshNilasKhatiwada/slack-clone/internal/middleware"
	"github.com/SandeshNilasKhatiwada/slack-clone/internal/repository"
	"github.com/SandeshNilasKhatiwada/slack-clone/internal/service"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	//adding cors
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:4200",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))
	//connect to the database
	data, err := database.Connect()

	// wire up repository, service, and handler
	repo := repository.NewUserRepository(data)
	svc := service.NewAuthService(repo)
	handler := auth.NewAuthHandler(svc)

	go chat.ChatHub.Run() // Start the chat hub in a separate goroutine
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

	app.Post("/api/user/register", handler.Register)
	app.Post("/api/user/login", handler.Login)

	user := app.Group("/api/user", middleware.RequireAuth)
	user.Get("/me", handler.GetMe)

	// --- WebSocket Routes ---
	// We use the Upgrade middleware first, then pass it to our WebSocket handler
	app.Use("/ws", chat.UpgradeToWebSocket)
	app.Get("/ws/chat", chat.HandleChat)

	// Start the server on port 8080
	app.Listen(":8080")

}
