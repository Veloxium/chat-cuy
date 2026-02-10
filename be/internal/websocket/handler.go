package websocket

import (
	"context"
	"encoding/json"
	"log"

	"github.com/Gylmynnn/chat-cuy-be/internal/service"
	ws "github.com/gofiber/contrib/websocket"
)

type WSHandler struct {
	hub            *Hub
	messageService service.MessageService
}

func NewWSHandler(hub *Hub, messageService service.MessageService) *WSHandler {
	return &WSHandler{
		hub:            hub,
		messageService: messageService,
	}
}

func (h *WSHandler) HandleWebSocket(c *ws.Conn) {
	userID := c.Locals("userID")
	if userID == nil {
		log.Println("WebSocket connection without userID")
		c.Close()
		return
	}

	uid, ok := userID.(int)
	if !ok {
		log.Println("WebSocket userID is not an int")
		c.Close()
		return
	}

	client := &Client{
		Conn:   c,
		UserID: uid,
		Rooms:  make(map[int]bool),
	}

	h.hub.Register() <- client

	defer func() {
		h.hub.Unregister() <- client
		c.Close()
	}()

	for {
		_, msgBytes, err := c.ReadMessage()
		if err != nil {
			log.Printf("WebSocket read error for user %d: %v", uid, err)
			break
		}

		var wsMsg WSMessage
		if err := json.Unmarshal(msgBytes, &wsMsg); err != nil {
			log.Printf("Invalid WebSocket message format: %v", err)
			continue
		}

		h.handleMessage(uid, &wsMsg)
	}
}

func (h *WSHandler) handleMessage(userID int, msg *WSMessage) {
	switch msg.Type {
	case "joinRoom":
		var payload JoinRoomPayload
		if err := json.Unmarshal(msg.Payload, &payload); err != nil {
			log.Printf("Invalid joinRoom payload: %v", err)
			return
		}
		log.Printf("User %d registered via joinRoom", payload.UserID)

	case "enterChatRoom":
		var payload EnterChatRoomPayload
		if err := json.Unmarshal(msg.Payload, &payload); err != nil {
			log.Printf("Invalid enterChatRoom payload: %v", err)
			return
		}
		h.hub.JoinRoom(userID, payload.RoomID)

	case "sendMessage":
		var payload SendMessagePayload
		if err := json.Unmarshal(msg.Payload, &payload); err != nil {
			log.Printf("Invalid sendMessage payload: %v", err)
			return
		}

		h.hub.JoinRoom(userID, payload.RoomID)

		msg, err := h.messageService.SendMessage(
			context.Background(),
			payload.UserID,
			payload.RoomID,
			payload.Content,
		)
		if err != nil {
			log.Printf("Failed to save message: %v", err)
			return
		}

		h.hub.BroadcastToRoom(payload.RoomID, "receiveMessage", msg)

	default:
		log.Printf("Unknown message type: %s", msg.Type)
	}
}
