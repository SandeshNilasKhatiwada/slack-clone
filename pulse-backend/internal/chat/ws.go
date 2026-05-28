package chat

import (
	"log"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

func UpgradeToWebSocket(c *fiber.Ctx) error {
	if websocket.IsWebSocketUpgrade(c) {
		c.Locals("allowed", true)
		return c.Next()
	}
	return fiber.ErrUpgradeRequired
}

var HandleChat = websocket.New(func(c *websocket.Conn) {
	// When a client connects, log it
	log.Println("New WebSocket connection established!")

	// Create an infinite loop to keep the connection alive and listen for messages
	for {
		var msg map[string]interface{}

		// 1. Read the message from the Angular client
		err := c.ReadJSON(&msg)
		if err != nil {
			log.Println("Client disconnected or error:", err)
			break // Break the loop if the client disconnects
		}

		log.Printf("Received message from Angular: %v\n", msg)

		// 2. Echo the message back to the client to prove it works
		response := map[string]string{
			"status": "Message received loud and clear!",
		}

		if err := c.WriteJSON(response); err != nil {
			log.Println("Error writing message:", err)
			break
		}
	}
})
