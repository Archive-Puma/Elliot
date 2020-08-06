package modules

import (
	"encoding/json"
)

// TODO: More info from Sucuri

// === MODULE METHOD ===

func ModuleSiteCheck(url string, channel *chan map[string]interface{}) {
	service := "Sucuri"
	// Data structure
	results := struct {
		Scan struct {
			Error string `json:"error"`
		} `json:"scan"`
		Warnings struct {
			ScanFailed []struct {
				Type string `json:"type"`
			} `json:"scan_failed"`
		} `json:"warnings"`

		Site struct {
			IP        []string `json:"ip"`
			URL       string   `json:"final_url"`
			Server    []string `json:"running_on"`
			Redirects []string `json:"redirects_to"`
		} `json:"site"`
		Spider struct {
			Links      []string `json:"urls"`
			JsLocal    []string `json:"js_local"`
			JSExternal []string `json:"js_external"`
		} `json:"links"`
		Rating struct {
			Total struct {
				Score string `json:"rating"`
			} `json:"total"`
		} `json:"ratings"`
	}{}
	// Get the data
	body, err := apiGet(service, "https://sitecheck.sucuri.net/api/v3/?scan=%s", url)
	if err != nil {
		*channel <- nil
		return
	}
	// Parse the JSON
	err = json.Unmarshal(body, &results)
	if err != nil || results.Scan.Error != "" || (len(results.Warnings.ScanFailed) > 0 && results.Warnings.ScanFailed[0].Type == "Scan failed") {
		*channel <- nil
		return
	}
	// Simplifies the data
	server := ""
	if len(results.Site.Server) > 0 {
		server = results.Site.Server[0]
	}
	ipv4, ipv6 := "", ""
	if len(results.Site.Server) > 0 {
		ipv4 = results.Site.IP[0]
	}
	if len(results.Site.Server) > 1 {
		ipv6 = results.Site.IP[1]
	}
	js := append(results.Spider.JsLocal, results.Spider.JSExternal...)
	*channel <- map[string]interface{}{
		"url":       results.Site.URL,
		"rating":    results.Rating.Total.Score,
		"server":    server,
		"redirects": results.Site.Redirects,
		"links":     results.Spider.Links,
		"js":        js,
		"ipv4":      ipv4,
		"ipv6":      ipv6,
	}
}

/*
// === IMPORTS ===

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



// === PUBLIC METHODS ===

// Subdomains is a concurrent method to obtain the sub-domains associated to a domain using different services.
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

func sucuri(domain string) ([]string, error) {
	// Compose the URL
	url := fmt.Sprintf("https://sitecheck.sucuri.net/api/v3/\?scan\=http://%s", domain)
	// Request the data
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != 200 {
		return nil, errors.New("Sucuri is not available")
	}
	// Grab the content
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("Sucuri does not respond correctly")
	}
	// Parse the JSON
	subdomains := struct {
		Ip []string `json:"subdomains"`
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
*/
