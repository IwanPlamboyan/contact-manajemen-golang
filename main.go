package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func NewServer(router *httprouter.Router) *http.Server {
	return &http.Server{Addr: ":8080", Handler: router}
}

func main() {
	server, err := InitializedServer()
	if err != nil {
		log.Fatal("failed to initialize app:", err)
	}

	// start server
	log.Println("Starting server on port 8080")
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal("server error:", err)
	}
}