package jira

import (
	"context"
	opengovernance "github.com/opengovern/og-describer-jira/discovery/pkg/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableJiraGroup(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "jira_group",
		Description: "Jira groups information.",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListGroup,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("group_id"),
			Hydrate:    opengovernance.GetGroup,
		},
		Columns: integrationColumns([]*plugin.Column{
			{
				Name:        "group_id",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.GroupID"),
				Description: "The unique identifier of the group.",
			},
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Name"),
				Description: "The name of the group.",
			},
		}),
	}
}
