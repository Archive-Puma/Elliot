package robots

import (
	"github.com/cosasdepuma/elliot/app/error"

	"fmt"
	"io/ioutil"
	"net/http"
)

// FindRobots TODO: Doc
func FindRobots(url string) (string, *error.MrRobotError) {
	path := fmt.Sprintf("%s/robots.txt", url)
	resp, err := http.Get(path)
	if err != nil {
		return "", error.NewWarning("Robots.txt not found")
	}

	defer resp.Body.Close()
	robots, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", error.NewWarning("Cannot read robots.txt")
	}

	return string(robots), nil
}
