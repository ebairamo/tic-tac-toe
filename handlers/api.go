package handlers

import (
	"encoding/json"
	"net/http"
	"tic/game"
	"tic/models"
	"tic/nicknames"
)

func HandleGetNickname(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	gamerID := game.GenerateGamerID()
	nick, err := nicknames.GetRandomNickname(nicknames.AvailableNicknames)
	if err != nil {
		http.Error(w, "Ошибка получения никнейма", http.StatusInternalServerError)
		return
	}
	onlineGamers := (len(game.GameMemory.ActiveGames) * 2) + len(game.GameMemory.SearchingGamers)
	activeGames := len(game.GameMemory.ActiveGames)
	var ResponseNickname models.NicknameResponse = models.NicknameResponse{
		Nickname:    nick,
		PlayerId:    gamerID,
		UsersOnline: onlineGamers,
		GamesOnline: activeGames,
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(ResponseNickname)
}

func HandleGetStats(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Метод запрещен", http.StatusMethodNotAllowed)
		return
	}

	onlineGamers := (len(game.GameMemory.ActiveGames) * 2) + len(game.GameMemory.SearchingGamers)
	activeGames := len(game.GameMemory.ActiveGames)
	messageStats := models.MessageStats{
		PlayerOnline: onlineGamers,
		GamesOnline:  activeGames,
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(messageStats)
}
