package hooks

import (
	"io"
	"net/http"

	"github.com/mattermost/mattermost-server/v5/model"
	"github.com/mattermost/mattermost-server/v5/plugin"
)

// OnActivate is part of the implementation of the `plugin.Hooks` interface.
func (nopPluginHooks) OnActivate() error {
	return nil
}

// Implemented is part of the implementation of the `plugin.Hooks` interface.
func (nopPluginHooks) Implemented() ([]string, error) {
	return nil, nil
}

// OnDeactivate is part of the implementation of the `plugin.Hooks` interface.
func (nopPluginHooks) OnDeactivate() error {
	return nil
}

// OnConfigurationChange is part of the implementation of the `plugin.Hooks` interface.
func (nopPluginHooks) OnConfigurationChange() error {
	return nil
}

// ServeHTTP is part of the implementation of the `plugin.Hooks` interface.
func (nopPluginHooks) ServeHTTP(c *plugin.Context, w http.ResponseWriter, r *http.Request) {
	return
}

// ExecuteCommand is part of the implementation of the `plugin.Hooks` interface.
func (nopPluginHooks) ExecuteCommand(c *plugin.Context, args *model.CommandArgs) (*model.CommandResponse, *model.AppError) {
	return nil, nil
}

// UserHasBeenCreated is part of the implementation of the `plugin.Hooks` interface.
func (nopPluginHooks) UserHasBeenCreated(c *plugin.Context, user *model.User) {
	return
}

// UserWillLogIn is part of the implementation of the `plugin.Hooks` interface.
func (nopPluginHooks) UserWillLogIn(c *plugin.Context, user *model.User) string {
	return ""
}

// UserHasLoggedIn is part of the implementation of the `plugin.Hooks` interface.
func (nopPluginHooks) UserHasLoggedIn(c *plugin.Context, user *model.User) {
	return
}

// MessageWillBePosted is part of the implementation of the `plugin.Hooks` interface.
func (nopPluginHooks) MessageWillBePosted(c *plugin.Context, post *model.Post) (*model.Post, string) {
	return nil, ""
}

// MessageWillBeUpdated is part of the implementation of the `plugin.Hooks` interface.
func (nopPluginHooks) MessageWillBeUpdated(c *plugin.Context, newPost, oldPost *model.Post) (*model.Post, string) {
	return nil, ""
}

// MessageHasBeenPosted is part of the implementation of the `plugin.Hooks` interface.
func (nopPluginHooks) MessageHasBeenPosted(c *plugin.Context, post *model.Post) {
	return
}

// MessageHasBeenUpdated is part of the implementation of the `plugin.Hooks` interface.
func (nopPluginHooks) MessageHasBeenUpdated(c *plugin.Context, newPost, oldPost *model.Post) {
	return
}

// ChannelHasBeenCreated is part of the implementation of the `plugin.Hooks` interface.
func (nopPluginHooks) ChannelHasBeenCreated(c *plugin.Context, channel *model.Channel) {
	return
}

// UserHasJoinedChannel is part of the implementation of the `plugin.Hooks` interface.
func (nopPluginHooks) UserHasJoinedChannel(c *plugin.Context, channelMember *model.ChannelMember, actor *model.User) {
	return
}

// UserHasLeftChannel is part of the implementation of the `plugin.Hooks` interface.
func (nopPluginHooks) UserHasLeftChannel(c *plugin.Context, channelMember *model.ChannelMember, actor *model.User) {
	return
}

// UserHasJoinedTeam is part of the implementation of the `plugin.Hooks` interface.
func (nopPluginHooks) UserHasJoinedTeam(c *plugin.Context, teamMember *model.TeamMember, actor *model.User) {
	return
}

// UserHasLeftTeam is part of the implementation of the `plugin.Hooks` interface.
func (nopPluginHooks) UserHasLeftTeam(c *plugin.Context, teamMember *model.TeamMember, actor *model.User) {
	return
}

// FileWillBeUploaded is part of the implementation of the `plugin.Hooks` interface.
func (nopPluginHooks) FileWillBeUploaded(c *plugin.Context, info *model.FileInfo, file io.Reader, output io.Writer) (*model.FileInfo, string) {
	return nil, ""
}
