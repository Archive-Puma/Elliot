package robots

import (
	"github.com/cosasdepuma/elliot/app/config"
	"github.com/cosasdepuma/elliot/app/error"
	"github.com/cosasdepuma/elliot/app/tui"
	"github.com/cosasdepuma/elliot/app/validator"

	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// Subcommand TODO: Doc
type Subcommand struct{}

// Help TODO: Doc
func (s Subcommand) Help() {
	config.Args.Print("-extended", "Full links")
	config.Args.Print("-disallow", "Filter only the disallowed links")
	config.Args.Print("-u, -url", "Target URL")
}

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
func (s Subcommand) Run() ([]string, []*error.MrRobotError) {
	results := make([]string, 0)
	errors := make([]*error.MrRobotError, 0)

	path := fmt.Sprintf("%s/robots.txt", config.Args.URL)
	resp, err := http.Get(path)
	if err != nil {
		return nil, append(errors, error.NewWarning("Robots.txt not found"))
	}

	defer resp.Body.Close()
	bRobots, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, append(errors, error.NewWarning("Cannot read robots.txt"))
	}

	results = strings.Split(strings.TrimSpace(string(bRobots)), "\n")

	if config.Args.Disallow {
		results = filterDisallow(results)
	}
	if config.Args.Extended {
		results = extendedMode(results)
	}

	return results, errors
}

func filterDisallow(robots []string) []string {
	disallowed := make([]string, 0)
	for _, robot := range robots {
		if strings.HasPrefix(robot, "Disallow: ") {
			disallowed = append(disallowed, strings.TrimSpace(strings.SplitN(robot, ": ", 2)[1]))
		}
	}
	return disallowed
}

func extendedMode(robots []string) []string {
	extended := make([]string, 0)
	for _, robot := range robots {
		if strings.HasPrefix(robot, "Allow: ") || strings.HasPrefix(robot, "Disallow: ") {
			splits := strings.SplitN(robot, "/", 2)
			url := config.Args.URL
			if !strings.HasSuffix(url, "/") {
				url = fmt.Sprintf("%s%s", url, "/")
			}
			robot = fmt.Sprintf("%s%s%s", splits[0], url, splits[1])
		}
		extended = append(extended, robot)
	}
	return extended
}
