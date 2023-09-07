package models

import "time"

// GitHubEvent represents a GitHub event
type GitHubEvent struct {
	Type      string    `json:"type"`
	Actor     Actor     `json:"actor"`
	Repo      Repo      `json:"repo"`
	Payload   Payload   `json:"payload"`
	CreatedAt time.Time `json:"created_at"`
}

// Actor represents the actor of a GitHub event
type Actor struct {
	Login string `json:"login"`
}

// Repo represents the repository of a GitHub event
type Repo struct {
	URL string `json:"url"`
}

// Payload represents the payload of a GitHub event
type Payload struct {
	Commits []Commit `json:"commits"`
}

// Commit represents a commit in the payload of a GitHub event
type Commit struct {
	Author Author `json:"author"`
}

// Author represents the author of a commit
type Author struct {
	Email string `json:"email"`
}
