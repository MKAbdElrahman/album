package main

import (
	"album/handlers"
	"fmt"
	"net/http"

	"github.com/charmbracelet/log"
)

func main() {
	// Config
	port := 3000
	hostAddr := fmt.Sprintf(":%d", port)

	// Router Setup
	mux := http.NewServeMux()

	htmlHandler := handlers.HTML{}
	mux.HandleFunc("GET /", htmlHandler.GetHomePage)

	// Server Setup
	serv := http.Server{
		Addr:    hostAddr,
		Handler: mux,
	}

	log.Info("starting server", "port", port)
	err := serv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
