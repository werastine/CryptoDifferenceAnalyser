// Package main
package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/werastine/CryptoDifferenceAnalyser/internal/api"
	"github.com/werastine/CryptoDifferenceAnalyser/service"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	client := &http.Client{Timeout: 5 * time.Second}
	waitGroup := &sync.WaitGroup{}

	srv := service.NewProviders(client, waitGroup)
	hdl := api.NewHandler(srv)

	router := hdl.RegisterRoutes()

	server := &http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Println("[ERROR] connecting to a port:", err)
		}
	}()

	<-ctx.Done()
	log.Println("[INFO] Recieved stop signal, making graceful shutdown...")
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Println("[ERROR] calling for graceful shutdown", err)
	}
	log.Println("Server is gracefuly stopped")
}
