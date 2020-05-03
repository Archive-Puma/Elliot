package elliot

import (
	"github.com/cosasdepuma/elliot/app/cli"
	"github.com/cosasdepuma/elliot/app/config"
	"github.com/cosasdepuma/elliot/app/error"

	"github.com/cosasdepuma/elliot/app/subcommand/robots"
	"github.com/cosasdepuma/elliot/app/subcommand/subdomain"

	"fmt"
	"os"
)

func startProcess(subcommand Subcommand) {
	now := cli.StartTime(&config.Args.Subcommand)
	cli.PrintInfo("Subcommand", config.Args.Subcommand)
	err := subcommand.Check()
	if err != nil {
		err.Resolve(true)
	}
	cli.Separator()
	if err != nil {
		os.Exit(1)
	}
	results, errors := subcommand.Run()
	for _, err := range errors {
		err.Resolve(true)
	}
	for _, result := range results {
		fmt.Println(result)
	}
	cli.EndTime(&now)
}

// Entrypoint TODO: Doc
func Entrypoint() {
	hasSubcommandError := false

	config.NewProgram("Elliot", "0.0.2")
	cli.Banner()

	modules := map[string]Subcommand{
		"robots":    robots.Subcommand{},
		"subdomain": subdomain.Subcommand{},
	}

	if config.Args.Subcommand == "help" || config.Args.Subcommand == "man" {
		if subcommand, ok := modules[os.Args[2]]; ok {
			cli.Separator()
			cli.PrintInfo("Subcommand", os.Args[2])
			cli.Separator()
			fmt.Println("Arguments:")
			subcommand.Help()
			cli.Separator()
		} else {
			hasSubcommandError = true
		}
	} else if subcommand, ok := modules[config.Args.Subcommand]; ok {
		startProcess(subcommand)
	} else {
		hasSubcommandError = true
	}

	if hasSubcommandError {
		error.NewWarning("A valid subcommand should be specified").Resolve(true)
		fmt.Println()
		config.ShowHelp()
	}
}

/*
	Documentation:
		- Subdomains: A lot of APIs (https://github.com/tomnomnom/assetfinder)
*/
