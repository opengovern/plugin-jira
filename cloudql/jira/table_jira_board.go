package jira

import (
	"context"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableJiraBoard(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "jira_board",
		Description: "Jira boards information.",
		List: &plugin.ListConfig{
			Hydrate: nil,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    nil,
		},
		Columns: commonColumns([]*plugin.Column{
			{
				Name:        "id",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.ID"),
				Description: "The unique identifier of the board.",
			},
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Name"),
				Description: "The name of the board.",
			},
			{
				Name:        "self",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Self"),
				Description: "The API URL of the board.",
			},
			{
				Name:        "type",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Type"),
				Description: "The type of the board (e.g., scrum, kanban).",
			},
		}),
	}
}
