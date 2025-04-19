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
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
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

	response, err := json.Marshal(newRoom)

	if err != nil {
		http.Error(w, "Unable to Convert Response into JSON", http.StatusInternalServerError)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}