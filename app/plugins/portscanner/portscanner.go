package portscanner

import (
	"errors"
	"fmt"
	"sort"
	"strings"

	"github.com/cosasdepuma/elliot/app/env"
	"github.com/cosasdepuma/elliot/app/validator"
)

// Plugin TODO: Doc
type Plugin struct{}

// Check TODO: Doc
func (plugin Plugin) Check() error {
	if !validator.IsValidDomain(env.Config.Target) {
		return errors.New("A valid domain should be specified")
	}
	if len(env.Config.Params.(string)) == 0 {
		return errors.New("Valid parameters should be specified")
	}
	ports, ok := validator.ParsePorts(env.Config.Params.(string))
	if !ok {
		return errors.New("Valid ports should be specified")
	}
	env.Config.Params = ports
	return nil
}

// Run TODO: Doc
func (plugin Plugin) Run() {
	if err := plugin.Check(); err != nil {
		env.Channels.Bad <- err
		return
	}

	results := map[int]string{}
	channel := make(chan *Port, len(env.Config.Params.([]int)))

	scanner := NewPortScanner(env.Config.Target)

	for _, port := range env.Config.Params.([]int) {
		go func(port int) { channel <- scanner.CheckTCPPort(port) }(port)
	}

	progress := len(env.Config.Params.([]int))
	for progress > 0 {
		progress--
		result := <-channel
		if result.IsOpen() {
			results[result.number] = result.String()
		}
	}

	sorted := []int{}
	for index := range results {
		sorted = append(sorted, index)
	}
	sort.Ints(sorted)

	buffer := fmt.Sprintf("%9s\t%-7s\t%-9s\t%s\n", "PORT", "STATE", "SERVICE", "BANNER")
	for _, port := range sorted {
		buffer = fmt.Sprintf("%s%s\n", buffer, results[port])
	}

	env.Channels.Ok <- strings.TrimSuffix(buffer, "\n")
}
