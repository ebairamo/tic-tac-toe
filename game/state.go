package game

import (
	"fmt"
	"tic/models"

	"github.com/gorilla/websocket"
)

var GameMemory models.GameMemory
var gamerIDCounter int
var gameIDCounter int
var activeConnections map[int]*websocket.Conn

func GenerateGamerID() int {
	gamerIDCounter++
	return gamerIDCounter
}
func GenerateGameID() int {
	gameIDCounter++
	return gameIDCounter
}

func InitGameMemory() {
	GameMemory.SearchingGamers = make(map[int]models.Gamer)
	GameMemory.ActiveGames = make(map[int]models.Game)
}
func InitActiveConnection() {
	activeConnections = make(map[int]*websocket.Conn)
}

func AddActiveConnection(id int, conn *websocket.Conn) error {
	if conn == nil {
		return fmt.Errorf("соединение пусто")
	}
	fmt.Println("✅ Добавлено соединение для игрока", id) // ← ДОБ АВЬ
	activeConnections[id] = conn
	return nil
}

func GetActiveConnection(id int) (*websocket.Conn, error) {
	activeConnection := activeConnections[id]
	if activeConnection != nil {
		return activeConnection, nil
	} else {
		return nil, fmt.Errorf("ошибка нет соединения")
	}
}

func RemoveActiveConnection(id int) error {
	activeConnection := activeConnections[id]
	if activeConnection != nil {
		delete(activeConnections, id)
		return nil
	} else {
		return fmt.Errorf("ошибка удаления такого айди нет")

	}
}
