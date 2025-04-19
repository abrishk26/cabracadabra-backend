package router

import (
	"net/http"

	"github.com/abrishk26/cabracadabra-backend/utils"
	"github.com/gorilla/websocket"
)

type Player struct {
	PlayerID    string
	Name        string
	IsConnected bool
	Result      TypingResult
	Conn        *websocket.Conn
}

type TypingResult struct {
	WPM      int
	Accuracy int
}

func JoinGame(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		conn.WriteJSON(map[string]string{
			"error": "Invalid JSON payload",
		})
		return

	}

	defer conn.Close()

	var joinReq struct {
		RoomID string `json:"room_id"`
		Name   string `json:"name"`
	}

	err = conn.ReadJSON(&joinReq)

	if err != nil {
		conn.WriteJSON(map[string]string{
			"error": "Invalid JSON payload",
		})
		return
	}

	newPlayer := Player{
		PlayerID:    utils.CreateID(),
		Name:        joinReq.Name,
		IsConnected: true,
		Result:      TypingResult{},
		Conn:        conn,
	}

	textToType := ""

	GameRooms.lock.Lock()
	isFound := false
	for _, room := range GameRooms.rooms {
		if room.RoomID == joinReq.RoomID {
			room.Players = append(room.Players, newPlayer)
			isFound = true
			textToType = room.TextToType
			break
		}
	}
	GameRooms.lock.Unlock()

	if !isFound {
		conn.WriteJSON(map[string]string{
			"error": "The Specified Room Does not exist",
		})
		return
	}

	mes := struct {
		TextToType string
	}{
		TextToType: textToType,
	}

	conn.WriteJSON(mes)
}
