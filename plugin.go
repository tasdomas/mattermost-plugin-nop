package main

import (
	"github.com/mattermost/mattermost-server/v5/plugin"

	"github.com/tasdomas/mattermost-plugin-nop/hooks"
)

func main() {
	plugin.ClientMain(&hooks.NopPlugin{})
}
