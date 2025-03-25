package router

import (
	"database/sql"
	"github.com/Gylmynnn/websocket-sesat/internal/contact"
	"github.com/Gylmynnn/websocket-sesat/internal/user"
	"github.com/Gylmynnn/websocket-sesat/internal/websocket"
	"github.com/Gylmynnn/websocket-sesat/protected"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

// initialzed gin engine
var app *gin.Engine

// initialzed router function
func InitRouter(db *sql.DB, userHandler *user.Handler, wsHandler *websocket.Handler, contactHandler *contact.Handler) {
	app = gin.Default()

	// wrap route with logger
	app.Use(protected.Logger(db))

	// app cors
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

	// user authentication route handler
	app.POST("/api/register", userHandler.CreateUser)
	app.POST("/api/login/default", userHandler.Login)
	app.POST("/api/login/google", userHandler.LoginWithGoogle)
	app.POST("/api/login/facebook", userHandler.LoginWithFacebook)
	app.GET("/api/logout", userHandler.Logout)

	// protected route with jwt
	authApp := app.Group("/")
	authApp.Use(protected.JWTAuthMiddleware())
	// contact route handler
	authApp.POST("/api/contact", contactHandler.AddContact)
	authApp.PUT("/api/contact/:id", contactHandler.DeleteContact)
	authApp.GET("/api/contact/:id", contactHandler.GetContactByID)
	authApp.GET("/api/contacts", contactHandler.GetAllContacts)

	// websocket route handler
	authApp.POST("/ws/createRoom", wsHandler.CreateRoom)
	authApp.GET("/ws/joinRoom/:roomId", wsHandler.JoinRoom)
	authApp.GET("/ws/getRooms", wsHandler.GetRooms)
	authApp.GET("/ws/getClients/:roomId", wsHandler.GetClient)
}

func Start(addr string) error {
	return app.Run(addr)
}
