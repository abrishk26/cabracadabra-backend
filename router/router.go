package router

import (
	"net/http"
	"sync"

	"github.com/julienschmidt/httprouter"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var GameRooms = struct {
	lock sync.Mutex
	rooms []GameRoom
} {
	lock: sync.Mutex{},
	rooms: []GameRoom{},
}



func SetUpRouter() *httprouter.Router {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/health_check", HealthCheck)
	router.HandlerFunc(http.MethodGet, "/ws", WebsocketTest)
	router.HandlerFunc(http.MethodPost, "/create_game", CreateGame)
	router.HandlerFunc(http.MethodGet, "/join_game", JoinGame)

	return router
}