package subdomain

import (
	"fmt"

	"github.com/cosasdepuma/elliot/app/config"
	"github.com/cosasdepuma/elliot/app/error"
	"github.com/cosasdepuma/elliot/app/tui"
	"github.com/cosasdepuma/elliot/app/validator"

	"sync"
)

// Subcommand TODO: Doc
type Subcommand struct{}

type function func(string) ([]string, *error.MrRobotError)

func concurrentMethod(method function, wg *sync.WaitGroup, channel *chan []string) {
	defer wg.Done()
	subdomains, err := method(config.Args.Domain)
	if err != nil {
		err.Resolve(config.Args.Verbose)
		*channel <- nil
	} else {
		*channel <- subdomains
	}
}

func filterDuplicates(data []string) []string {
	var subdomains []string
	duplicates := make(map[string]int)
	// Iterate over all the subdomains
	for _, subdomain := range data {
		duplicates[subdomain]++
		if duplicates[subdomain] == 1 {
			subdomains = append(subdomains, subdomain)
		}
	}
	return subdomains
}

// Check TODO: Doc
func (s Subcommand) Check() *error.MrRobotError {
	if validator.IsValidDomain(config.Args.Domain) {
		tui.PrintInfo("Domain", config.Args.Domain)
	} else {
		return error.NewWarning("A valid domain should be specified")
	}

	return nil
}

// Run TODO: Doc
func (s Subcommand) Run() []*error.MrRobotError {
	errors := []*error.MrRobotError{}

	availableMethods := map[string]function{
		"crtsh":        methodCtrSh,
		"hackertarget": methodHackerTarget,
		"threatcrowd":  methodThreatCrowd,
	}

	nMethods := len(availableMethods)
	subdomains := make([]string, 0)
	wg := sync.WaitGroup{}
	channel := make(chan []string, 0)
	defer close(channel)

	wg.Add(nMethods)

	for _, method := range availableMethods {
		go concurrentMethod(method, &wg, &channel)
	}

	for nMethods > 0 {
		nMethods--
		subdomains = append(subdomains, <-channel...)
	}

	for _, subdomain := range filterDuplicates(subdomains) {
		fmt.Println(subdomain)
	}

	return errors
}
