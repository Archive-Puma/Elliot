package plugins

import (
	"github.com/cosasdepuma/elliot/app/error"
	//"github.com/cosasdepuma/elliot/app/plugins/portscanner"
	//"github.com/cosasdepuma/elliot/app/plugins/robots"
	//"github.com/cosasdepuma/elliot/app/plugins/subdomain"
)

// Plugin TODO: Doc
type Plugin interface {
	Help()
	Check() *error.MrRobotError
	Run() ([]string, []*error.MrRobotError)
}

// Plugins TODO: Doc
var Plugins = map[string]Plugin{
	"portscanner": nil, //portscanner.Plugin{},
	"subdomain":   nil, //subdomain.Plugin{},
	"robots":      nil, //robots.Plugin{},
}
