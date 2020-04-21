package main

import (
	"github.com/mrrobotproject/mrrobot/arguments"

	"fmt"
)



func main() {
	args := arguments.NewProgram("0.0.1")

	switch args.Subcommand {
	case "subdomain": fmt.Println(*args.Domain)
	default: args.ShowHelp()
	}
}
