package main

import (
	"math/rand"
	"net/http"
	"tic/game"
	"tic/handlers"

	"tic/models"
	"tic/nicknames"
	"time"
)

var GameMemory models.GameMemory
var Nickname map[string]bool

func main() {
	rand.Seed(time.Now().UnixNano())
	nicknames.InitNicknames()
	game.InitGameMemory()
	game.InitActiveConnection()

	http.HandleFunc("/ws", handlers.HandleWebSocket)
	http.HandleFunc("/api/nickname", handlers.HandleGetNickname)
	http.ListenAndServe(":8000", nil)
}
