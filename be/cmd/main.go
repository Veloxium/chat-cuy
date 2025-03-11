package main

import (
	"fmt"
	"github.com/Gylmynnn/websocket-sesat/database"
	"github.com/Gylmynnn/websocket-sesat/internal/contact"
	"github.com/Gylmynnn/websocket-sesat/internal/user"
	"github.com/Gylmynnn/websocket-sesat/internal/websocket"
	"github.com/Gylmynnn/websocket-sesat/router"
)

func main() {

	dbConn := database.NewDatabaseConn()
	database.InitFirebase()

	userRepository := user.NewRepository(dbConn.GetDB())
	userService := user.NewService(userRepository)
	userHandler := user.NewHundler(userService)

	contactRepository := contact.NewRepository(dbConn.GetDB())
	contactService := contact.NewService(contactRepository)
	contactHandler := contact.NewHundler(contactService)

	hub := websocket.NewHub()
	wsHandler := websocket.NewHandler(hub)
	go hub.Run()

	router.InitRouter(dbConn.GetDB(), userHandler, wsHandler, contactHandler)
	router.Start("0.0.0.0:8080")
	fmt.Println("server running on port : 8080")

}
