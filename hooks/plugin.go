package hooks

import (
	"github.com/mattermost/mattermost-server/v5/plugin"
)

// NopPlugin is the main plugin object.
type NopPlugin struct {
	plugin.MattermostPlugin
	// nopPluginHooks contains empty implementations for all defined hook methods.
	nopPluginHooks
}

// Check that the `plugin.Hooks` interface is implemented by `nopPluginHooks`.
var _ plugin.Hooks = (*nopPluginHooks)(nil)

// Generate the implementation.
//go:generate go run ../tools/generate.go -interface github.com/mattermost/mattermost-server/v5/plugin.Hooks -pkg hooks -struct nopPluginHooks -output hooks_generated.go
