package api

import (
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

func GetEventCounts(w http.ResponseWriter, r *http.Request) {
	// Lock and read eventTypeCount
	eventTypeCountMutex.Lock()
	defer eventTypeCountMutex.Unlock()

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

func GetUniqueActors(w http.ResponseWriter, r *http.Request) {
	// Lock and read uniqueActors
	uniqueActorsMutex.Lock()
	defer uniqueActorsMutex.Unlock()

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

func GetUniqueRepoURLs(w http.ResponseWriter, r *http.Request) {
	// Lock and read uniqueRepoURLs
	uniqueRepoURLsMutex.Lock()
	defer uniqueRepoURLsMutex.Unlock()

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

func GetUniqueEmails(w http.ResponseWriter, r *http.Request) {
	// Lock and read uniqueEmails
	uniqueEmailsMutex.Lock()
	defer uniqueEmailsMutex.Unlock()

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
