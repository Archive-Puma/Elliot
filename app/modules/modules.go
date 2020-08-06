package modules

import (
	"github.com/cosasdepuma/elliot/app/server/database"
)

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
