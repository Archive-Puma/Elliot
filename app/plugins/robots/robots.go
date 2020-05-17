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

// Plugin TODO: Doc
type Plugin struct{}

// Check TODO: Doc
func (plgn Plugin) Check() error {
	if !validator.IsValidURL(env.Params.Target) {
		return errors.New("A valid URL should be specified")
	}

	return nil
}

// Run TODO: Doc
func (plgn Plugin) Run() ([]string, error) {
	if err := plgn.Check(); err != nil {
		return nil, err
	}

	path := fmt.Sprintf("%s/robots.txt", env.Params.Target)
	resp, err := http.Get(path)
	if err != nil {
		return nil, errors.New("Robots.txt not found")
	}
	defer resp.Body.Close()

	bRobots, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("Cannot read robots.txt")
	}

	results := strings.Split(strings.TrimSpace(string(bRobots)), "\n")

	// TODO: Implement disallow only
	// results = filterDisallow(results)

	// TODO: Implement extended
	// results = extendedMode(results)

	return results, nil
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
		if strings.HasPrefix(robot, "Allow: ") || strings.HasPrefix(robot, "Disallow: ") || strings.HasPrefix(robot, "/") {
			splits := strings.SplitN(robot, "/", 2)
			url := env.Params.Target
			if !strings.HasSuffix(url, "/") {
				url = fmt.Sprintf("%s%s", url, "/")
			}
			robot = fmt.Sprintf("%s%s%s", splits[0], url, splits[1])
		}
		extended = append(extended, robot)
	}
	return extended
}
