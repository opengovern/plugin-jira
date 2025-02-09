package main

import (
	"fmt"
	"github.com/andygrunwald/go-jira"
)

type JiraServerInfo struct {
	BaseURL     string `json:"baseUrl"`
	ServerTitle string `json:"serverTitle"`
}

// JiraIntegrationDiscovery fetches Jira server info details using the provided url, username and token.
func JiraIntegrationDiscovery(baseUrl, username, pass string, isLocal bool) (*JiraServerInfo, error) {
	var instance JiraServerInfo
	var finalURL string
	if isLocal {
		finalURL = "rest/api/2/serverInfo"
	} else {
		finalURL = "rest/api/3/serverInfo"
	}

	tp := jira.BasicAuthTransport{
		Username: username,
		Password: pass,
	}

	client, err := jira.NewClient(tp.Client(), baseUrl)
	if err != nil {
		return nil, err
	}

	req, err := client.NewRequest("GET", finalURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	_, err = client.Do(req, &instance)
	if err != nil {
		return nil, fmt.Errorf("request execution failed: %w", err)
	}

	return &instance, nil
}
