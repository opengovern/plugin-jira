package maps

import (
	"github.com/opengovern/og-describer-jira/discovery/pkg/es"
)

var ResourceTypesToTables = map[string]string{
  "Jira/Project": "jira_project",
  "Jira/Issue": "jira_issue",
  "Jira/Board": "jira_board",
  "Jira/Group": "jira_group",
  "Jira/User": "jira_user",
}

var ResourceTypeToDescription = map[string]interface{}{
  "Jira/Project": opengovernance.Project{},
  "Jira/Issue": opengovernance.Issue{},
  "Jira/Board": opengovernance.Board{},
  "Jira/Group": opengovernance.Group{},
  "Jira/User": opengovernance.User{},
}

var TablesToResourceTypes = map[string]string{
  "jira_project": "Jira/Project",
  "jira_issue": "Jira/Issue",
  "jira_board": "Jira/Board",
  "jira_group": "Jira/Group",
  "jira_user": "Jira/User",
}
