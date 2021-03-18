package main

import (
	"fmt"
	"strings"
	"sync"

	"github.com/mattermost/mattermost-server/v5/model"
	"github.com/mattermost/mattermost-server/v5/plugin"
	"github.com/pkg/errors"
)

// Plugin implements the interface expected by the Mattermost server to communicate between the server and plugin processes.
type Plugin struct {
	plugin.MattermostPlugin

	// configurationLock synchronizes access to the configuration.
	configurationLock sync.RWMutex

	// configuration is the active plugin configuration. Consult getConfiguration and
	// setConfiguration for usage.
	configuration *configuration
}

// registering a slashcommand /hello
func (p *Plugin) OnActivate() error {

	if err := p.API.RegisterCommand(&model.Command{
		Trigger:          "hello",
		AutoComplete:     true,
		AutoCompleteHint: "",
		AutoCompleteDesc: "responds world",
	}); err != nil {
		return errors.Wrapf(err, "failed to register  command")
	}
	return nil

}

// adding functionality to to slashcommand
func (p *Plugin) ExecuteCommand(c *plugin.Context, args *model.CommandArgs) (*model.CommandResponse, *model.AppError) {

	// reading the slashcommand
	trigger := strings.TrimPrefix(strings.Fields(args.Command)[0], "/")
	switch trigger {
	case "hello":
		return &model.CommandResponse{
			ResponseType: model.COMMAND_RESPONSE_TYPE_EPHEMERAL,
			Text:         fmt.Sprintf("World"),
		}, nil

	default:
		return &model.CommandResponse{
			ResponseType: model.COMMAND_RESPONSE_TYPE_EPHEMERAL,
			Text:         fmt.Sprintf("Unknown__ command: " + args.Command),
		}, nil
	}
}
