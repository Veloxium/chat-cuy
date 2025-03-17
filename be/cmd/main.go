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
	getDB := dbConn.GetDB()
	database.InitFirebase()

	userRepository := user.NewRepository(getDB)
	userService := user.NewService(userRepository)
	userHandler := user.NewHundler(userService)

	contactRepository := contact.NewRepository(getDB)
	contactService := contact.NewService(contactRepository)
	contactHandler := contact.NewHundler(contactService)

	hub := websocket.NewHub()
	wsHandler := websocket.NewHandler(hub)
	go hub.Run()

	router.InitRouter(getDB, userHandler, wsHandler, contactHandler)
	if err := router.Start("localhost:8080"); err != nil {
		fmt.Println("error when running server")
	}
	fmt.Println("server running on port : 8080")

}
