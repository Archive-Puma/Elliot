package subdomain

import (
	"errors"
	"fmt"
	"sync"

	"github.com/cosasdepuma/elliot/app/env"
	"github.com/cosasdepuma/elliot/app/validator"
)

// Plugin TODO: Doc
type Plugin struct{}

type function func(string) ([]string, error)

// Check TODO: Doc
func (plgn Plugin) Check() error {
	if !validator.IsValidDomain(env.Params.Target) {
		return errors.New("A valid domain should be specified")
	}

	return nil
}

// Run TODO: Doc
// -- Fixme: Can't rerun module
func (plgn Plugin) Run() {
	subdomains := make([]string, 0)

	availableMethods := map[string]function{
		"crtsh":        methodCtrSh,
		"hackertarget": methodHackerTarget,
		"threatcrowd":  methodThreatCrowd,
	}

	wg := sync.WaitGroup{}
	nMethods := len(availableMethods)
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

	result := ""
	for _, subdomain := range filterDuplicates(subdomains) {
		result = fmt.Sprintf("%s\n%s", result, subdomain)
	}

	env.Channels.Ok <- result
}

func concurrentMethod(method function, wg *sync.WaitGroup, channel *chan []string) {
	defer wg.Done()
	subdomains, err := method(env.Params.Target)
	if err != nil {
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
