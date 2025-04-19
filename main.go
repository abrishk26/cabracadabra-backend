package main

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type HealthCheckMessage struct {
	Message string
}

func HealthCheck(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	message := HealthCheckMessage {
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



func main() {
	router := httprouter.New()
	router.GET("/health_check", HealthCheck)

	server := &http.Server{
		Addr: ":8080",
		Handler: router,
	}

	server.ListenAndServe()
}