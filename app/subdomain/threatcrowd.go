package subdomain

import (
	"github.com/cosasdepuma/elliot/app/error"

	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type formatThreatCrowd struct {
	Subdomains []string `json:"subdomains"`
}

func fetchThreatCrowd(domain string) ([]string, *error.MrRobotError) {
	// Compose the URL
	url := fmt.Sprintf("https://www.threatcrowd.org/searchApi/v2/domain/report/?domain=%s", domain)
	// Request the data
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != 200 {
		return nil, error.NewWarning("ThreatCrowd is not available")
	}
	// Grab the content
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, error.NewWarning("ThreatCrowd does not respond correctly")
	}
	// Parse the JSON
	subdomains := formatThreatCrowd{}
	err = json.Unmarshal([]byte(body), &subdomains)
	if err != nil {
		return nil, error.NewWarning("Bad JSON format using ThreatCrowd")
	}
	// Return the JSON
	return subdomains.Subdomains, nil
}

func methodThreatCrowd(domain string) ([]string, *error.MrRobotError) {
	data, err := fetchThreatCrowd(domain)
	if err != nil {
		return nil, err
	}
	return data, nil
}
