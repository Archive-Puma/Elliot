package config

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var (
	// Args TODO: Doc
	Args = sArguments{}
)

type sArguments struct {
	// Parser
	argParse *flag.FlagSet
	// Special information
	Subcommand  string
	ProgramName string
	Version     string
	// Options
	DisplayHelp    bool
	DisplayVersion bool
	// Arguments
	Domain   string
	Disallow bool
	Extended bool
	URL      string
	Verbose  bool
}

// NewProgram TODO: Doc
func NewProgram(name string, version string) {
	Args.ProgramName = name
	Args.Version = version
	Args.argParse = flag.NewFlagSet("Arguments", flag.ExitOnError)
	Args.argParse.Usage = func() { fmt.Println(); ShowHelp() }

	Args.argParse.BoolVar(&Args.DisplayHelp, "h", false, "")
	Args.argParse.BoolVar(&Args.DisplayHelp, "help", false, "")

	Args.argParse.BoolVar(&Args.DisplayVersion, "version", false, "")

	Args.argParse.BoolVar(&Args.Disallow, "disallow", false, "")

	Args.argParse.StringVar(&Args.Domain, "d", "", "")
	Args.argParse.StringVar(&Args.Domain, "domain", "", "")

	Args.argParse.BoolVar(&Args.Extended, "extended", false, "")

	Args.argParse.StringVar(&Args.URL, "u", "", "")
	Args.argParse.StringVar(&Args.URL, "url", "", "")

	Args.argParse.BoolVar(&Args.Verbose, "v", false, "")
	Args.argParse.BoolVar(&Args.Verbose, "verbose", false, "")

	if len(os.Args) == 1 {
		ShowHelp()
	}

	Args.argParse.Parse(os.Args[1:])

	if len(os.Args) > 2 {
		Args.Subcommand = strings.ToLower(os.Args[1])
		_ = Args.argParse.Parse(os.Args[2:])
	}
}

// ShowHelp TODO: Doc
func ShowHelp() {
	fmt.Printf("%s v%s - Just another hacking framework\n\n", Args.ProgramName, Args.Version)
	fmt.Printf("Usage: %s [subcommand] <args...>\n", filepath.Base(os.Args[0]))
	fmt.Println("Subcommands:")
	Args.Print("help, man", "Show help message for a given subcommand")
	Args.Print("robots", "Find and displays robots.txt file from an URL")
	Args.Print("subdomain", "Find subdomains related to a domain")
	fmt.Println("Options:")
	Args.Print("-h, -help", "Display this message")
	Args.Print("-version", "Display the version")
	fmt.Println("Arguments:")
	Args.Print("-d, -domain", "Target domain")
	Args.Print("-u, -url", "Target URL")
	Args.Print("-v, -verbose", "Verbosity")
	os.Exit(1)
}

// ShowVersion TODO: Doc
func ShowVersion() {
	fmt.Printf("%s v%s\n", Args.ProgramName, Args.Version)
	os.Exit(0)
}

// Print TODO: Doc
func (args sArguments) Print(command string, description string) {
	fmt.Printf("    %-20s%s\n", command, description)
}
