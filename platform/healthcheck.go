package main

import (
	"fmt"
	"github.com/andygrunwald/go-jira"
)

// JiraIntegrationHealthcheck checks if the given Jira credentials is valid.
func JiraIntegrationHealthcheck(baseUrl, username, pass string, isLocal bool) (bool, error) {
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
		return false, err
	}

	req, err := client.NewRequest("GET", finalURL, nil)
	if err != nil {
		return false, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	_, err = client.Do(req, &instance)
	if err != nil {
		return false, fmt.Errorf("request execution failed: %w", err)
	}

	return true, nil
}
