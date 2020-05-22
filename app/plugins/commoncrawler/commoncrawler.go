package commoncrawler

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/cosasdepuma/elliot/app/env"
	"github.com/cosasdepuma/elliot/app/validator"
)

// Plugin allows it to be executed by Elliot
type Plugin struct{}

// Check that all parameters are defined correctly
func (plgn Plugin) Check() error {
	if !validator.IsValidURL(env.Config.Target) && !validator.IsValidDomain(env.Config.Target) {
		return errors.New("A valid domain or URL should be specified")
	}

	return nil
}

// Run is the entrypoint of the plugin
func (plgn Plugin) Run() {
	// Check the parameters
	if err := plgn.Check(); err != nil {
		env.Channels.Bad <- err
		return
	}
	// Compose the API
	indexYear := "2020-16" // This value should be periodically updated
	api := fmt.Sprintf("http://index.commoncrawl.org/CC-MAIN-%s-index?output=json&url=%s/*", indexYear, env.Config.Target)
	// Request the data
	resp, err := http.Get(api)
	if err != nil || resp.StatusCode != 200 {
		env.Channels.Bad <- errors.New("CommonCrawl is not available")
		return
	}
	defer resp.Body.Close()
	// Grab the content
	results := ""
	sc := bufio.NewScanner(resp.Body)
	for sc.Scan() {
		link := struct {
			URL string `json:"url"`
		}{}
		err := json.Unmarshal([]byte(sc.Text()), &link)
		if err != nil {
			env.Channels.Bad <- errors.New("No results found")
			return
		}
		results = fmt.Sprintf("%s%s\n", results, link.URL)
	}
	env.Channels.Ok <- strings.TrimSpace(results)
}
