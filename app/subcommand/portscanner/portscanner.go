package portscanner

import (
	"fmt"

	"github.com/cosasdepuma/elliot/app/cli"
	"github.com/cosasdepuma/elliot/app/config"
	"github.com/cosasdepuma/elliot/app/error"
	"github.com/cosasdepuma/elliot/app/validator"
)

// Subcommand TODO: Doc
type Subcommand struct{}

// Help TODO: Doc
func (s Subcommand) Help() {
	config.Args.Print("-d, -domain", "Target domain")
	config.Args.Print("-p, -port, -ports", "Target ports")
	config.Args.Print("-t, -timeout", "Connection timeout")
}

// Check TODO: Doc
func (s Subcommand) Check() *error.MrRobotError {
	if validator.IsValidDomain(config.Args.Domain) {
		cli.PrintInfo("Domain", config.Args.Domain)
	} else {
		return error.NewWarning("A valid domain should be specified")
	}

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

	scanner := NewPortScanner(config.Args.Domain)

	for _, port := range config.Args.Ports {
		result := scanner.CheckTCPPort(port)
		fmt.Println(result.String())
	}

	return results, errors
}
