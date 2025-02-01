package router

import (
	"github.com/Gylmynnn/websocket-sesat/internal/user"
	"github.com/Gylmynnn/websocket-sesat/internal/websocket"
	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func InitRouter(userHandler *user.Handler, wsHandler *websocket.Handler) {
	r = gin.Default()
	r.POST("/signup", userHandler.CreateUser)
	r.POST("/login", userHandler.Login)
	r.GET("/logout", userHandler.Logout)

	r.POST("/ws/createRoom", wsHandler.CreateRoom)
	r.GET("/ws/joinRoom/:roomId", wsHandler.JoinRoom)
	r.GET("/ws/getRooms", wsHandler.GetRooms)
	r.GET("/ws/getClients/:roomId", wsHandler.GetClient)
}

func Start(addr string) error {
	return r.Run(addr)
}
