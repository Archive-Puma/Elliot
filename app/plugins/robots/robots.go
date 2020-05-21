package robots

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/cosasdepuma/elliot/app/env"
	"github.com/cosasdepuma/elliot/app/validator"
)

// Plugin allows it to be executed by Elliot
type Plugin struct{}

// Check that all parameters are defined correctly
func (plgn Plugin) Check() error {
	if !validator.IsValidURL(env.Config.Target) {
		return errors.New("A valid URL should be specified")
	}

	return nil
}

// Run is the entrypoint of the plugin
func (plgn Plugin) Run() {
	if err := plgn.Check(); err != nil {
		env.Channels.Bad <- err
		return
	}

	path := fmt.Sprintf("%s/robots.txt", env.Config.Target)
	resp, err := http.Get(path)
	if err != nil {
		env.Channels.Bad <- errors.New("Robots.txt not found")
		return
	}
	defer resp.Body.Close()

	bRobots, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		env.Channels.Bad <- errors.New("Cannot read robots.txt")
		return
	}
	results := strings.TrimSpace(string(bRobots))
	env.Channels.Ok <- results
}
