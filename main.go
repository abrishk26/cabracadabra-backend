package main

import (
	"net/http"

	"github.com/abrishk26/cabracadabra-backend/router"
)




func main() {
	router := router.SetUpRouter()

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	server.ListenAndServe()
}
