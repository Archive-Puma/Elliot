package main

import (
	"github.com/mrrobotproject/mrrobot/arguments"
	"github.com/mrrobotproject/mrrobot/error"
	"github.com/mrrobotproject/mrrobot/subdomain"
	"github.com/mrrobotproject/mrrobot/tui"
	"github.com/mrrobotproject/mrrobot/validator"

	"fmt"
)

func main() {
	args := arguments.NewProgram("0.0.1")
	tui.Banner(&args.Version)

	switch args.Subcommand {
	case "subdomain":
		if ! validator.IsValidDomain(*args.Domain) {
			error.NewCritical("A valid domain should be specified").Resolve()
		} else {
			tui.PrintInfo("Subcommand", args.Subcommand)
			tui.PrintInfo("Domain", *args.Domain)
			now := tui.StartTime(&args.Subcommand)
			subdomains := subdomain.GetAllConcurrent(*args.Domain)
			for _, sDomain := range subdomains { fmt.Println(sDomain) }
			tui.EndTime(&now)
		}
	default: args.ShowHelp()
	}
}


/*
	Documentation:
		- Subdomains: A lot of APIs (https://github.com/tomnomnom/assetfinder)
 */
