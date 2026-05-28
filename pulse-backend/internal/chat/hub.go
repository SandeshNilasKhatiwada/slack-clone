package chat

import (
	"log"

	"github.com/gofiber/contrib/websocket"
)

// ChatMessage represents the JSON structure of our messages
type ChatMessage struct {
	Text string `json:"text"`
}

// Hub maintains the set of active clients and broadcasts messages to them.
type Hub struct {
	// Registered clients (using a map as a fast lookup set)
	Clients map[*websocket.Conn]bool

	// Inbound messages from the clients
	Broadcast chan ChatMessage

	// Register requests from the clients
	Register chan *websocket.Conn

	// Unregister requests from clients
	Unregister chan *websocket.Conn
}

// Global instance of our Hub
var ChatHub = Hub{
	Broadcast:  make(chan ChatMessage),
	Register:   make(chan *websocket.Conn),
	Unregister: make(chan *websocket.Conn),
	Clients:    make(map[*websocket.Conn]bool),
}

// Run starts the infinite loop that listens to our channels
func (h *Hub) Run() {
	for {
		// 'select' waits until one of these channels receives data
		select {
		case conn := <-h.Register:
			h.Clients[conn] = true
			log.Println("User joined! Total connected:", len(h.Clients))

		case conn := <-h.Unregister:
			if _, ok := h.Clients[conn]; ok {
				delete(h.Clients, conn)
				conn.Close()
				log.Println("User left! Total connected:", len(h.Clients))
			}

		case message := <-h.Broadcast:
			// Send the message to ALL connected clients
			for conn := range h.Clients {
				if err := conn.WriteJSON(message); err != nil {
					log.Println("Error broadcasting:", err)
					// If a connection is dead, unregister it
					conn.Close()
					delete(h.Clients, conn)
				}
			}
		}
	}
}
