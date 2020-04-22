package subdomain

import (
	"github.com/mrrobotproject/mrrobot/error"

	"sync"
)

type methodFunc func(string) ([]string, *error.MrRobotError)

func concurrentMethod(method methodFunc, domain string, wg sync.WaitGroup, channel *chan []string) {
	defer wg.Done()
	subdomains, err := method(domain)
	if err != nil {
		err.Resolve()
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
		if duplicates[subdomain] == 1 { subdomains = append(subdomains, subdomain) }
	}
	return subdomains
}

func GetAllConcurrent(domain string) []string {
	nMethods := 2
	subdomains := make([]string, 0)
	wg := sync.WaitGroup{}
	channel := make(chan []string, 0)
	defer close(channel)

	wg.Add(nMethods)
	go concurrentMethod(MethodCtrSh, domain, wg, &channel)
	go concurrentMethod(MethodHackerTarget, domain, wg, &channel)

	for nMethods > 0 {
		nMethods--
		subdomains = append(subdomains, <- channel...)
	}

	return filterDuplicates(subdomains)
}