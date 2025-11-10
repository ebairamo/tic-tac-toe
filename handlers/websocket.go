package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"tic/game"
	"tic/models"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("playerId")
	playerId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("ошибка преобразования player id")
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	defer conn.Close()

	err = game.AddActiveConnection(playerId, conn)
	if err != nil {
		fmt.Println("ошибка добавления активного соединения")
		return
	}
	defer game.RemoveActiveConnection(playerId)
	for {
		var msg models.Message
		err := conn.ReadJSON(&msg)
		if err != nil {
			fmt.Println(err)
			break
		}
		err = ProcessMessage(msg)
		if err != nil {
			fmt.Println(err)
			break
		}
	}
}

func ProcessMessage(msg models.Message) error {

	if msg.Action == "quickgame" {
		gamer := models.Gamer{
			ID:   msg.PlayerId,
			Name: msg.Nickname,
		}
		game.GameMemory.SearchingGamers[msg.PlayerId] = gamer
		if len(game.GameMemory.SearchingGamers) >= 2 {
			var player1, player2 models.Gamer
			var id1, id2 int
			var count int
			for id, gamer := range game.GameMemory.SearchingGamers {
				if count == 0 {
					player1 = gamer
					id1 = id
					player1.Symbol = "X"

					count++
				}
				if count == 1 {
					player2 = gamer
					id2 = id
					player2.Symbol = "O"
					break
				}

			}
			var matchedGame models.Game
			var grid ([3][3]string)
			gameId := game.GenerateGameID()
			matchedGame = models.Game{
				ID:      gameId,
				Grid:    grid,
				Player1: player1,
				Player2: player2,
			}
			delete(game.GameMemory.SearchingGamers, id1)
			delete(game.GameMemory.SearchingGamers, id2)
			game.GameMemory.ActiveGames[gameId] = matchedGame
		}

	}
	return nil
}
