package portscanner

import (
	"errors"
	"fmt"
	"sort"
	"strings"

	"github.com/cosasdepuma/elliot/app/env"
	"github.com/cosasdepuma/elliot/app/validator"
	"github.com/sirupsen/logrus"
)

// Plugin allows it to be executed by Elliot
type Plugin struct{}

// Check that all parameters are defined correctly
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

// Run is the entrypoint of the plugin
func (plugin Plugin) Run() {
	if err := plugin.Check(); err != nil {
		env.Channels.Bad <- err
		return
	}

	results := map[int]*sPort{}
	channel := make(chan *sPort, len(env.Config.Params.([]int)))

	scanner := newPortScanner(env.Config.Target)

	for _, port := range env.Config.Params.([]int) {
		go func(port int) { channel <- scanner.checkTCPPort(port) }(port)
	}

	progress := len(env.Config.Params.([]int))
	for progress > 0 {
		progress--
		result := <-channel
		if result.isOpen() {
			results[result.number] = result
		}
	}

	sorted := []int{}
	for index := range results {
		sorted = append(sorted, index)
	}
	sort.Ints(sorted)

	slice := make([]string, 0)
	buffer := fmt.Sprintf("%9s\t%-7s\t%-9s\t%s\n", "PORT", "STATE", "SERVICE", "BANNER")
	for _, port := range sorted {
		slice = append(slice, results[port].format())
		buffer = fmt.Sprintf("%s%s\n", buffer, results[port].string())
	}
	env.Channels.Ok <- strings.TrimSuffix(buffer, "\n")
	if err := plugin.Save(slice); err != nil {
		logrus.Error(err)
	}
}
