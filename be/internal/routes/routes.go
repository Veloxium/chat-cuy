package routes

import (
	"github.com/Gylmynnn/chat-cuy-be/internal/handler"
	"github.com/Gylmynnn/chat-cuy-be/internal/middleware"
	wsHandler "github.com/Gylmynnn/chat-cuy-be/internal/websocket"
	"github.com/Gylmynnn/chat-cuy-be/pkg/utils"
	ws "github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(
	app *fiber.App,
	userHandler *handler.UserHandler,
	chatRoomHandler *handler.ChatRoomHandler,
	messageHandler *handler.MessageHandler,
	contactHandler *handler.ContactHandler,
	storyHandler *handler.StoryHandler,
	wsH *wsHandler.WSHandler,
	jwtSecret string,
) {
	api := app.Group("/api")

	// Public routes
	api.Post("/login", userHandler.Login)
	api.Post("/register", userHandler.Register)

	// Protected routes
	protected := api.Group("", middleware.AuthMiddleware(jwtSecret))
	protected.Get("/profile", userHandler.GetProfile)
	protected.Put("/profile/avatar", userHandler.UploadAvatar)
	protected.Get("/users/search", userHandler.SearchUsers)

	// Chat rooms
	protected.Get("/chatrooms/:userId", chatRoomHandler.GetUserChatRooms)
	protected.Post("/chatrooms", chatRoomHandler.CreatePrivateRoom)

	// Messages
	protected.Get("/messages/:roomId", messageHandler.GetMessages)

	// Contacts
	protected.Get("/contacts", contactHandler.GetContacts)
	protected.Post("/contacts", contactHandler.AddContact)

	// Stories
	protected.Post("/stories", storyHandler.CreateStory)
	protected.Get("/stories", storyHandler.GetStories)
	protected.Get("/stories/me", storyHandler.GetMyStories)
	protected.Delete("/stories/:id", storyHandler.DeleteStory)

	// WebSocket - with JWT auth via query param
	app.Use("/ws", func(c *fiber.Ctx) error {
		if ws.IsWebSocketUpgrade(c) {
			token := c.Query("token")
			if token != "" {
				claims, err := utils.ValidateJWT(jwtSecret, token)
				if err == nil {
					c.Locals("userID", claims.UserID)
					c.Locals("username", claims.Username)
				}
			}
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	app.Get("/ws", ws.New(wsH.HandleWebSocket))
}
