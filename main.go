package main

import (
	"github.com/raw-labs/steampipe-plugin-confluence/confluence"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{PluginFunc: confluence.Plugin})
}
