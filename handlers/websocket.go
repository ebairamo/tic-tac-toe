package handlers

import (
	"fmt"
	"net/http"
	"tic/models"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	defer conn.Close()

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
	
	if msg.Action == "quickgame"{
		var gamer models.Gamer

		gameMemory.SearchingGamers[msg.PlayerId] =
	}
	return nil
}
