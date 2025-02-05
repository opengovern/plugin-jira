package main

import (
	"github.com/opengovern/og-describer-jira/cloudql/jira"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{PluginFunc: jira.Plugin})
}
