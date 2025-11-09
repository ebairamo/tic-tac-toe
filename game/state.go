package game

import "tic/models"

var GameMemory models.GameMemory
var gamerIDCounter int

func GenerateGamerID() int {
	gamerIDCounter++
	return gamerIDCounter
}

func InitGameMemory() {
	GameMemory.SearchingGamers = make(map[int]models.Gamer)
	GameMemory.ActiveGames = make(map[int]models.Game)
}
