package main

import (
	"sync"

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

// registering a slash command /hello
func (p *Plugin) OnActivate() error {

	if err := p.registerCommands(); err != nil {
		return errors.Wrapf(err, "failed to register  command")
	}
	return nil
}
