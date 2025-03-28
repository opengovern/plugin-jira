package maps
import (
	"github.com/opengovern/og-describer-jira/discovery/describers"
	"github.com/opengovern/og-describer-jira/discovery/provider"
	"github.com/opengovern/og-describer-jira/platform/constants"
	"github.com/opengovern/og-util/pkg/integration/interfaces"
	model "github.com/opengovern/og-describer-jira/discovery/pkg/models"
)
var ResourceTypes = map[string]model.ResourceType{

	"Jira/Project": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Jira/Project",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeListByJira(describers.ListProjects),
		GetDescriber:         provider.DescribeSingleByJira(describers.GetProject),
	},

	"Jira/Issue": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Jira/Issue",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeListByJira(describers.ListIssues),
		GetDescriber:         provider.DescribeSingleByJira(describers.GetIssue),
	},

	"Jira/Board": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Jira/Board",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeListByJira(describers.ListBoards),
		GetDescriber:         provider.DescribeSingleByJira(describers.GetBoard),
	},

	"Jira/Group": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Jira/Group",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeListByJira(describers.ListGroups),
		GetDescriber:         nil,
	},

	"Jira/User": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Jira/User",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeListByJira(describers.ListUsers),
		GetDescriber:         nil,
	},
}


var ResourceTypeConfigs = map[string]*interfaces.ResourceTypeConfiguration{

	"Jira/Project": {
		Name:         "Jira/Project",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		Params:           	[]interfaces.Param{
			{
				Name:  "project_key",
				Description: "Please provide the project key",
				Required:    true,
				Default:     nil,
			},
			      },
		
	},

	"Jira/Issue": {
		Name:         "Jira/Issue",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		Params:           	[]interfaces.Param{
			{
				Name:  "fields",
				Description: "Please provide the fields",
				Required:    false,
				Default:     nil,
			},
			
			{
				Name:  "project_key",
				Description: "Please provide the project key",
				Required:    true,
				Default:     nil,
			},
			
			{
				Name:  "status",
				Description: "Please provide the status",
				Required:    true,
				Default:     nil,
			},
			
			{
				Name:  "status_category",
				Description: "Please provide the status category",
				Required:    true,
				Default:     nil,
			},
			      },
		
	},

	"Jira/Board": {
		Name:         "Jira/Board",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Jira/Group": {
		Name:         "Jira/Group",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Jira/User": {
		Name:         "Jira/User",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},
}


var ResourceTypesList = []string{
  "Jira/Project",
  "Jira/Issue",
  "Jira/Board",
  "Jira/Group",
  "Jira/User",
}