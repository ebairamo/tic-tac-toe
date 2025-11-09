package main

import (
	"math/rand"
	"net/http"
	"tic/handlers"
	"tic/models"
	"time"
)

var gameMemory models.GameMemory
var nickname map[string]bool

func main() {
	rand.Seed(time.Now().UnixNano())
	gameMemory.SearchingGamers = make(map[int]models.Gamer)
	gameMemory.ActiveGames = make(map[int]models.Game)

	http.HandleFunc("/ws", handlers.HandleWebSocket)
	http.ListenAndServe(":8000", nil)
}
