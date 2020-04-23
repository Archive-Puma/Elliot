package elliot

import (
	"github.com/cosasdepuma/elliot/app/arguments"
	"github.com/cosasdepuma/elliot/app/error"
	"github.com/cosasdepuma/elliot/app/subdomain"
	"github.com/cosasdepuma/elliot/app/tui"
	"github.com/cosasdepuma/elliot/app/validator"

	"fmt"
)

func Entrypoint() {
	args := arguments.NewProgram("Elliot", "0.0.1")
	tui.Banner(&args.ProgramName, &args.Version)

	switch args.Subcommand {
	case "subdomain":
		if ! validator.IsValidDomain(args.Domain) {
			error.NewCritical("A valid domain should be specified").Resolve()
		} else {
			tui.PrintInfo("Subcommand", args.Subcommand)
			tui.PrintInfo("Domain", args.Domain)
			now := tui.StartTime(&args.Subcommand)
			subdomains := subdomain.GetAllConcurrent(args.Domain)
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