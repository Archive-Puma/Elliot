package portscanner

import (
	"github.com/cosasdepuma/elliot/app/cli"
	"github.com/cosasdepuma/elliot/app/config"
	"github.com/cosasdepuma/elliot/app/error"
	"github.com/cosasdepuma/elliot/app/validator"

	"fmt"
)

// Subcommand TODO: Doc
type Subcommand struct{}

// Help TODO: Doc
func (s Subcommand) Help() {
	config.Args.Print("-p, -port, -ports", "Target ports")
	config.Args.Print("-t, -timeout", "Connection timeout")
}

// Check TODO: Doc
func (s Subcommand) Check() *error.MrRobotError {
	ports, ok := validator.ParsePorts(config.Args.RawPorts)

	if ok {
		cli.PrintInfo("Port(s)", config.Args.RawPorts)
		config.Args.Ports = ports
	} else {
		return error.NewWarning("Valid ports should be specified")
	}

	return nil
}

// Run TODO: Doc
func (s Subcommand) Run() ([]string, []*error.MrRobotError) {
	results := make([]string, 0)
	errors := make([]*error.MrRobotError, 0)

	for _, p := range config.Args.Ports {
		fmt.Println(p)
	}

	return results, errors
}
