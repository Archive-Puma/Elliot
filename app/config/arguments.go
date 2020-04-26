package config

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
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
	Domain  string
	URL     string
	Verbose bool
}

// NewProgram TODO: Doc
func NewProgram(name string, version string) {
	Args.ProgramName = name
	Args.Version = version
	Args.argParse = flag.NewFlagSet("Arguments", flag.ExitOnError)
	Args.argParse.Usage = func() { fmt.Println(); ShowHelp() }

	Args.argParse.BoolVar(&Args.DisplayHelp, "h", false, "Display this message")
	Args.argParse.BoolVar(&Args.DisplayHelp, "help", false, "Display this message")

	Args.argParse.BoolVar(&Args.DisplayVersion, "version", false, "Display the version")

	Args.argParse.StringVar(&Args.Domain, "d", "", "Specify the domain")
	Args.argParse.StringVar(&Args.Domain, "domain", "", "Specify the domain")

	Args.argParse.StringVar(&Args.URL, "u", "", "Specify an URL")
	Args.argParse.StringVar(&Args.URL, "url", "", "Specify an URL")

	Args.argParse.BoolVar(&Args.Verbose, "v", false, "Verbose output")
	Args.argParse.BoolVar(&Args.Verbose, "verbose", false, "Verbose output")

	if len(os.Args) == 1 {
		ShowHelp()
	}

	Args.argParse.Parse(os.Args[1:])

	if len(os.Args) > 2 {
		Args.Subcommand = os.Args[1]
		_ = Args.argParse.Parse(os.Args[2:])
	}
}

// ShowHelp TODO: Doc
func ShowHelp() {
	fmt.Printf("%s v%s - Just another hacking framework\n\n", Args.ProgramName, Args.Version)
	fmt.Printf("Usage: %s [subcommand] <args...>\n", filepath.Base(os.Args[0]))
	fmt.Println("Subcommands:")
	fmt.Println("\trobots\t\t\tFind and display robots.txt file from an URL")
	fmt.Println("\tsubdomain\t\tFind subdomains related to a given domain")
	fmt.Println("Options:")
	fmt.Println("\t-h, -help\t\tDisplay this message")
	fmt.Println("\t-version\t\tDisplay the version")
	fmt.Println("Arguments:")
	fmt.Println("\t-d, -domain\t\tSpecify a domain")
	fmt.Println("\t-u, -url\t\tSpecify an URL")
	fmt.Println("\t-v, -verbose\t\tVerbose output")
	os.Exit(1)
}

// ShowVersion TODO: Doc
func ShowVersion() {
	fmt.Printf("%s v%s\n", Args.ProgramName, Args.Version)
	os.Exit(0)
}
