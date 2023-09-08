package api

import (
	"database/sql"
	"github.com/gorilla/mux"
)

// SetupRoutes sets up the API routes.
func SetupRoutes(router *mux.Router, db *sql.DB) {
	router.HandleFunc("/event-counts", GetEventCounts(db)).Methods("GET")
	router.HandleFunc("/unique-actors", GetUniqueActors(db)).Methods("GET")
	router.HandleFunc("/unique-repo-urls", GetUniqueRepoURLs(db)).Methods("GET")
	router.HandleFunc("/unique-emails", GetUniqueEmails(db)).Methods("GET")
}
