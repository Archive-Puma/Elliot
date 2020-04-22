package subdomain

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func fetchHackerTarget(domain string) []string {
	// Compose the URL
	url := fmt.Sprintf("https://api.hackertarget.com/hostsearch/?q=%s", domain)
	// Request the data
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != 200 { return nil }
	// Grab the content
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil { return nil }
	// Parse the Response
	subdomains := make([]string, 0)
	sc := bufio.NewScanner(bytes.NewReader(body))
	for sc.Scan() {
		splitter := strings.SplitN(sc.Text(), ",", 2)
		subdomains = append(subdomains, splitter[0])
	}
	return subdomains
}

func filterHackerTarget(data []string) []string {
	var subdomains []string
	duplicates := make(map[string]int)
	// Iterate over all the subdomains
	for _, subdomain := range data {
		duplicates[subdomain]++
		if duplicates[subdomain] == 1 { subdomains = append(subdomains, subdomain) }
	}
	return subdomains
}

func MethodHackerTarget(domain string) []string {
	data := fetchHackerTarget(domain)
	return filterHackerTarget(data)
}