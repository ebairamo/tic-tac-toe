package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"tic/game"
	"tic/models"
	"tic/nicknames"

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
			return err
		}

		draw := Draw(msg.GameId)
		winner := CheckWin(msg.GameId)
		if winner != "" {
			fmt.Println(winner + "winer")

		}
		SendBoardUpdate(msg.GameId, winner, draw)
		fmt.Println("🎮 Ход от игрока", msg.PlayerId, "на позицию", msg.Move.Row, msg.Move.Col)
	}
	fmt.Println("✅ ProcessMessage завершена успешно")
	return nil
}

func SendBoardUpdate(gameId int, winner string, gameStatus string) error {
	thisGame := game.GameMemory.ActiveGames[gameId]
	var finalWinner string
	var finalStatus string
	if winner != "" {
		finalStatus = "finished"
		finalWinner = winner
	} else if gameStatus == "draw" {
		finalStatus = gameStatus

	}
	BoardUpdate := models.BoardUpdate{
		GameId:      thisGame.ID,
		Grid:        thisGame.Grid,
		CurrentTurn: thisGame.CurrentTurn,
		GameStatus:  finalStatus,
		Winner:      finalWinner,
	}

	player1Conn, err := game.GetActiveConnection(thisGame.Player1.ID)
	if err != nil {

		return fmt.Errorf("ошибка получения соединения первого игрока")
	}
	player2Conn, err := game.GetActiveConnection(thisGame.Player2.ID)
	if err != nil {

		return fmt.Errorf("ошибка получения соединения второго игрока")
	}
	err = player1Conn.WriteJSON(BoardUpdate)
	if err != nil {

		return fmt.Errorf("ошибка отправления Json первому игроку")
	}
	err = player2Conn.WriteJSON(BoardUpdate)
	if err != nil {

		return fmt.Errorf("ошибка отправления Json второму игроку")
	}
	if BoardUpdate.GameStatus == "finished" || BoardUpdate.GameStatus == "draw" {

		player1Conn.Close()
		player2Conn.Close()
		delete(game.GameMemory.ActiveGames, gameId)
		nicknames.ReleaseNickname(thisGame.Player1.Name, thisGame.Player2.Name)
	}
	return nil
}

func CheckWin(gameId int) string {
	g := game.GameMemory.ActiveGames[gameId]
	patterns := [8][3][2]int{
		{{0, 0}, {0, 1}, {0, 2}}, // горизонталь 1
		{{1, 0}, {1, 1}, {1, 2}}, // горизонталь 2
		{{2, 0}, {2, 1}, {2, 2}}, // горизонталь 3
		{{0, 0}, {1, 0}, {2, 0}}, // вертикаль 1
		{{0, 1}, {1, 1}, {2, 1}}, // вертикаль 2
		{{0, 2}, {1, 2}, {2, 2}}, // вертикаль 3
		{{0, 0}, {1, 1}, {2, 2}}, // диагональ 1
		{{0, 2}, {1, 1}, {2, 0}}, // диагональ 2
	}

	for _, pattern := range patterns {
		a := g.Grid[pattern[0][0]][pattern[0][1]]
		b := g.Grid[pattern[1][0]][pattern[1][1]]
		c := g.Grid[pattern[2][0]][pattern[2][1]]

		if a != "" && a == b && b == c {
			return a // возвращаем "X" или "O"
		}
	}
	return "" // нет выигрыша
}

func Draw(gameId int) string {
	g := game.GameMemory.ActiveGames[gameId]
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if g.Grid[i][j] == "" {
				return "playing"
			}
		}
	}
	return "draw"
}
