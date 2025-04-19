package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/julienschmidt/httprouter"
	"github.com/google/uuid"
)

type GameRoom struct {
	RoomID string
	TextToType string
	Players []Player
	GameState string
	GameDuration int
}

type Player struct {
	PlayerID string
	Name string
	IsConnected bool
	Result TypingResult
}

type TypingResult struct {
	WPM int
	Accuracy int
}

type HealthCheckMessage struct {
	Message string
}

type CreateGameRequest struct {
	GameDuration int `json:"game_duration"`
}

var GameRooms = struct {
	lock sync.Mutex
	rooms []GameRoom
} {
	lock: sync.Mutex{},
	rooms: []GameRoom{},
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	message := HealthCheckMessage{
		Message: "The API is Working",
	}

	response, err := json.Marshal(message)

	if err != nil {
		http.Error(w, "INTERNAL SERVER ERROR", http.StatusInternalServerError)
		return
	}

	response = append(response, '\n')
	w.Header().Add("Content-Type", "application/json")
	w.Write(response)
}

func CreateGame(w http.ResponseWriter, r *http.Request) {
	
	var req CreateGameRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	roomID := createID()

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

func createID() string {
	return uuid.New().String()
}

func WebsocketTest(w http.ResponseWriter, r *http.Request) {
	// Upgrage the HTTP connection to a WebSocket connection
	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Fatal(err)
		return
	}

	defer conn.Close()

	messageType, p, err := conn.ReadMessage()
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println("Received", string(p))

	if err = conn.WriteMessage(messageType, p); err != nil {
		log.Fatal(err)
	}
}

func main() {
	router := httprouter.New()
	router.HandlerFunc(http.MethodGet, "/health_check", HealthCheck)
	router.HandlerFunc(http.MethodGet, "/ws", WebsocketTest)
	router.HandlerFunc(http.MethodPost, "/create_game", CreateGame)

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	server.ListenAndServe()
}
