package websocket

import (
	"net/http"

	"github.com/Gylmynnn/websocket-sesat/utils"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type Handler struct {
	hub *Hub
}

func NewHandler(h *Hub) *Handler {
	return &Handler{
		hub: h,
	}
}

type CreateRoomReq struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (h *Handler) CreateRoom(c *gin.Context) {
	var req CreateRoomReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, utils.ResFormatter{
			Success:    false,
			StatusCode: http.StatusBadRequest,
			Message:    "error :" + err.Error(),
			Data:       nil,
		})
		return
	}

	h.hub.Rooms[req.ID] = &Room{
		ID:      req.ID,
		Name:    req.Name,
		Clients: make(map[string]*Client),
	}

	c.JSON(http.StatusOK, utils.ResFormatter{
		Success:    true,
		StatusCode: http.StatusOK,
		Message:    "create room successfully",
		Data:       req,
	})

}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		// _ := r.Header.Get("Origin")
		return true
	},
}

func (h *Handler) JoinRoom(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ResFormatter{
			Success:    false,
			StatusCode: http.StatusBadRequest,
			Message:    "error :" + err.Error(),
			Data:       nil,
		})
		return
	}

	roomId := c.Param("roomId")
	clientId := c.Query("userId")
	username := c.Query("username")

	client := &Client{
		Conn:     conn,
		Message:  make(chan *Message, 10),
		ID:       clientId,
		RoomID:   roomId,
		Username: username,
	}

	message := &Message{
		Content:  "A new user has join room",
		RoomID:   roomId,
		Username: username,
	}

	h.hub.Register <- client
	h.hub.Broadcast <- message

	go client.writeMessage()
	client.readMessage(h.hub)
}

type RoomsRes struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (h *Handler) GetRooms(c *gin.Context) {
	rooms := make([]RoomsRes, 0)

	for _, r := range h.hub.Rooms {
		rooms = append(rooms, RoomsRes{
			ID:   r.ID,
			Name: r.Name,
		})
	}

	c.JSON(http.StatusOK, utils.ResFormatter{
		Success:    true,
		StatusCode: http.StatusOK,
		Message:    "get rooms successfully",
		Data:       rooms,
	})

}

type ClientsRes struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}

func (h *Handler) GetClient(c *gin.Context) {
	var clients []ClientsRes
	roomId := c.Param("roomId")

	if _, ok := h.hub.Rooms[roomId]; !ok {
		clients = make([]ClientsRes, 0)
		c.JSON(http.StatusOK, utils.ResFormatter{
			Success:    true,
			StatusCode: http.StatusOK,
			Message:    "get clients successfully",
			Data:       clients,
		})
	}

	for _, c := range h.hub.Rooms[roomId].Clients {
		clients = append(clients, ClientsRes{
			ID:       c.ID,
			Username: c.Username,
		})
	}

	c.JSON(http.StatusOK, utils.ResFormatter{
		Success:    true,
		StatusCode: http.StatusOK,
		Message:    "get clients successfully",
		Data:       clients,
	})
}
