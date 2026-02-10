package main

import (
	"log"

	"github.com/Gylmynnn/chat-cuy-be/internal/config"
	"github.com/Gylmynnn/chat-cuy-be/internal/handler"
	"github.com/Gylmynnn/chat-cuy-be/internal/repository"
	"github.com/Gylmynnn/chat-cuy-be/internal/routes"
	"github.com/Gylmynnn/chat-cuy-be/internal/service"
	"github.com/Gylmynnn/chat-cuy-be/internal/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	// Load configuration
	cfg := config.LoadConfig()

	// Connect to database
	db := cfg.ConnectDB()
	defer db.Close()

	// Initialize repositories
	userRepo := repository.NewUserRepository(db)
	chatRoomRepo := repository.NewChatRoomRepository(db)
	messageRepo := repository.NewMessageRepository(db)
	contactRepo := repository.NewContactRepository(db)
	storyRepo := repository.NewStoryRepository(db)

	// Initialize services
	userService := service.NewUserService(userRepo)
	chatRoomService := service.NewChatRoomService(chatRoomRepo)
	messageService := service.NewMessageService(messageRepo)
	contactService := service.NewContactService(contactRepo, userRepo)
	storyService := service.NewStoryService(storyRepo)

	// Initialize handlers
	userHandler := handler.NewUserHandler(userService, cfg)
	chatRoomHandler := handler.NewChatRoomHandler(chatRoomService)
	messageHandler := handler.NewMessageHandler(messageService)
	contactHandler := handler.NewContactHandler(contactService)
	storyHandler := handler.NewStoryHandler(storyService)

	// Initialize WebSocket
	hub := websocket.NewHub()
	go hub.Run()
	wsHandler := websocket.NewWSHandler(hub, messageService)

	// Create Fiber app
	app := fiber.New(fiber.Config{
		AppName:   "Chat Cuy API",
		BodyLimit: 10 * 1024 * 1024, // 10MB limit for file uploads
	})

	// Middleware
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET, POST, PUT, DELETE, OPTIONS",
	}))

	// Serve uploaded files as static
	app.Static("/uploads", "./uploads")

	// Setup routes
	routes.SetupRoutes(app, userHandler, chatRoomHandler, messageHandler, contactHandler, storyHandler, wsHandler, cfg.JWTSecret)

	// Health check
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok"})
	})

	// Start server
	log.Printf("Server starting on port %s", cfg.AppPort)
	if err := app.Listen(":" + cfg.AppPort); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
