package portscanner

import (
	"fmt"

	"github.com/cosasdepuma/elliot/app/error"
	"github.com/cosasdepuma/elliot/app/validator"
)

// Plugin TODO: Doc
type Plugin struct{}

// Check TODO: Doc
func (s Plugin) Check() *error.MrRobotError {
	if validator.IsValidDomain(config.Args.Domain) {
		return error.NewWarning("A valid domain should be specified")
	}

	ports, ok := validator.ParsePorts(config.Args.RawPorts)
	if !ok {
		return error.NewWarning("Valid ports should be specified")
	}

	return nil
}

// Run TODO: Doc
func (s Plugin) Run() ([]string, []*error.MrRobotError) {
	results := make([]string, 0)
	errors := make([]*error.MrRobotError, 0)

	scanner := NewPortScanner(config.Args.Domain)

	for _, port := range config.Args.Ports {
		result := scanner.CheckTCPPort(port)
		fmt.Println(result.String())
	}

	return results, errors
}
