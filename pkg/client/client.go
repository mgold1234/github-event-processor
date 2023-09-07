package client

import "net/http"

func CreateGitHubClient() *http.Client {
	client := &http.Client{}
	return client
}
