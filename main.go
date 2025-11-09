package main

import (
	"math/rand"
	"net/http"
	"tic/handlers"
	"tic/models"
	"tic/nicknames"
	"time"
)

var gameMemory models.GameMemory
var Nickname map[string]bool

func main() {
	rand.Seed(time.Now().UnixNano())
	Nickname = nicknames.GenerateNicknames()
	gameMemory.SearchingGamers = make(map[int]models.Gamer)
	gameMemory.ActiveGames = make(map[int]models.Game)

	http.HandleFunc("/ws", handlers.HandleWebSocket)
	http.ListenAndServe(":8000", nil)
}
