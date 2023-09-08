package main

import (
	"awsomeProject/api"
	"awsomeProject/events"
	"database/sql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Initialize your database and other dependencies here

	// Call FetchAndProcessEvents once when the program starts
	events.FetchAndProcessEvents()

	// Create a new Gorilla Mux router instance
	router := mux.NewRouter()

	// Create an HTTP server and bind it to your router
	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	// Read PostgreSQL connection details from environment variables
	dbUser := os.Getenv("PGSQL_USER")
	dbPassword := os.Getenv("PGSQL_PASSWORD")
	dbHost := os.Getenv("PGSQL_HOSTNAME")
	dbPort := os.Getenv("PGSQL_PORT")
	dbName := os.Getenv("PGSQL_DATABASE")

	// Create the PostgreSQL connection string
	connStr := "user=" + dbUser + " password=" + dbPassword + " host=" + dbHost + " port=" + dbPort + " dbname=" + dbName + " sslmode=disable"

	// Connect to the PostgreSQL database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	// Configure your API routes
	api.SetupRoutes(router, db)

	// Create a channel to signal the server to shut down
	shutdownChan := make(chan struct{})

	// Start the server in a separate goroutine
	go func() {
		if err := server.ListenAndServe(); err != nil {
			// Handle any errors that occur while starting the server
			if err != http.ErrServerClosed {
				panic(err)
			}
		}
	}()

	// Handle shutdown via HTTP endpoint
	http.HandleFunc("/shutdown", func(w http.ResponseWriter, r *http.Request) {
		// Gracefully shut down the server
		if err := server.Shutdown(nil); err != nil {
			// Handle any errors that occur during shutdown
			panic(err)
		}
		close(shutdownChan)
	})

	// Call initiateShutdown() here
	//	events.InitiateShutdown()

	// Listen for an interrupt signal (e.g., Ctrl+C)
	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Block until a signal or shutdown request is received
	select {
	case <-interruptChan:
		// Received interrupt signal
		log.Println("Received interrupt signal. Shutting down...")
	case <-shutdownChan:
		// Received shutdown request
		log.Println("Received shutdown request. Shutting down...")
	}

	os.Exit(0)
}
