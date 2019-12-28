/*
Package main implements the plugin's executable.

The nop plugin has empty implementations for all plugin
hooks available. It is meant for testing purposes.
*/
package main

import (
	"github.com/mattermost/mattermost-server/v5/plugin"

	"github.com/tasdomas/mattermost-plugin-nop/hooks"
)

func main() {
	plugin.ClientMain(&hooks.NopPlugin{})
}
