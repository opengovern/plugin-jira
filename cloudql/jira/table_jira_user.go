package jira

import (
	"context"
	opengovernance "github.com/opengovern/og-describer-jira/discovery/pkg/es"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableJiraUser(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "jira_user",
		Description: "Jira users information.",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListUser,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("account_id"),
			Hydrate:    opengovernance.GetUser,
		},
		Columns: integrationColumns([]*plugin.Column{
			{
				Name:        "account_id",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("AccountID"),
				Description: "The unique account identifier of the user.",
			},
			{
				Name:        "account_type",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("AccountType"),
				Description: "The type of account (e.g., atlassian).",
			},
			{
				Name:        "active",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Active"),
				Description: "Indicates whether the user is active.",
			},
			{
				Name:        "display_name",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("DisplayName"),
				Description: "The display name of the user.",
			},
			{
				Name:        "key",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Key"),
				Description: "The user key (deprecated in Jira Cloud).",
			},
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Name"),
				Description: "The username (deprecated in Jira Cloud).",
			},
			{
				Name:        "self",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Self"),
				Description: "The URL of the user's REST API endpoint.",
			},
			{
				Name:        "avatar_urls",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("AvatarURLs"),
				Description: "The avatar URLs of the user.",
			},
		}),
	}
}
