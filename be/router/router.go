package router

import (
	"github.com/Gylmynnn/websocket-sesat/internal/user"
	"github.com/Gylmynnn/websocket-sesat/internal/websocket"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

var app *gin.Engine

func InitRouter(userHandler *user.Handler, wsHandler *websocket.Handler) {
	app = gin.Default()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "http://localhost:5173"
		},
		MaxAge: 12 * time.Hour,
	}))
	app.POST("/signup", userHandler.CreateUser)
	app.POST("/login", userHandler.Login)
	app.POST("/withgoogle", userHandler.LoginWithGoogle)
	app.POST("/withfacebook", userHandler.LoginWithFacebook)
	app.GET("/logout", userHandler.Logout)

	app.POST("/ws/createRoom", wsHandler.CreateRoom)
	app.GET("/ws/joinRoom/:roomId", wsHandler.JoinRoom)
	app.GET("/ws/getRooms", wsHandler.GetRooms)
	app.GET("/ws/getClients/:roomId", wsHandler.GetClient)
}

func Start(addr string) error {
	return app.Run(addr)
}
