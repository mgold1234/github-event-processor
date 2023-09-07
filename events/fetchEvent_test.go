package events

import (
	"awsomeProject/pkg/models"
	"database/sql"
	"os"
	"testing"
)

var testDB *sql.DB

func TestMain(m *testing.M) {
	// Setup test database connection here
	var err error
	testDB, err = sql.Open("postgres", "user=username dbname=testdb sslmode=disable")
	if err != nil {
		panic("Error connecting to the test database: " + err.Error())
	}
	defer testDB.Close()

	// Create necessary tables or perform migrations for the test database

	// Run tests
	exitCode := m.Run()

	// Teardown code here, if necessary

	// Exit with the appropriate exit code
	os.Exit(exitCode)
}

func TestStoreGitHubEvent(t *testing.T) {
	// Create a test GitHub event
	event := models.GitHubEvent{
		Type: "PushEvent",
		Actor: models.Actor{
			Login: "JohnDoe"},
		Repo: models.Repo{
			URL: "https://github.com/example/repo",
		},
	}

	// Call the storeGitHubEvent function
	err := storeGitHubEvent(event)
	if err != nil {
		t.Fatalf("Error storing GitHub event: %v", err)
	}

	// Add assertions to verify that the event was stored correctly in the test database
}

func TestGetGitHubEvents(t *testing.T) {
	// Call the getGitHubEvents function to retrieve GitHub events from the test database
	_, err := getGitHubEvents()
	if err != nil {
		t.Fatalf("Error retrieving GitHub events: %v", err)
	}

}

func TestAddUniqueActor(t *testing.T) {
	actors := []string{"actor1", "actor2", "actor3"}
	newActor := "actor4"

	result := addUniqueActor(newActor, actors)

	// Check that the new actor is added
	if result[len(result)-1] != newActor {
		t.Errorf("Expected actor to be added, got %v", result)
	}

	// Check that the length of the slice is limited to 50
	if len(result) > 50 {
		t.Errorf("Expected length to be limited to 50, got %v", len(result))
	}
}

func TestAddUniqueRepoURL(t *testing.T) {
	repoURLs := []string{"url1", "url2", "url3"}
	newRepoURL := "url4"

	result := addUniqueRepoURL(newRepoURL, repoURLs)

	// Check that the new repo URL is added
	if result[len(result)-1] != newRepoURL {
		t.Errorf("Expected repo URL to be added, got %v", result)
	}

	// Check that the length of the slice is limited to 20
	if len(result) > 20 {
		t.Errorf("Expected length to be limited to 20, got %v", len(result))
	}
}

func TestAddUniqueEmail(t *testing.T) {
	emails := []string{"email1@example.com", "email2@example.com", "email3@example.com"}
	newEmail := "email4@example.com"

	result := addUniqueEmail(newEmail, emails)

	// Check that the new email is added
	if result[len(result)-1] != newEmail {
		t.Errorf("Expected email to be added, got %v", result)
	}
}

func TestCountEventTypes(t *testing.T) {
	events := []models.GitHubEvent{
		{Type: "PushEvent"},
		{Type: "PullRequestEvent"},
		{Type: "PushEvent"},
		{Type: "IssuesEvent"},
	}

	result := countEventTypes(events)

	// Check that event types are counted correctly
	if result["PushEvent"] != 2 {
		t.Errorf("Expected 2 PushEvents, got %v", result["PushEvent"])
	}
	if result["PullRequestEvent"] != 1 {
		t.Errorf("Expected 1 PullRequestEvent, got %v", result["PullRequestEvent"])
	}
	if result["IssuesEvent"] != 1 {
		t.Errorf("Expected 1 IssuesEvent, got %v", result["IssuesEvent"])
	}
}
