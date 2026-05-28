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

// HandleChat upgrades the HTTP connection and connects the user to the Hub
var HandleChat = websocket.New(func(c *websocket.Conn) {
	// 1. Register this new connection to the Hub
	ChatHub.Register <- c

	// 2. Ensure we unregister the user when they disconnect
	defer func() {
		ChatHub.Unregister <- c
	}()

	// 3. Listen for incoming messages from THIS specific client
	for {
		var msg ChatMessage
		err := c.ReadJSON(&msg)
		if err != nil {
			log.Println("Client disconnected or error:", err)
			break // Break the loop, which triggers the defer func above
		}

		// 4. Send the message to the Hub's broadcast channel
		ChatHub.Broadcast <- msg
	}
})
