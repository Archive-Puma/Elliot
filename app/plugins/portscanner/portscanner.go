package portscanner

import (
	"errors"
	"fmt"

	"github.com/cosasdepuma/elliot/app/env"
	"github.com/cosasdepuma/elliot/app/validator"
)

// Plugin TODO: Doc
type Plugin struct{}

// Check TODO: Doc
func (plugin Plugin) check() error {
	if !validator.IsValidDomain(env.Params.Target) {
		return errors.New("A valid domain should be specified")
	}
	if len(env.Params.Arguments) == 0 {
		return errors.New("Valid parameters should be specified")
	}
	ports, ok := validator.ParsePorts(env.Params.Arguments[0].(string))
	if !ok {
		return errors.New("Valid ports should be specified")
	}
	env.Params.Arguments[0] = ports
	return nil
}

// Run TODO: Doc
func (plugin Plugin) Run() {
	if err := plugin.Check(); err != nil {
		env.Channels.Bad <- err
	}
	/*
		results := make([]string, 0)
		errors := make([]*error.MrRobotError, 0)

		scanner := NewPortScanner(config.Args.Domain)

		for _, port := range config.Args.Ports {
			result := scanner.CheckTCPPort(port)
			fmt.Println(result.String())
		}

		return results, errors
	*/
	env.Channels.Ok <- fmt.Sprintln("Test done: ", env.Params.Arguments)
}
