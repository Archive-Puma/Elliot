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
  ProgramName   string
  Version		string
  // Options
  Help          bool
  // Arguments
  Domain		string
}

// Constructor
func NewProgram(name string, version string) *Arguments {
  arguments := &Arguments{ ProgramName: name, Version: version }
  if len(os.Args) < 3 { arguments.ShowHelp() }
  arguments.config()
  return arguments
}

func (arguments *Arguments) config() {
  arguments.argParse = flag.NewFlagSet("Arguments", flag.ContinueOnError)
  arguments.Subcommand = os.Args[1]

  arguments.argParse.BoolVar(&arguments.Help, "h", false, "Display this message")
  arguments.argParse.BoolVar(&arguments.Help, "help", false, "Display this message")

  arguments.argParse.StringVar(&arguments.Domain, "d", "", "Specify the domain")
  arguments.argParse.StringVar(&arguments.Domain, "domain", "", "Specify the domain")
  _ = arguments.argParse.Parse(os.Args[2:])

  if arguments.Help { arguments.ShowHelp() }
}

func (arguments Arguments) ShowHelp() {
  fmt.Printf("MrRobot v%s - Just another hacking framework\n\n", arguments.Version)
  fmt.Printf("Usage: %s [subcommand] <args...>\n", filepath.Base(os.Args[0]))
  fmt.Printf("Subcommands:\n\tsubdomain\t\tFind subdomains related to a given domain\n")
  fmt.Println("Options:")
  fmt.Println("\t-h, -help\t\tDisplay this message")
  fmt.Println("Arguments:")
  fmt.Println("\t-d, -domain\t\tSpecify a domain")
  os.Exit(1)
}
