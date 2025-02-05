package jira

import (
	"context"
	opengovernance "github.com/opengovern/og-describer-jira/discovery/pkg/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableJiraProject(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "jira_project",
		Description: "Jira projects information.",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListProject,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    opengovernance.GetProject,
		},
		Columns: integrationColumns([]*plugin.Column{
			{
				Name:        "id",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ID"),
				Description: "The unique identifier of the project.",
			},
			{
				Name:        "key",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Key"),
				Description: "The key of the project.",
			},
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Name"),
				Description: "The name of the project.",
			},
			{
				Name:        "self",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Self"),
				Description: "The API URL of the project.",
			},
			{
				Name:        "simplified",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Simplified"),
				Description: "Indicates whether the project is simplified.",
			},
			{
				Name:        "style",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Style"),
				Description: "The style of the project (e.g., classic, next-gen).",
			},
			{
				Name:        "avatar_urls",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.AvatarUrls"),
				Description: "Avatar URLs for different sizes of the project icon.",
			},
			{
				Name:        "insight",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Insight"),
				Description: "Insight data related to the project.",
			},
			{
				Name:        "project_category",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ProjectCategory"),
				Description: "Category information for the project.",
			},
		}),
	}
}
