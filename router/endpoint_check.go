package router

import (
	"net/http"
	"encoding/json"
	"log"
	"fmt"
)
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	message := struct {
		Message string
	}{
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