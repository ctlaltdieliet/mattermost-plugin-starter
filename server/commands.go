package main

import (
	"fmt"
	"strings"

	"github.com/mattermost/mattermost-server/v5/model"
	"github.com/mattermost/mattermost-server/v5/plugin"
	"github.com/pkg/errors"
)

func (p *Plugin) registerCommands() error {
	commandName := p.getConfiguration().getCommandName()

	if err := p.API.RegisterCommand(&model.Command{
		Trigger:          commandName,
		AutoComplete:     true,
		AutoCompleteDesc: "responds world",
	}); err != nil {
		return errors.Wrapf(err, "failed to register  command")
	}

	return nil
}

func (p *Plugin) ExecuteCommand(c *plugin.Context, args *model.CommandArgs) (*model.CommandResponse, *model.AppError) {
	trigger := strings.TrimPrefix(strings.Fields(args.Command)[0], "/")
	commandName := p.getConfiguration().getCommandName()
	username := ""
	UserId := args.UserId
	user, appErr := p.API.GetUser(UserId)
	if appErr == nil {
		username = p.getConfiguration().DefaultGreeting
	}
	if username == "" {
		username = p.getConfiguration().DefaultGreeting

	}

	switch trigger {
	case commandName:
		return &model.CommandResponse{
			ResponseType: model.COMMAND_RESPONSE_TYPE_EPHEMERAL,
			Text:         fmt.Sprintf("Hello " + username),
		}, nil

	default:
		return &model.CommandResponse{
			ResponseType: model.COMMAND_RESPONSE_TYPE_EPHEMERAL,
			Text:         fmt.Sprintf("Unknown command: " + args.Command),
		}, nil
	}
}
