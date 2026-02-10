package websocket

import (
	"encoding/json"
	"log"
	"sync"

	ws "github.com/gofiber/contrib/websocket"
)

type Client struct {
	Conn   *ws.Conn
	UserID int
	Rooms  map[int]bool
	mu     sync.Mutex
}

type WSMessage struct {
	Type    string          `json:"type"`
	Payload json.RawMessage `json:"payload"`
}

type JoinRoomPayload struct {
	UserID int `json:"userId"`
}

type EnterChatRoomPayload struct {
	RoomID int `json:"roomId"`
	UserID int `json:"userId"`
}

type SendMessagePayload struct {
	RoomID  int    `json:"roomId"`
	Content string `json:"content"`
	UserID  int    `json:"userId"`
}

type Hub struct {
	clients    map[int]*Client      // userID -> Client
	rooms      map[int]map[int]bool // roomID -> set of userIDs
	register   chan *Client
	unregister chan *Client
	broadcast  chan BroadcastMessage
	mu         sync.RWMutex
}

type BroadcastMessage struct {
	RoomID  int
	Message interface{}
	Type    string
}

func NewHub() *Hub {
	return &Hub{
		clients:    make(map[int]*Client),
		rooms:      make(map[int]map[int]bool),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		broadcast:  make(chan BroadcastMessage, 256),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.mu.Lock()
			h.clients[client.UserID] = client
			h.mu.Unlock()
			log.Printf("Client connected: userID=%d", client.UserID)

		case client := <-h.unregister:
			h.mu.Lock()
			if _, ok := h.clients[client.UserID]; ok {
				delete(h.clients, client.UserID)
				for roomID := range client.Rooms {
					if h.rooms[roomID] != nil {
						delete(h.rooms[roomID], client.UserID)
					}
				}
			}
			h.mu.Unlock()
			log.Printf("Client disconnected: userID=%d", client.UserID)

		case msg := <-h.broadcast:
			h.mu.RLock()
			if users, ok := h.rooms[msg.RoomID]; ok {
				response := map[string]interface{}{
					"type":    msg.Type,
					"payload": msg.Message,
				}
				data, err := json.Marshal(response)
				if err != nil {
					h.mu.RUnlock()
					continue
				}
				for userID := range users {
					if client, ok := h.clients[userID]; ok {
						client.mu.Lock()
						err := client.Conn.WriteMessage(ws.TextMessage, data)
						client.mu.Unlock()
						if err != nil {
							log.Printf("Error sending message to user %d: %v", userID, err)
						}
					}
				}
			}
			h.mu.RUnlock()
		}
	}
}

func (h *Hub) JoinRoom(userID, roomID int) {
	h.mu.Lock()
	defer h.mu.Unlock()

	if h.rooms[roomID] == nil {
		h.rooms[roomID] = make(map[int]bool)
	}
	h.rooms[roomID][userID] = true

	if client, ok := h.clients[userID]; ok {
		client.Rooms[roomID] = true
	}

	log.Printf("User %d joined room %d", userID, roomID)
}

func (h *Hub) BroadcastToRoom(roomID int, msgType string, message interface{}) {
	h.broadcast <- BroadcastMessage{
		RoomID:  roomID,
		Message: message,
		Type:    msgType,
	}
}

func (h *Hub) Register() chan<- *Client {
	return h.register
}

func (h *Hub) Unregister() chan<- *Client {
	return h.unregister
}
