package elliot

import "github.com/cosasdepuma/elliot/app/modules"

// RunDomainOSINT obtains the information related to a domain using Open Sources
func RunDomainOSINT(domain string) {
	// Set the target
	Backend.DB.SetDomain(domain)
	// Initialize the channels
	xsubdomains := make(chan []string, 1)
	xwhois := make(chan *modules.SWhois, 1)
	// Close the channels
	defer close(xsubdomains)
	defer close(xwhois)
	// Run the concurrent methods
	go modules.Subdomains(domain, &xsubdomains)
	go modules.Whois(domain, &xwhois)
	// Receive the data
	i := 2 // FIXME Better implementation using arrays or maps?
	for i > 0 {
		select {
		case subdomains := <-xsubdomains:
			Backend.DB.SetSubdomains(subdomains)
		case whois := <-xwhois:
			Backend.DB.SetWhois(whois)
		}
		i--
	}
}
