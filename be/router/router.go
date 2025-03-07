package router

import (
	"time"

	"github.com/Gylmynnn/websocket-sesat/internal/contact"
	"github.com/Gylmynnn/websocket-sesat/internal/user"
	"github.com/Gylmynnn/websocket-sesat/internal/websocket"
	"github.com/Gylmynnn/websocket-sesat/protected"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var app *gin.Engine

func InitRouter(userHandler *user.Handler, wsHandler *websocket.Handler, contactHandler *contact.Handler) {
	app = gin.Default()
	app.Use(protected.Logger())

	app.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
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

	authApp := app.Group("/")
	authApp.Use(protected.JWTAuthMiddleware())

	authApp.POST("/contact", contactHandler.AddContact)

	authApp.POST("/ws/createRoom", wsHandler.CreateRoom)
	authApp.GET("/ws/joinRoom/:roomId", wsHandler.JoinRoom)
	authApp.GET("/ws/getRooms", wsHandler.GetRooms)
	authApp.GET("/ws/getClients/:roomId", wsHandler.GetClient)
}

func Start(addr string) error {
	return app.Run(addr)
}
