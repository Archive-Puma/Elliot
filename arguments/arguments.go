package arguments

import (
  "flag"
  "fmt"
  "os"
  "path/filepath"
)

type Arguments struct {
  // Parser
  argParse *flag.FlagSet
  // Special information
  Subcommand	string
  Version		string
  // Arguments
  Domain		string
}

// Constructor
func NewProgram(version string) *Arguments {
  arguments := &Arguments{ Version: version }
  if len(os.Args) < 3 { arguments.ShowHelp() }
  arguments.config()
  return arguments
}

func (arguments *Arguments) config() {
  arguments.argParse = flag.NewFlagSet("Arguments", flag.ContinueOnError)
  arguments.Subcommand = os.Args[1]

  arguments.argParse.StringVar(&arguments.Domain, "d", "", "Specify the domain")
  arguments.argParse.StringVar(&arguments.Domain, "domain", "", "Specify the domain")
  _ = arguments.argParse.Parse(os.Args[2:])
}

func (arguments Arguments) ShowHelp() {
  fmt.Printf("MrRobot v%s - Just another hacking framework\n\n", arguments.Version)
  fmt.Printf("Usage: %s [subcommand] <args...>\n", filepath.Base(os.Args[0]))
  fmt.Printf("Subcommands:\n\tsubdomain\tFind subdomains related to a given domain\n")
  fmt.Println("Arguments:")
  fmt.Println("\t-d, -domain\t\tSpecies a domain")
  os.Exit(1)
}
