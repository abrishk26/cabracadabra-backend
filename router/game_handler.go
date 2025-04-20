package router

import (
	"net/http"
	"encoding/json"

	"github.com/abrishk26/cabracadabra-backend/utils"
)

type GameRoom struct {
	RoomID string
	TextToType string
	Players []Player
	GameState string
	GameDuration int
}

func CreateGame(w http.ResponseWriter, r *http.Request) {
	
	var req struct {
		GameDuration int `json:"game_duration"`	
	}

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		utils.BadRequestResponse(w, r, err)	
		return
	}

	roomID := utils.CreateID()

	newRoom := GameRoom{
		RoomID: roomID,
		TextToType: "Sample text for typing",
		Players: []Player{},
		GameState: "waiting",
		GameDuration: req.GameDuration,
	}

	GameRooms.lock.Lock()
	GameRooms.rooms = append(GameRooms.rooms, newRoom)
	GameRooms.lock.Unlock()

	err = utils.WriteJSON(w, http.StatusAccepted, map[string]any{"room": newRoom}, nil)

	if err != nil {
		utils.ServerErrorResponse(w, r)
		return
	}
}