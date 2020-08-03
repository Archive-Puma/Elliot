package modules

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"

	"github.com/cosasdepuma/elliot/app/utils"
)

func Subdomains(domain string, output *chan []string) {
	availableMethods := [](func(string) ([]string, error)){
		subdomainsInHackerTarget, subdomainsInThreatCrowd,
	}
	// Concurrency
	wg := sync.WaitGroup{}
	wg.Add(len(availableMethods))
	channel := make(chan []string, len(availableMethods))
	defer close(channel)

	// Initialize the concurrency
	for _, method := range availableMethods {
		go concurrent(method, domain, &wg, &channel)
	}

	// Retrieve the results
	subdomains, i := make([]string, 0), 0
	for i < len(availableMethods) {
		i++
		subdomains = append(subdomains, <-channel...)
	}

	// Filter the duplicates
	if len(subdomains) == 0 {
		*output <- nil
		return
	}

	subdomains = utils.FilterDuplicates(subdomains)
	*output <- utils.FilterDuplicates(subdomains)
}

func concurrent(method func(string) ([]string, error), domain string, wg *sync.WaitGroup, channel *chan []string) {
	defer wg.Done()
	result, err := method(domain)
	if err != nil {
		*channel <- nil
	} else {
		*channel <- result
	}
}

// ==== Subdomain Methods ====

func subdomainsInThreatCrowd(domain string) ([]string, error) {
	// Compose the URL
	url := fmt.Sprintf("https://www.threatcrowd.org/searchApi/v2/domain/report/?domain=%s", domain)
	// Request the data
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != 200 {
		return nil, errors.New("ThreatCrowd is not available")
	}
	// Grab the content
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("ThreatCrowd does not respond correctly")
	}
	// Parse the JSON
	subdomains := struct {
		Results []string `json:"subdomains"`
	}{}
	err = json.Unmarshal([]byte(body), &subdomains)
	if err != nil {
		return nil, errors.New("Bad JSON format using ThreatCrowd")
	}
	// Return the JSON
	return subdomains.Results, nil
}

func subdomainsInHackerTarget(domain string) ([]string, error) {
	// Compose the URL
	url := fmt.Sprintf("https://api.hackertarget.com/hostsearch/?q=%s", domain)
	// Request the data
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != 200 {
		return nil, errors.New("HackerTarget is not available")
	}
	// Grab the content
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("HackerTarget does not respond correctly")
	}
	// Parse the Response
	subdomains := make([]string, 0)
	sc := bufio.NewScanner(bytes.NewReader(body))
	for sc.Scan() {
		splitter := strings.SplitN(sc.Text(), ",", 2)
		subdomains = append(subdomains, splitter[0])
	}
	return subdomains, nil
}

/*
import (
	"errors"
	"fmt"
	"sync"

)

// Plugin allows it to be executed by Elliot
type  struct{}

type function func(string) ([]string, error)

// Check that all parameters are defined correctly
func (plgn Plugin) Check() error {
	if !validator.IsValidDomain(env.Config.Target) {
		return errors.New("A valid domain should be specified")
	}

	return nil
}

// Run is the entrypoint of the plugin
func (plgn Plugin) Run() {
	if err := plgn.Check(); err != nil {
		env.Channels.Bad <- err
		return
	}

	subdomains := make([]string, 0)

	availableMethods := map[string]function{
		"crtsh":        methodCtrSh,
		"hackertarget": methodHackerTarget,
		"threatcrowd":  methodThreatCrowd,
	}

	wg := sync.WaitGroup{}
	nMethods := len(availableMethods)
	channel := make(chan []string, 0)
	defer close(channel)

	wg.Add(nMethods)

	for _, method := range availableMethods {
		go concurrentMethod(method, env.Config.Target, &wg, &channel)
	}

	for nMethods > 0 {
		nMethods--
		subdomains = append(subdomains, <-channel...)
	}

	result := ""
	filtered := filterDuplicates(subdomains)
	for _, subdomain := range filtered {
		if len(subdomain) > 0 && subdomain != "error check your search parameter" {
			result = fmt.Sprintf("%s%s\n", result, subdomain)
		}
	}

	env.Channels.Ok <- result
	if err := plgn.Save(filtered); err != nil {
		logrus.Error(err)
	}
}

func concurrentMethod(method function, target string, wg *sync.WaitGroup, channel *chan []string) {
	defer wg.Done()
	subdomains, err := method(target)
	if err != nil {
		*channel <- nil
	} else {
		*channel <- subdomains
	}
}

func filterDuplicates(data []string) []string {
	var subdomains []string
	duplicates := make(map[string]int)
	// Iterate over all the subdomains
	for _, subdomain := range data {
		duplicates[subdomain]++
		if duplicates[subdomain] == 1 {
			subdomains = append(subdomains, subdomain)
		}
	}
	return subdomains
}

*/
