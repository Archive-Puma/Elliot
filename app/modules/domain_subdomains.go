package modules

// === IMPORTS ===

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
	"sync"

	"github.com/cosasdepuma/elliot/app/utils"
)

// === MODULE METHOD ===

func moduleSubdomains(domain string, output *chan []string) {
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
		go concurrentSubdomainer(method, domain, &wg, &channel)
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

// === PRIVATE METHODS ===

// ==== Subconcurrency Method ====

func concurrentSubdomainer(method func(string) ([]string, error), domain string, wg *sync.WaitGroup, channel *chan []string) {
	defer wg.Done()
	result, err := method(domain)
	if err != nil {
		*channel <- nil
	} else {
		*channel <- result
	}
}

// ==== Subdomain Methods ====

// FIXME: API count exceeded - Increase Quota with Membership

func subdomainsInThreatCrowd(domain string) ([]string, error) {
	service := "ThreadCrowd"
	// Data structure
	results := struct {
		Subdomains []string `json:"subdomains"`
	}{}
	// Get the data
	body, err := apiGet(service,
		"https://www.threatcrowd.org/searchApi/v2/domain/report/?domain=%s",
		domain)
	// Parse the JSON
	err = json.Unmarshal(body, &results)
	if err != nil {
		return nil, fmt.Errorf("Bad JSON format using %s", service)
	}
	return results.Subdomains, err
}

func subdomainsInHackerTarget(domain string) ([]string, error) {
	service := "HackerTarget"
	// Get the data
	body, err := apiGet(service,
		"https://api.hackertarget.com/hostsearch/?q=%s",
		domain)
	if err != nil {
		return nil, err
	}
	// Parse the data
	subdomains := make([]string, 0)
	sc := bufio.NewScanner(bytes.NewReader(body))
	for sc.Scan() {
		splitter := strings.SplitN(sc.Text(), ",", 2)
		subdomains = append(subdomains, splitter[0])
	}
	return subdomains, nil
}
