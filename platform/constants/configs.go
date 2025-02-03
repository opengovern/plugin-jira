package constants

import "github.com/opengovern/og-util/pkg/integration"
import _ "embed"

//go:embed ui-spec.json
var UISpec []byte

//go:embed manifest.yaml
var Manifest []byte

//go:embed Setup.md
var SetupMd []byte

const (
	IntegrationName = integration.Type("jira") // example: aws_cloud, azure_subscription, github_account
)

const (
	DescriberDeploymentName = "og-describer-jira"
	DescriberRunCommand     = "/og-describer-jira"
)
