package main

import (
	"github.com/mrrobotproject/mrrobot/arguments"
	"github.com/mrrobotproject/mrrobot/validator"

	"fmt"
)



func main() {
	args := arguments.NewProgram("0.0.1")

	switch args.Subcommand {
	case "subdomain":
		if validator.IsValidDomain(*args.Domain) {
			fmt.Println("[+] Valid")
		} else {
			fmt.Println("[!] Invalid")
		}
	default: args.ShowHelp()
	}
}
