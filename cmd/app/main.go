// Package main
package main

import (
	"log"
	"net/http"
	"sync"

	"time"

	"github.com/werastine/CryptoDifferenceAnalyser/internal/api"
	"github.com/werastine/CryptoDifferenceAnalyser/service"
)

func main() {
	client := &http.Client{Timeout: 5 * time.Second}
	waitGroup := &sync.WaitGroup{}

	srv := service.NewProviders(client, waitGroup)
	hdl := api.NewHandler(srv)
	hdl.RegisterRoutes()

	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Println("[ERROR] connecting to a port:", err)
	}

}
