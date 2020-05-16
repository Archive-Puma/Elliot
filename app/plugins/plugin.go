package plugins

import (
	"github.com/cosasdepuma/elliot/app/error"
)

// Plugins TODO: Doc
var (
	Plugins = map[string]interface {
		Help()
		Check() *error.MrRobotError
		Run() ([]string, []*error.MrRobotError)
	}{
		"portscanner": nil,
		"subdomain":   nil,
		"robots":      nil,
	}
	Amount = len(Plugins)
)
