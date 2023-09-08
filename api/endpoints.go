package api

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"sync"
)

// Define data structures to store the collected data
var eventTypeCount = make(map[string]int)
var uniqueActors []string
var uniqueRepoURLs []string
var uniqueEmails []string

var (
	eventTypeCountMutex sync.Mutex
	uniqueActorsMutex   sync.Mutex
	uniqueRepoURLsMutex sync.Mutex
	uniqueEmailsMutex   sync.Mutex
)

// API endpoints for retrieving data

func GetEventCounts(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Query the database to get event counts
		rows, err := db.Query("SELECT event_type, COUNT(*) FROM github GROUP BY event_type")
		if err != nil {
			http.Error(w, "Error querying the database", http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		// Create a map to store the event counts
		eventTypeCount := make(map[string]int)

		// Iterate through the query results and populate the event counts map
		for rows.Next() {
			var eventType string
			var count int
			err := rows.Scan(&eventType, &count)
			if err != nil {
				http.Error(w, "Error scanning database rows", http.StatusInternalServerError)
				return
			}
			eventTypeCount[eventType] = count
		}

		// Convert eventTypeCount to JSON and write it to the response
		data, err := json.Marshal(eventTypeCount)
		if err != nil {
			http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(data)
	}
}

func GetUniqueActors(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Query the database to get unique actors
		rows, err := db.Query("SELECT DISTINCT actor FROM github")
		if err != nil {
			http.Error(w, "Error querying the database", http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		// Create a slice to store the unique actors
		var uniqueActors []string

		// Iterate through the query results and add actors to the slice
		for rows.Next() {
			var actor string
			err := rows.Scan(&actor)
			if err != nil {
				http.Error(w, "Error scanning database rows", http.StatusInternalServerError)
				return
			}
			uniqueActors = append(uniqueActors, actor)
		}

		// Convert uniqueActors to JSON and write it to the response
		data, err := json.Marshal(uniqueActors)
		if err != nil {
			http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(data)
	}
}

func GetUniqueRepoURLs(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Query the database to get unique repo URLs
		rows, err := db.Query("SELECT DISTINCT repo_url FROM github")
		if err != nil {
			http.Error(w, "Error querying the database", http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		// Create a slice to store the unique repo URLs
		var uniqueRepoURLs []string

		// Iterate through the query results and populate the unique repo URLs slice
		for rows.Next() {
			var repoURL string
			err := rows.Scan(&repoURL)
			if err != nil {
				http.Error(w, "Error scanning database rows", http.StatusInternalServerError)
				return
			}
			uniqueRepoURLs = append(uniqueRepoURLs, repoURL)
		}

		// Convert uniqueRepoURLs to JSON and write it to the response
		data, err := json.Marshal(uniqueRepoURLs)
		if err != nil {
			http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(data)
	}
}

func GetUniqueEmails(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Query the database to get unique emails
		rows, err := db.Query("SELECT DISTINCT actor FROM github")
		if err != nil {
			http.Error(w, "Error querying the database", http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		// Create a slice to store the unique emails
		var uniqueEmails []string

		// Iterate through the query results and populate the unique emails slice
		for rows.Next() {
			var actor string
			err := rows.Scan(&actor)
			if err != nil {
				http.Error(w, "Error scanning database rows", http.StatusInternalServerError)
				return
			}
			uniqueEmails = append(uniqueEmails, actor)
		}

		// Convert uniqueEmails to JSON and write it to the response
		data, err := json.Marshal(uniqueEmails)
		if err != nil {
			http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(data)
	}
}
