package main

import (
	"album/handlers"
	"fmt"
	"net/http"

	"github.com/charmbracelet/log"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	// Config
	port := 3000
	hostAddr := fmt.Sprintf(":%d", port)

	// Router Setup
	mux := http.NewServeMux()

	htmlHandler := handlers.HTML{}
	mux.HandleFunc("GET /", htmlHandler.GetHomePage)
	mux.HandleFunc("GET /contact", htmlHandler.GetContactPage)
	
	// Server Setup
	serv := http.Server{
		Addr:    hostAddr,
		Handler: middleware.Logger(mux),
	}

	log.Info("starting server", "port", port)
	err := serv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
