package subdomain

import (
	"github.com/cosasdepuma/elliot/error"

	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type formatCrtSh struct { Name string `json:"name_value"` }

func fetchCrtSh(domain string) ([]formatCrtSh, *error.MrRobotError) {
	// Compose the URL
	url := fmt.Sprintf("https://crt.sh/?q=%s&output=json", domain)
	// Request the data
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != 200 { return nil, error.NewWarning("Crt.sh is not available") }
	// Grab the content
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil { return nil, error.NewWarning("Crt.sh does not respond correctly") }
	// Parse the JSON
	subdomains := make([]formatCrtSh, 0)
	err = json.Unmarshal([]byte(body), &subdomains)
	if err != nil { return nil, error.NewWarning("Bad JSON format using Crt.sh") }
	// Return the JSON
	return subdomains, nil
}

func filterCrtSh(data []formatCrtSh) []string {
	var subdomains []string
	duplicates := make(map[string]int)
	// Iterate over all the subdomains
	for _, sub := range data {
		splitter := strings.Split(sub.Name, "\n")
		subdomain := splitter[len(splitter) - 1]

		duplicates[subdomain]++
		if duplicates[subdomain] == 1 { subdomains = append(subdomains, subdomain) }
	}
	return subdomains
}

func MethodCtrSh(domain string) ([]string, *error.MrRobotError) {
	data, err := fetchCrtSh(domain)
	if err != nil { return nil, err }
	return filterCrtSh(data), nil
}