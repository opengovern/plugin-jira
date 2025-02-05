package jira

import (
	"context"
	opengovernance "github.com/opengovern/og-describer-jira/discovery/pkg/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableJiraIssue(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "jira_issue",
		Description: "Jira issues information.",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListIssue,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    opengovernance.GetBoard,
		},
		Columns: integrationColumns([]*plugin.Column{
			{
				Name:        "id",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ID"),
				Description: "The unique identifier of the issue.",
			},
			{
				Name:        "key",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Key"),
				Description: "The key of the issue.",
			},
			{
				Name:        "self",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Self"),
				Description: "The API URL of the issue.",
			},
			{
				Name:        "expand",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Expand"),
				Description: "Additional expand options for the issue.",
			},
			{
				Name:        "fields",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Fields"),
				Description: "Detailed fields related to the issue.",
			},
		}),
	}
}
