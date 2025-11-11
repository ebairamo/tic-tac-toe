package game

import "fmt"

func ValidateMove(gameId int, playerId int, row int, col int) error {

	game := GameMemory.ActiveGames[gameId]
	if game.CurrentTurn == "" {
		game.CurrentTurn = "X"
	}
	fmt.Println(gameId, playerId, row, col, game.Grid)
	fmt.Println(game.Grid[row][col], game.CurrentTurn, game)
	fmt.Println(game.Player1.ID, "==", playerId, game.CurrentTurn, "==", game.Player1.Symbol)

	if game.Player1.ID == playerId && game.CurrentTurn == game.Player1.Symbol {
		fmt.Println(game.Grid[row][col])
		if game.Grid[row][col] == "" {
			game.Grid[row][col] = game.CurrentTurn
			if game.CurrentTurn == "X" {
				fmt.Println(game.CurrentTurn)
				game.CurrentTurn = "O"

			} else {
				fmt.Println(game.CurrentTurn)
				game.CurrentTurn = "X"
			}
		} else {
			return fmt.Errorf("ошибка клетка не пуста")
		}

	}

	if game.Player2.ID == playerId && game.CurrentTurn == game.Player2.Symbol {
		fmt.Println(game.Grid[row][col])
		if game.Grid[row][col] == "" {
			game.Grid[row][col] = game.CurrentTurn
			if game.CurrentTurn == "X" {
				fmt.Println(game.CurrentTurn)
				game.CurrentTurn = "O"

			} else {
				fmt.Println(game.CurrentTurn)
				game.CurrentTurn = "X"
			}
		} else {
			return fmt.Errorf("ошибка клетка не пуста")
		}

	}

	GameMemory.ActiveGames[gameId] = game
	return nil
}
