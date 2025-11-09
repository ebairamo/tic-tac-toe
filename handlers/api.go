package handlers

import (
	"net/http"
	"tic/nicknames"
)

func HandleGetNickname(w http.ResponseWriter, r *http.Request) {
	nick, err := nicknames.GetRandomNickname(Nickname)
}
