package hooks

import (
	"io"
	"net/http"

	"github.com/mattermost/mattermost-server/v5/model"
	"github.com/mattermost/mattermost-server/v5/plugin"
)

func (*nopPluginHooks) OnActivate() error {
	return nil
}

func (*nopPluginHooks) Implemented() ([]string, error) {
	return nil, nil
}

func (*nopPluginHooks) OnDeactivate() error {
	return nil
}

func (*nopPluginHooks) OnConfigurationChange() error {
	return nil
}

func (*nopPluginHooks) ServeHTTP(c *plugin.Context, w http.ResponseWriter, r *http.Request) {
	return
}

func (*nopPluginHooks) ExecuteCommand(c *plugin.Context, args *model.CommandArgs) (*model.CommandResponse, *model.AppError) {
	return nil, nil
}

func (*nopPluginHooks) UserHasBeenCreated(c *plugin.Context, user *model.User) {
	return
}

func (*nopPluginHooks) UserWillLogIn(c *plugin.Context, user *model.User) string {
	return ""
}

func (*nopPluginHooks) UserHasLoggedIn(c *plugin.Context, user *model.User) {
	return
}

func (*nopPluginHooks) MessageWillBePosted(c *plugin.Context, post *model.Post) (*model.Post, string) {
	return nil, ""
}

func (*nopPluginHooks) MessageWillBeUpdated(c *plugin.Context, newPost, oldPost *model.Post) (*model.Post, string) {
	return nil, ""
}

func (*nopPluginHooks) MessageHasBeenPosted(c *plugin.Context, post *model.Post) {
	return
}

func (*nopPluginHooks) MessageHasBeenUpdated(c *plugin.Context, newPost, oldPost *model.Post) {
	return
}

func (*nopPluginHooks) ChannelHasBeenCreated(c *plugin.Context, channel *model.Channel) {
	return
}

func (*nopPluginHooks) UserHasJoinedChannel(c *plugin.Context, channelMember *model.ChannelMember, actor *model.User) {
	return
}

func (*nopPluginHooks) UserHasLeftChannel(c *plugin.Context, channelMember *model.ChannelMember, actor *model.User) {
	return
}

func (*nopPluginHooks) UserHasJoinedTeam(c *plugin.Context, teamMember *model.TeamMember, actor *model.User) {
	return
}

func (*nopPluginHooks) UserHasLeftTeam(c *plugin.Context, teamMember *model.TeamMember, actor *model.User) {
	return
}

func (*nopPluginHooks) FileWillBeUploaded(c *plugin.Context, info *model.FileInfo, file io.Reader, output io.Writer) (*model.FileInfo, string) {
	return nil, ""
}
