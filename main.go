package main

import (
	"github.com/mrrobotproject/mrrobot/arguments"
	"github.com/mrrobotproject/mrrobot/subdomain"
	"github.com/mrrobotproject/mrrobot/validator"

	"fmt"
)

func main() {
	args := arguments.NewProgram("0.0.1")

	switch args.Subcommand {
	case "subdomain":
		if ! validator.IsValidDomain(*args.Domain) {
			fmt.Println("[!] Domain should be specified")
		} else {
			fmt.Printf("[+] Results for domain %s: \n%v", *args.Domain, subdomain.MethodHackerTarget(*args.Domain))
		}

	default: args.ShowHelp()
	}
}



/*
	Documentation:
		- Subdomains: A lot of APIs (https://github.com/tomnomnom/assetfinder)
 */
