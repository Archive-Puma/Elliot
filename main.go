package main

import (
	"context"
	"fmt"
	"os"
	"sort"

	"github.com/cosasdepuma/elliot/pkg/app"
	"github.com/cosasdepuma/elliot/pkg/modules"
	"github.com/cosasdepuma/elliot/pkg/modules/scanner"
)

var subcommands = map[string]modules.Module{
	"nmap": scanner.Nmap,
}

func main() {
	// Define a context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	// Create a new core
	core := app.NewCore()
	core.ParentCtx = ctx

	// Check arguments
	if len(os.Args) == 1 {
		fmt.Fprintf(os.Stderr, "Usage: %s <module> [args...]\n", os.Args[0])
		os.Exit(1)
	}
	// Check subcommand
	if subcommand, ok := subcommands[os.Args[1]]; ok {
		subcommand.Flag.Parse(os.Args[2:])
		subcommand.Run(core)
	} else {
		values := make([]string, 0, len(subcommands))
		for value := range subcommands {
			values = append(values, value)
		}
		sort.Strings(values)
		fmt.Fprintf(os.Stderr, "[-] Available subcommands: %v\n", values)
	}
}
