package modules

// === IMPORTS ===

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/cosasdepuma/elliot/app/server/database"
)

// === PUBLIC METHODS ===

// RunDomain obtains the information related to a domain using Open Sources
func RunDomain(domain string, db *database.Database) {
	// Set the target
	db.SetDomain(domain)
	// Initialize the channels
	xsubdomains := make(chan []string, 1)
	xsitecheck := make(chan map[string]interface{}, 1)
	xwhois := make(chan *whois, 1)
	// Close the channels
	defer close(xsubdomains)
	defer close(xsitecheck)
	defer close(xwhois)
	// Run the concurrent methods
	go moduleSubdomains(domain, &xsubdomains)
	go moduleSiteCheck(domain, &xsitecheck)
	go moduleWhois(domain, &xwhois)
	// Receive the data
	i := 3 // FIXME Better implementation using arrays or maps?
	for i > 0 {
		select {
		case s := <-xsubdomains:
			db.SetDomainSubdomains(s)
		case sc := <-xsitecheck:
			if sc != nil {
				if value, ok := sc["ipv4"].(string); ok {
					db.SetDomainIPv4(value)
				}
				if value, ok := sc["ipv6"].(string); ok {
					db.SetDomainIPv6(value)
				}
				if value, ok := sc["url"].(string); ok {
					db.SetDomainWebUrl(value)
				}
				if value, ok := sc["rating"].(string); ok {
					db.SetDomainWebRating(value)
				}
				if value, ok := sc["server"].(string); ok {
					db.SetDomainWebServer(value)
				}
				if value, ok := sc["redirects"].([]string); ok {
					db.SetDomainWebRedirects(value)
				}
				if value, ok := sc["links"].([]string); ok {
					db.SetDomainWebLinks(value)
				}
				if value, ok := sc["js"].([]string); ok {
					db.SetDomainWebJS(value)
				}
			}
		case w := <-xwhois:
			db.SetDomainWhoisTLD(w.tld)
			db.SetDomainWhoisStatus(w.status)
			db.SetDomainWhoisCreated(w.created)
			db.SetDomainWhoisChanged(w.changed)
			db.SetDomainWhoisPhones(w.phone)
			db.SetDomainWhoisEmails(w.email)
		}
		i--
	}
}

// === PRIVATE METHODS ===

func apiGet(service string, base string, data string) ([]byte, error) {
	// Compose the URL
	api := fmt.Sprintf(base, data)
	// Request the data
	resp, err := http.Get(api)
	if err != nil || resp.StatusCode != 200 {
		return nil, fmt.Errorf("%s is not available", service)
	}
	// Grab the content
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("%s does not respond correctly", service)
	}
	return body, nil
}
