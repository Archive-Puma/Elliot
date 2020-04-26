package robots

import (
	"strings"

	"github.com/cosasdepuma/elliot/app/config"
	"github.com/cosasdepuma/elliot/app/error"
	"github.com/cosasdepuma/elliot/app/tui"
	"github.com/cosasdepuma/elliot/app/validator"

	"fmt"
	"io/ioutil"
	"net/http"
)

// Subcommand TODO: Doc
type Subcommand struct{}

// Check TODO: Doc
func (s Subcommand) Check() *error.MrRobotError {
	if validator.IsValidURL(config.Args.URL) {
		tui.PrintInfo("URL", config.Args.URL)
	} else {
		return error.NewWarning("A valid URL should be specified")
	}

	return nil
}

// Run TODO: Doc
func (s Subcommand) Run() []*error.MrRobotError {
	errors := []*error.MrRobotError{}

	path := fmt.Sprintf("%s/robots.txt", config.Args.URL)
	resp, err := http.Get(path)
	if err != nil {
		return append(errors, error.NewWarning("Robots.txt not found"))
	}

	defer resp.Body.Close()
	bRobots, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return append(errors, error.NewWarning("Cannot read robots.txt"))
	}

	robots := strings.TrimSpace(string(bRobots))

	fmt.Println(robots)

	return errors
}
