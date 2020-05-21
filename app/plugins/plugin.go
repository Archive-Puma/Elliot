package plugins

import (
	"github.com/cosasdepuma/elliot/app/plugins/robots"
	"github.com/cosasdepuma/elliot/app/plugins/subdomain"
)

// Plugins TODO: Doc
var (
	Plugins = map[string]interface {
		Check() error
		Run()
	}{
		// "portscanner": portscanner.Plugin{},
		"robots.txt": robots.Plugin{},
		"subdomain":  subdomain.Plugin{},
	}
	Required = map[string]string{
		"portscanner": "Ports",
	}
	Amount = len(Plugins)
)
