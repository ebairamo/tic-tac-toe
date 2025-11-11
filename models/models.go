package models

type Gamer struct {
	ID     int
	Name   string
	Symbol string
}

type Game struct {
	ID          int
	Grid        [3][3]string
	Player1     Gamer
	Player2     Gamer
	CurrentTurn string
}

type GameMemory struct {
	ActiveGames     map[int]Game
	SearchingGamers map[int]Gamer
}

type Message struct {
	Action   string `json:"action"`
	PlayerId int    `json:"playerId"`
	Nickname string `json:"nickname"`
	GameId   int    `json:"gameId,omitempty"`
	Move     *Move  `json:"move,omitempty"`
}

type Move struct {
	Row int `json:"row"`
	Col int `json:"col"`
}

type NicknameResponse struct {
	Nickname    string `json:"nickname"`
	PlayerId    int    `json:"playerId"`
	UsersOnline int    `json:"usersOnline"`
	GamesOnline int    `json:"gamesOnline"`
}

type MessageGameFound struct {
	GameId     int    `json:"gameId"`
	YourSymbol string `json:"yourSymbol"`
	Enemy      Gamer  `json:"enemy"`
}
