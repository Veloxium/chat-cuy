package chat

import (
	"context"
	"log"
	"net/http"
	"sync"

	"github.com/Gylmynnn/websocket-sesat/utils"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type Handler struct {
	Service Service
	clients map[*websocket.Conn]string
	mu      sync.Mutex
}

func NewHandler(s Service) *Handler {
	return &Handler{
		Service: s,
		clients: make(map[*websocket.Conn]string),
	}
}

func (h *Handler) AddChat(c *gin.Context) {
	var chat CreateChatReq
	if err := c.ShouldBindJSON(&chat); err != nil {
		c.JSON(http.StatusBadRequest, utils.ResFormatter{
			Success:    false,
			StatusCode: http.StatusBadRequest,
			Message:    "Error :" + err.Error(),
			Data:       nil,
		})
	}

	res, err := h.Service.CreateChat(c.Request.Context(), &chat)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ResFormatter{
			Success:    false,
			StatusCode: http.StatusInternalServerError,
			Message:    "Error :" + err.Error(),
			Data:       nil,
		})
		return
	}

	c.JSON(http.StatusOK, utils.ResFormatter{
		Success:    true,
		StatusCode: http.StatusOK,
		Message:    "Chat room create successfully",
		Data:       res,
	})

}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Bisa diperbaiki agar lebih aman
	},
}

// Handle koneksi WebSocket
func (h *Handler) HandleWebsocket(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("WebSocket Upgrade Error:", err)
		return
	}

	userID := c.Query("user_id")
	chatID := c.Query("chat_id")

	h.mu.Lock()
	h.clients[conn] = userID
	h.mu.Unlock()

	log.Println("Client connected:", userID)

	go h.readMessages(c.Request.Context(), conn, chatID)

}

func (h *Handler) readMessages(ctx context.Context, conn *websocket.Conn, chatID string) {
	defer func() {
		h.mu.Lock()
		delete(h.clients, conn)
		h.mu.Unlock()
		conn.Close()
		log.Println("Client disconnected:", conn.RemoteAddr())
	}()

	for {

		var msg CreateMessageReq
		err := conn.ReadJSON(&msg)
		if err != nil {
			log.Println("Read error:", err)
			break
		}
		msg.ChatID = chatID
		if msg.Content == "" {
			log.Println("Empty message received, skipping")
			continue
		}

		res, err := h.Service.CreateMessage(ctx, &msg)
		if err != nil {
			log.Println("Database insert error:", err)
			continue
		}
		h.broadcastMessage(res)
	}
}

func (h *Handler) broadcastMessage(msg *CreateMessageRes) {
	h.mu.Lock()
	defer h.mu.Unlock()
	for client, userID := range h.clients {
		if userID != msg.SenderID {
			err := client.WriteJSON(&msg)
			if err != nil {
				log.Println("Write error:", err)
				client.Close()
				delete(h.clients, client)
			}
		}
	}
}
