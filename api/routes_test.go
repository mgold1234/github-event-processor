// api/api_test.go

package api

import (
	"database/sql"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func TestSetupRoutes(t *testing.T) {
	// Create a new router
	router := mux.NewRouter()

	// Create a test database connection (replace with your database details)
	testDB, err := sql.Open("postgres", "user=mydb dbname=mydb sslmode=disable")
	if err != nil {
		t.Fatalf("Failed to open database connection: %v", err)
	}
	defer testDB.Close()

	// Set up routes with the test database
	SetupRoutes(router, testDB)

	// Define test cases for the routes
	testCases := []struct {
		path       string
		method     string
		statusCode int
	}{
		{"/event-counts", "GET", http.StatusOK},
		{"/unique-actors", "GET", http.StatusOK},
		{"/unique-repo-urls", "GET", http.StatusOK},
		{"/unique-emails", "GET", http.StatusOK},
		// Add more test cases as needed
	}

	// Iterate through test cases
	for _, tc := range testCases {
		t.Run(tc.path, func(t *testing.T) {
			// Create a request to the test server
			req, err := http.NewRequest(tc.method, tc.path, nil)
			if err != nil {
				t.Fatalf("Failed to create request: %v", err)
			}

			// Create a ResponseRecorder to record the response
			rr := httptest.NewRecorder()

			// Serve the request using the router
			router.ServeHTTP(rr, req)

			// Check the response status code
			if rr.Code != tc.statusCode {
				t.Errorf("Handler returned wrong status code: got %v, want %v", rr.Code, tc.statusCode)
			}
		})
	}
}
