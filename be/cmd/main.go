package main

import (
	"fmt"
	"log"

	"github.com/Gylmynnn/websocket-sesat/database"
	"github.com/Gylmynnn/websocket-sesat/internal/user"
	"github.com/Gylmynnn/websocket-sesat/internal/websocket"
	"github.com/Gylmynnn/websocket-sesat/router"
)

func main() {
	dbConnection, err := database.NewDatabaseConn()
	if err != nil {
		log.Fatalf("error when initial database %s", err)
    }


	userRepository := user.NewRepository(dbConnection.GetDB())
	userService := user.NewService(userRepository)
	userHandler := user.NewHundler(userService)

	hub := websocket.NewHub()
	wsHandler := websocket.NewHandler(hub)
   go hub.Run()

	router.InitRouter(userHandler, wsHandler)
	router.Start("0.0.0.0:8080")
	fmt.Println("server running on port : 8080")
}
