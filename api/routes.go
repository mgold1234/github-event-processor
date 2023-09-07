package api

import (
	"github.com/gorilla/mux"
)

// SetupRoutes sets up the API routes.
func SetupRoutes(router *mux.Router) {
	router.HandleFunc("/event-counts", GetEventCounts).Methods("GET")
	router.HandleFunc("/unique-actors", GetUniqueActors).Methods("GET")
	router.HandleFunc("/unique-repo-urls", GetUniqueRepoURLs).Methods("GET")
	router.HandleFunc("/unique-emails", GetUniqueEmails).Methods("GET")
}
