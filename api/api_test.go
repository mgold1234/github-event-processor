package api

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"
	"time"
)

func TestGetEventCounts(t *testing.T) {
	// Set up a temporary test database
	db := setupTestDatabase(t)
	defer db.Close()

	// Set up test data in the database
	setupTestData(db)

	// Set up a request to the endpoint
	req, err := http.NewRequest("GET", "/event-counts", nil)
	if err != nil {
		t.Fatalf("error creating request: %v", err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Call the handler function with the request and response recorder
	handler := GetEventCounts(db)
	handler(rr, req)

	// Check the response status code
	if rr.Code != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", rr.Code, http.StatusOK)
	}

	// Parse the actual response body
	var actualMap map[string]int
	if err := json.Unmarshal(rr.Body.Bytes(), &actualMap); err != nil {
		t.Fatalf("error unmarshalling response body: %v", err)
	}

	// Compare the actual map with the expected map
	expected := map[string]int{
		"EventType1": 10,
		"EventType2": 20,
		"EventType3": 5,
	}
	if !reflect.DeepEqual(actualMap, expected) {
		t.Errorf("handler returned unexpected body: got %v want %v", actualMap, expected)
	}
}

// Helper function to set up a temporary test database
func setupTestDatabase(t *testing.T) *sql.DB {
	// Open a database connection to a test database

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
		t.Fatalf("error connecting to database: %v", err)
	}

	// Ensure the database tables are created (you may need to modify this)
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS github (
		id SERIAL PRIMARY KEY,
		event_type VARCHAR(255),
		...
	);`)
	if err != nil {
		t.Fatalf("error creating table: %v", err)
	}

	return db
}

func setupTestData(db *sql.DB) {
	// Insert test data into the database as needed for your test
	_, err := db.Exec("INSERT INTO github (event_type, actor, repo_url, created_at) VALUES ($1, $2, $3, $4)", "EventType1", "Actor1", "RepoURL1", time.Now())
	if err != nil {
		log.Fatalf("error inserting test data: %v", err)
	}

	// Insert more test data as needed
	_, err = db.Exec("INSERT INTO github (event_type, actor, repo_url, created_at) VALUES ($1, $2, $3, $4)", "EventType2", "Actor2", "RepoURL2", time.Now())
	if err != nil {
		log.Fatalf("error inserting test data: %v", err)
	}

}

func TestGetUniqueActors(t *testing.T) {
	// Set up a temporary test database
	db := setupTestDatabase(t)
	defer db.Close()

	// Set up test data in the database
	setupTestData(db)

	// Set up a request to the endpoint
	req, err := http.NewRequest("GET", "/unique-actors", nil)
	if err != nil {
		t.Fatalf("error creating request: %v", err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Call the handler function with the request and response recorder
	handler := GetUniqueActors(db)
	handler(rr, req)

	// Check the response status code
	if rr.Code != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", rr.Code, http.StatusOK)
	}

	// Parse the actual response body
	var actual []string
	if err := json.Unmarshal(rr.Body.Bytes(), &actual); err != nil {
		t.Fatalf("error unmarshalling response body: %v", err)
	}

	// Compare the actual slice with the expected slice
	expected := []string{"Actor1", "Actor2", "Actor3"}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("handler returned unexpected body: got %v want %v", actual, expected)
	}
}

func TestGetUniqueEmails(t *testing.T) {
	// Set up a temporary test database
	db := setupTestDatabase(t)
	defer db.Close()

	// Set up test data in the database
	setupTestData(db)

	var testUniqueEmails = []string{"email1@example.com", "email2@example.com", "email3@example.com"}

	// Set the test data for unique emails in the database
	_, err := db.Exec("INSERT INTO github (event_type, actor, repo_url, created_at) VALUES ($1, $2, $3, $4)", "EventType1", "Actor1", "url1", time.Now())
	if err != nil {
		t.Fatalf("error inserting test data: %v", err)
	}

	_, err = db.Exec("INSERT INTO github (event_type, actor, repo_url, created_at) VALUES ($1, $2, $3, $4)", "EventType2", "Actor2", "url2", time.Now())
	if err != nil {
		t.Fatalf("error inserting test data: %v", err)
	}

	_, err = db.Exec("INSERT INTO github (event_type, actor, repo_url, created_at) VALUES ($1, $2, $3, $4)", "EventType3", "Actor3", "url3", time.Now())
	if err != nil {
		t.Fatalf("error inserting test data: %v", err)
	}

	req, err := http.NewRequest("GET", "/unique-emails", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	// Create a handler function that uses the database connection
	handler := GetUniqueEmails(db)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Convert the test data to JSON
	expected, err := json.Marshal(testUniqueEmails)
	if err != nil {
		t.Fatal(err)
	}

	// Check the response body
	if rr.Body.String() != string(expected) {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), string(expected))
	}
}
