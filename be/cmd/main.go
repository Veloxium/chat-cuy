package main

import (
	"fmt"
	"github.com/Gylmynnn/websocket-sesat/database"
	"github.com/Gylmynnn/websocket-sesat/internal/chat"
	"github.com/Gylmynnn/websocket-sesat/internal/contact"
	"github.com/Gylmynnn/websocket-sesat/internal/user"
	"github.com/Gylmynnn/websocket-sesat/internal/websocket"
	"github.com/Gylmynnn/websocket-sesat/router"
	"github.com/Gylmynnn/websocket-sesat/utils"
)

func main() {

	databaseConn := database.NewDatabaseConn()
	getDatabase := databaseConn.GetDB()

	err := databaseConn.Ping()
	utils.HandleErr("connecting to database failed", err)

	err = database.InitFirebase()
	utils.HandleErr("initialize firebase failed", err)

	userRepository := user.NewRepository(getDatabase)
	userService := user.NewService(userRepository)
	userHandler := user.NewHundler(userService)

	contactRepository := contact.NewRepository(getDatabase)
	contactService := contact.NewService(contactRepository)
	contactHandler := contact.NewHundler(contactService)

	chatRepository := chat.NewRepository(getDatabase)
	chatService := chat.NewService(chatRepository)
	chatHandler := chat.NewHandler(chatService)

	hub := websocket.NewHub()
	wsHandler := websocket.NewHandler(hub)
	go hub.Run()

	router.InitRouter(
		getDatabase,
		userHandler,
		wsHandler,
		contactHandler,
		chatHandler)

	err = router.Start("localhost:8080")
	utils.HandleErr("running server :", err)
	fmt.Println("server running on port : 8080")

}
