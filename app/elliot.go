package elliot

import (
	"github.com/cosasdepuma/elliot/app/tui"
)

/*

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

// EntrypointOld TODO: Doc
func EntrypointOld() {
	hasSubcommandError := false

	config.NewProgram("Elliot", "0.0.2")
	cli.Banner()

	modules := map[string]Subcommand{
		"nmap":        portscanner.Subcommand{},
		"portscanner": portscanner.Subcommand{},
		"robots":      robots.Subcommand{},
		"subdomain":   subdomain.Subcommand{},
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

*/

// Entrypoint TODO: Doc
func Entrypoint() {
	if err := tui.ShowTUI(); err != nil {
		err.Resolve(true)
	}
}

/*
	Documentation:
		- Subdomains: A lot of APIs (https://github.com/tomnomnom/assetfinder)
*/
