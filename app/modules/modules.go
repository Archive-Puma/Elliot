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
	xwhois := make(chan *whois, 1)
	// Close the channels
	defer close(xsubdomains)
	defer close(xwhois)
	// Run the concurrent methods
	go moduleSubdomains(domain, &xsubdomains)
	go moduleWhois(domain, &xwhois)
	// Receive the data
	i := 2 // FIXME Better implementation using arrays or maps?
	for i > 0 {
		select {
		case s := <-xsubdomains:
			db.SetDomainSubdomains(s)
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
