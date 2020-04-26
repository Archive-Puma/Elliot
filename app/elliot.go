package elliot

import (
	"github.com/cosasdepuma/elliot/app/config"
	"github.com/cosasdepuma/elliot/app/error"
	"github.com/cosasdepuma/elliot/app/robots"
	"github.com/cosasdepuma/elliot/app/tui"

	"fmt"
	"os"
)

func startProcess(subcommand Subcommand) {
	now := tui.StartTime(&config.Args.Subcommand)
	tui.PrintInfo("Subcommand", config.Args.Subcommand)
	err := subcommand.Check()
	if err != nil {
		err.Resolve(true)
	}
	tui.Separator()
	if err != nil {
		os.Exit(1)
	}
	errors := subcommand.Run()
	for _, err := range errors {
		err.Resolve(true)
	}
	tui.EndTime(&now)
}

// Entrypoint TODO: Doc
func Entrypoint() {
	config.NewProgram("Elliot", "0.0.2")
	tui.Banner()

	switch config.Args.Subcommand {
	case "robots":
		startProcess(robots.Subcommand{})
	default:
		error.NewWarning("A valid subcommand should be specified").Resolve(true)
		fmt.Println()
		config.ShowHelp()
	}
}

/*
	Documentation:
		- Subdomains: A lot of APIs (https://github.com/tomnomnom/assetfinder)
*/
