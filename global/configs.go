package global

import "github.com/opengovern/og-util/pkg/integration"

const (
	IntegrationTypeLower = "jira"                                    // example: aws, azure
	IntegrationName      = integration.Type("jira_cloud")            // example: aws_account, github_account
	OGPluginRepoURL      = "github.com/opengovern/og-describer-jira" // example: github.com/opengovern/og-describer-aws
)

type IntegrationCredentials struct {
	Username string `json:"username"`
	APIKey   string `json:"api_key"`
	BaseURL  string `json:"base_url"`
	Password string `json:"password"`
}
