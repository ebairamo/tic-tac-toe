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
		if _, exists := game.GameMemory.SearchingGamers[msg.PlayerId]; exists {
			return fmt.Errorf("ты уже ищешь соперника")
		}

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

				} else if count == 1 {
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
			player1conn, err := game.GetActiveConnection(id1)
			if err != nil {
				fmt.Println("ошибка получения соединеня первого игрока")
				return err
			}
			messagePlayer1 := models.MessageGameFound{
				GameId:     gameId,
				YourSymbol: "X",
				Enemy: models.Gamer{
					ID:     id2,
					Name:   player2.Name,
					Symbol: "O",
				},
			}
			err = player1conn.WriteJSON(messagePlayer1)
			if err != nil {
				fmt.Println("ошибка отправки сообщения первому игроку")
			}
			player2conn, err := game.GetActiveConnection(id2)
			if err != nil {
				fmt.Println("ошибка получения соединеня второго игрока")
				return err
			}
			messagePlayer2 := models.MessageGameFound{
				GameId:     gameId,
				YourSymbol: "O",
				Enemy: models.Gamer{
					ID:     id1,
					Name:   player1.Name,
					Symbol: "X",
				},
			}
			err = player2conn.WriteJSON(messagePlayer2)
			fmt.Println("✉️ Отправлено игроку 1 (ID:", id1, player1, ")")
			fmt.Println("✉️ Отправлено игроку 2 (ID:", id2, player2, ")")
			if err != nil {
				fmt.Println("ошибка отправки сообщения второму игроку")
			}
		}

	} else if msg.Action == "move" {
		err := game.ValidateMove(msg.GameId, msg.PlayerId, msg.Move.Row, msg.Move.Col)
		if err != nil {
		}
		fmt.Println("🎮 Ход от игрока", msg.PlayerId, "на позицию", msg.Move.Row, msg.Move.Col)
	}
	fmt.Println("✅ ProcessMessage завершена успешно")
	return nil
}
