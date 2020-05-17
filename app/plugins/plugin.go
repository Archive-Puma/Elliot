package plugins

import "github.com/cosasdepuma/elliot/app/plugins/robots"

// Plugins TODO: Doc
var (
	Plugins = map[string]interface {
		Check() error
		Run()
	}{
		"portscanner": nil,
		"subdomain":   nil,
		"robots":      robots.Plugin{},
	}
	Amount = len(Plugins)
)
