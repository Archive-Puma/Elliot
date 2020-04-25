package elliot

import (
	"github.com/cosasdepuma/elliot/app/arguments"
	"github.com/cosasdepuma/elliot/app/error"
	"github.com/cosasdepuma/elliot/app/robots"
	"github.com/cosasdepuma/elliot/app/subdomain"
	"github.com/cosasdepuma/elliot/app/tui"
	"github.com/cosasdepuma/elliot/app/validator"

	"fmt"
	"strconv"
)

// Entrypoint TODO: Doc
func Entrypoint() {
	args := arguments.NewProgram("Elliot", "0.0.2")
	tui.Banner(&args.ProgramName, &args.Version)

	switch args.Subcommand {
	case "robots":
		if !validator.IsValidURL(args.URL) {
			error.NewCritical("A valid url should be specified").Resolve()
		}

		tui.PrintInfo("Subcommand", args.Subcommand)
		tui.PrintInfo("URL", args.URL)
		tui.PrintInfo("Verbose", strconv.FormatBool(args.Verbose))
		now := tui.StartTime(&args.Subcommand)
		robots, err := robots.FindRobots(args.URL)
		if err != nil && args.Verbose {
			err.Resolve()
		}
		fmt.Println(robots)
		tui.EndTime(&now)
	case "subdomain":
		if !validator.IsValidDomain(args.Domain) {
			error.NewCritical("A valid domain should be specified").Resolve()
		}

		tui.PrintInfo("Subcommand", args.Subcommand)
		tui.PrintInfo("Domain", args.Domain)
		tui.PrintInfo("Verbose", strconv.FormatBool(args.Verbose))
		now := tui.StartTime(&args.Subcommand)
		subdomains := subdomain.FindAllConcurrent(args.Domain, args.Verbose)
		for _, sDomain := range subdomains {
			fmt.Println(sDomain)
		}
		tui.EndTime(&now)
	default:
		error.NewWarning("A valid subcommand should be specified").Resolve()
		fmt.Println()
		args.ShowHelp()
	}
}

/*
	Documentation:
		- Subdomains: A lot of APIs (https://github.com/tomnomnom/assetfinder)
*/
