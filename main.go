package main

import (
	"awsomeProject/api"
	"awsomeProject/events"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Initialize the events fetching process
	events.FetchGitHubEvents()

	// Create a new Gorilla Mux router instance
	router := mux.NewRouter()

	// Call the SetupRoutes function to configure your API routes
	api.SetupRoutes(router)

	// Create an HTTP server and bind it to your router
	server := &http.Server{
		Addr:    ":8080", // Choose the port you want to run your API on
		Handler: router,
	}

	// Start the HTTP server in a goroutine
	go func() {
		if err := server.ListenAndServe(); err != nil {
			// Handle any errors that occur while starting the server
			panic(err)
		}
	}()

	// Listen for an interrupt signal (e.g., Ctrl+C)
	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Block until an interrupt signal is received
	<-interruptChan

	// Gracefully shut down the server
	if err := server.Shutdown(nil); err != nil {
		// Handle any errors that occur during shutdown
		panic(err)
	}
}
