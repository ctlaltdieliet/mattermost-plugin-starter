package main

import (
	"fmt"
	"strings"

	"github.com/mattermost/mattermost-server/v5/model"
	"github.com/mattermost/mattermost-server/v5/plugin"
	"github.com/pkg/errors"
)

func (p *Plugin) registerCommands() error {
	if err := p.API.RegisterCommand(&model.Command{
		Trigger:          "hello",
		AutoComplete:     true,
		AutoCompleteDesc: "responds world",
	}); err != nil {
		return errors.Wrapf(err, "failed to register  command")
	}

	return nil
}

func (p *Plugin) ExecuteCommand(c *plugin.Context, args *model.CommandArgs) (*model.CommandResponse, *model.AppError) {
	trigger := strings.TrimPrefix(strings.Fields(args.Command)[0], "/")
	username := ""
	switch trigger {
	case "hello":
		UserId := args.UserId

		user, appErr := p.API.GetUser(UserId)
		if appErr != nil {
			username = "there"
		} else {
			username = user.GetFullName()
		}

		//fmt.Sprintf(Apperr.Message)
		return &model.CommandResponse{
			ResponseType: model.COMMAND_RESPONSE_TYPE_EPHEMERAL,
			Text:         fmt.Sprintf("Hi " + username),
		}, nil

	default:
		return &model.CommandResponse{
			ResponseType: model.COMMAND_RESPONSE_TYPE_EPHEMERAL,
			Text:         fmt.Sprintf("Unknown__ command: " + args.Command),
		}, nil
	}
}

func (p *Plugin) executeAutocompleteTest(args *model.CommandArgs) *model.CommandResponse {
	return &model.CommandResponse{
		ResponseType: model.COMMAND_RESPONSE_TYPE_EPHEMERAL,
		Text:         fmt.Sprintf("World"),
	}
}
