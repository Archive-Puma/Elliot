package subdomainer

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type formatCrtSh struct {
	Name string `json:"name_value"`
}

func fetchCrtSh(domain string) ([]formatCrtSh, error) {
	// Compose the URL
	url := fmt.Sprintf("https://crt.sh/?q=%s&output=json", domain)
	// Request the data
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != 200 {
		return nil, errors.New("Crt.sh is not available")
	}
	// Grab the content
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("Crt.sh does not respond correctly")
	}
	// Parse the JSON
	subdomains := make([]formatCrtSh, 0)
	err = json.Unmarshal([]byte(body), &subdomains)
	if err != nil {
		return nil, errors.New("Bad JSON format using Crt.sh")
	}
	// Return the JSON
	return subdomains, nil
}

func filterCrtSh(data []formatCrtSh) []string {
	var subdomains []string
	duplicates := make(map[string]int)
	// Iterate over all the subdomains
	for _, sub := range data {
		splitter := strings.Split(sub.Name, "\n")
		subdomain := splitter[len(splitter)-1]

		duplicates[subdomain]++
		if duplicates[subdomain] == 1 {
			subdomains = append(subdomains, subdomain)
		}
	}
	return subdomains
}

func methodCtrSh(domain string) ([]string, error) {
	data, err := fetchCrtSh(domain)
	if err != nil {
		return nil, err
	}
	return filterCrtSh(data), nil
}
