package maps

import (
	"github.com/opengovern/og-describer-jira/discovery/pkg/es"
)

var ResourceTypesToTables = map[string]string{
  "Jira/Project": "jira_project",
  "Jira/Issue": "jira_issue",
  "Jira/Board": "jira_board",
}

var ResourceTypeToDescription = map[string]interface{}{
  "Jira/Project": opengovernance.Project{},
  "Jira/Issue": opengovernance.Issue{},
  "Jira/Board": opengovernance.Board{},
}

var TablesToResourceTypes = map[string]string{
  "jira_project": "Jira/Project",
  "jira_issue": "Jira/Issue",
  "jira_board": "Jira/Board",
}
