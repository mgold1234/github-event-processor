package events

import (
	"awsomeProject/pkg/client"
	"awsomeProject/pkg/models"
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

var events []models.GitHubEvent

// Declare data structures and slices
var eventTypeCount = make(map[string]int)
var uniqueActors []string
var uniqueRepoURLs []string
var uniqueEmails []string

var (
	eventTypeCountMutex sync.Mutex
	uniqueActorsMutex   sync.Mutex
	uniqueRepoURLsMutex sync.Mutex
	uniqueEmailsMutex   sync.Mutex
	fetchMutex          sync.Mutex // Create a mutex to ensure only one goroutine is running FetchAndProcessEvents
)
var isFetching = false // Flag to track if FetchAndProcessEvents is already running
var fetchOnce sync.Once

var db *sql.DB

func init() {
	var err error

	// Read PostgreSQL connection details from environment variables
	dbUser := os.Getenv("PGSQL_USER")
	dbPassword := os.Getenv("PGSQL_PASSWORD")
	dbHost := os.Getenv("PGSQL_HOSTNAME")
	dbPort := os.Getenv("PGSQL_PORT")
	dbName := os.Getenv("PGSQL_DATABASE")

	// Create the PostgreSQL connection string
	connStr := "user=" + dbUser + " password=" + dbPassword + " host=" + dbHost + " port=" + dbPort + " dbname=" + dbName + " sslmode=disable"

	// Connect to the PostgreSQL database
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	createGitHubEventsTable()
}

func createGitHubEventsTable() {
	createTableSQL := `
        CREATE TABLE IF NOT EXISTS github (
            id serial PRIMARY KEY,
            event_type varchar(255),
            actor varchar(255),
            repo_url varchar(255),
            created_at timestamp NOT NULL
        );
    `

	_, err := db.Exec(createTableSQL)
	if err != nil {
		log.Fatalf("Error creating github table: %v", err)
	}
}

// storeGitHubEvent to store GitHub event data in the database
func storeGitHubEvent(event models.GitHubEvent) error {
	_, err := db.Exec("INSERT INTO github (event_type, actor, repo_url, created_at) VALUES ($1, $2, $3, $4)", event.Type, event.Actor.Login, event.Repo.URL, event.CreatedAt)
	return err
}

// getGitHubEvents to retrieve GitHub event data from the database
func getGitHubEvents() ([]models.GitHubEvent, error) {
	rows, err := db.Query("SELECT event_type, actor, repo_url, created_at FROM github")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []models.GitHubEvent
	for rows.Next() {
		var event models.GitHubEvent
		err := rows.Scan(&event.Type, &event.Actor.Login, &event.Repo.URL, &event.CreatedAt)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	return events, nil
}

func CleanupOldData(db *sql.DB) {
	// Calculate the date threshold based on your retention policy
	retentionThreshold := time.Now().AddDate(0, 0, -1) // Retain data for 30 days

	// Execute SQL query to delete old data
	_, err := db.Exec("DELETE FROM github WHERE created_at < $1", retentionThreshold)
	if err != nil {
		log.Printf("Error cleaning up data: %v", err)
	}
}

// FetchGitHubEvents fetches GitHub events periodically.
func FetchGitHubEvents() {
	ticker := time.NewTicker(1 * time.Minute) // Fetch events every 1 minute

	for {
		select {
		case <-ticker.C:
			if !isFetching {
				isFetching = true
				go func() {
					defer func() {
						isFetching = false
					}()
					FetchAndProcessEvents()
				}()
			}
		}
	}
}

// FetchAndProcessEvents fetches and processes GitHub events.
func FetchAndProcessEvents() {
	// The code inside this function will run only once, regardless of how many times it's called.
	fetchOnce.Do(func() {
		// Your fetch and processing logic here
		fmt.Println("Fetching and processing events...")
		githubToken := os.Getenv("GITHUB_ACCESS_TOKEN")
		if githubToken == "" {
			fmt.Println("GitHub access token not found. Please set the GITHUB_ACCESS_TOKEN environment variable.")
			return
		}
		client := client.CreateGitHubClient()

		// Create an HTTP GET request
		req, err := http.NewRequest("GET", "https://api.github.com/events", nil)
		if err != nil {
			log.Fatalf("Error creating HTTP request: %v", err)
			return
		}

		// Set headers, including the Authorization header for authentication
		req.Header.Set("Authorization", "token "+githubToken)

		// Make the GET request
		resp, err := client.Do(req)
		if err != nil {
			log.Fatalf("Error making GET request: %v", err)
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			log.Printf("Non-200 response: %d", resp.StatusCode) // Log the response status code
			return
		}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Printf("Error reading response body: %v", err) // Log the error
			return
		}
		// Parse JSON response
		if err := json.Unmarshal(body, &events); err != nil {
			log.Printf("Error parsing JSON: %v", err) // Log the error
			return
		}
		for _, event := range events {
			// Increment event type count
			eventType := event.Type
			eventTypeCount[eventType]++

			// Add actor to unique actors slice
			actorLogin := event.Actor.Login
			uniqueActors = addUniqueActor(actorLogin, uniqueActors)

			// Add repository URL to unique repo URLs slice
			repoURL := event.Repo.URL
			uniqueRepoURLs = addUniqueRepoURL(repoURL, uniqueRepoURLs)

			// Extract and add unique email addresses from commits
			for _, commit := range event.Payload.Commits {
				authorEmail := commit.Author.Email
				uniqueEmails = addUniqueEmail(authorEmail, uniqueEmails)
			}
			err := storeGitHubEvent(event)
			if err != nil {
				log.Printf("Error storing GitHub event: %v", err)
				// Handle the error as needed (e.g., log it, return an error response, etc.)
			}
		}
		//	CleanupOldData(db)
	})
}

// countEventTypes counts event types.
func countEventTypes(events []models.GitHubEvent) map[string]int {
	eventCounts := make(map[string]int)
	for _, event := range events {
		eventType := event.Type
		eventCounts[eventType]++
	}
	return eventCounts
}

// addUniqueActor adds unique actors to the slice (maintaining the last 50).
func addUniqueActor(actor string, actors []string) []string {
	uniqueActorsMutex.Lock()         // Lock the mutex before modifying the shared data
	defer uniqueActorsMutex.Unlock() // Ensure the mutex is unlocked when the function exits

	// Check if the actor is already in the slice
	for _, a := range actors {
		if a == actor {
			return actors // Actor is already in the slice
		}
	}

	// Add the actor to the slice
	actors = append(actors, actor)

	// Limit the slice size to 50 (remove the oldest actor if necessary)
	if len(actors) > 50 {
		actors = actors[1:]
	}

	return actors
}

// addUniqueRepoURL adds unique repository URLs to the slice (maintaining the last 20).
func addUniqueRepoURL(repoURL string, repoURLs []string) []string {
	uniqueRepoURLsMutex.Lock()         // Lock the mutex before modifying the shared data
	defer uniqueRepoURLsMutex.Unlock() // Ensure the mutex is unlocked when the function exits
	// Check if the repoURL is already in the slice
	for _, url := range repoURLs {
		if url == repoURL {
			return repoURLs // Repo URL is already in the slice
		}
	}

	// Add the repoURL to the slice
	repoURLs = append(repoURLs, repoURL)

	// Limit the slice size to 20 (remove the oldest URL if necessary)
	if len(repoURLs) > 20 {
		repoURLs = repoURLs[1:]
	}

	return repoURLs
}

// addUniqueEmail adds unique email addresses to the slice.
func addUniqueEmail(email string, emails []string) []string {
	uniqueEmailsMutex.Lock()         // Lock the mutex before modifying the shared data
	defer uniqueEmailsMutex.Unlock() // Ensure the mutex is unlocked when the function exits
	// Check if the email is already in the slice
	for _, e := range emails {
		if e == email {
			return emails // Email is already in the slice
		}
	}

	// Add the email to the slice
	uniqueEmails = append(uniqueEmails, email)

	return uniqueEmails
}
