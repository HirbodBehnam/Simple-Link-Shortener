package main

import (
	"RabinLink/api"
	"RabinLink/config"
	"RabinLink/database"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// Load config
	if len(os.Args) > 1 {
		config.LoadConfig(os.Args[1])
	} else {
		config.LoadConfig("config.json")
	}
	// Connect to database
	database.MainDatabase = database.New(config.Config.Database)
	defer database.MainDatabase.Close()
	// Start the webserver
	r := &http.ServeMux{}
	r.HandleFunc("/", api.Endpoint)
	srv := &http.Server{ // create a server
		Handler:     r,
		Addr:        config.Config.Listen,
		ReadTimeout: 10 * time.Second,
	}
	done := make(chan os.Signal, 1) // channel to get the system signals for graceful shutdown
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	// Gracefully shutdown the server
	go func() {
		var err error
		err = srv.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatalln("Cannot server the server:", err)
		}
	}()
	<-done // wait for ctrl + c
	log.Println("Received the stop signal")
}
