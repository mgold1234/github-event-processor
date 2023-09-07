package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestGetEventCounts(t *testing.T) {
	// Set up test data
	eventTypeCount = map[string]int{
		"EventType1": 10,
		"EventType2": 20,
		"EventType3": 5,
	}
	// Set up a request to the endpoint
	req, err := http.NewRequest("GET", "/event_counts", nil)
	if err != nil {
		t.Fatalf("error creating request: %v", err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Call the handler function with the request and response recorder
	GetEventCounts(rr, req)

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
	expected := map[string]int{"EventType1": 10, "EventType2": 20, "EventType3": 5}
	if !reflect.DeepEqual(actualMap, expected) {
		t.Errorf("handler returned unexpected body: got %v want %v", actualMap, expected)
	}
}

func TestGetUniqueActors(t *testing.T) {
	// Set up test data
	uniqueActors = []string{
		"Actor1",
		"Actor2",
		"Actor3",
	}
	// Set up a request to the endpoint
	req, err := http.NewRequest("GET", "/unique_actors", nil)
	if err != nil {
		t.Fatalf("error creating request: %v", err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Call the handler function with the request and response recorder
	GetUniqueActors(rr, req)

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

func TestGetUniqueRepoURLs(t *testing.T) {
	var testUniqueRepoURLs = []string{"url1", "url2", "url3"}

	// Set the test data for unique repo URLs
	uniqueRepoURLs = testUniqueRepoURLs

	req, err := http.NewRequest("GET", "/unique-repo-urls", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	http.HandlerFunc(GetUniqueRepoURLs).ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Convert the test data to JSON
	expected, err := json.Marshal(testUniqueRepoURLs)
	if err != nil {
		t.Fatal(err)
	}

	// Check the response body
	if rr.Body.String() != string(expected) {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), string(expected))
	}
}

func TestGetUniqueEmails(t *testing.T) {
	var testUniqueEmails = []string{"email1@example.com", "email2@example.com", "email3@example.com"}
	// Set the test data for unique emails
	uniqueEmails = testUniqueEmails

	req, err := http.NewRequest("GET", "/unique-emails", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	http.HandlerFunc(GetUniqueEmails).ServeHTTP(rr, req)

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
