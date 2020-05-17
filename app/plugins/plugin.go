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
		// "portscanner": nil,
		"subdomain":  subdomain.Plugin{},
		"robots.txt": robots.Plugin{},
	}
	Amount = len(Plugins)
)
