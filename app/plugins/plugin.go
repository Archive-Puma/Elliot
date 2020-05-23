package plugins

import (
	"github.com/cosasdepuma/elliot/app/plugins/commoncrawler"
	"github.com/cosasdepuma/elliot/app/plugins/portscanner"
	"github.com/cosasdepuma/elliot/app/plugins/robots"
	"github.com/cosasdepuma/elliot/app/plugins/subdomainer"
)

var (
	// Plugins collects all available plugins that can be run
	Plugins = map[string]interface {
		Run()
		Check() error
		Save([]string) error
	}{
		"commoncrawler": new(commoncrawler.Plugin),
		"portscanner":   new(portscanner.Plugin),
		"robots.txt":    new(robots.Plugin),
		"subdomainer":   new(subdomainer.Plugin),
	}
	// Required specifies what parameters are necessary to run a plugin
	Required = map[string]string{
		"portscanner": "Ports",
	}
	// Amount specifies how many plugins are available
	Amount = len(Plugins)
)
