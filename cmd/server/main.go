package main

import (
	"context"
	"go-api-playground/internal/di"
	"go-api-playground/pkg/middlewares"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-chi/chi"
)

func main() {
	log.Println("Hello, Playground!")

	if err := di.RegisterMediatorHandlers(); err != nil {
		log.Fatalf("Failed to register Mediator handlers: %v", err)
	}

	router := chi.NewRouter()

	router.Use(middlewares.LoggerMiddleware)

	server := &http.Server{
		Addr:         ":8081",
		Handler:      router,
		ReadTimeout:  3 * time.Second,
		IdleTimeout:  3 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	go launchServer(server)

	<-stop
	shutdownServer(server)
}

func launchServer(server *http.Server) {
	log.Printf("Starting now the server on localhost%v...🚀", server.Addr)

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Failed to start the server on localhost%v:%v", server.Addr, err)
	}
}

func shutdownServer(server *http.Server) {
	log.Printf("Shutting down the server on localhost%v...", server.Addr)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Failed to shutdown server: %v", err)
	}

	log.Println("Server gracefully stopped.")
}
