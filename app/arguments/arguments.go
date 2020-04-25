package arguments

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

// Arguments TODO: Doc
type Arguments struct {
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
func NewProgram(name string, version string) *Arguments {
	arguments := &Arguments{ProgramName: name, Version: version}
	arguments.argParse = flag.NewFlagSet("Arguments", flag.ContinueOnError)

	arguments.argParse.BoolVar(&arguments.DisplayHelp, "h", false, "Display this message")
	arguments.argParse.BoolVar(&arguments.DisplayHelp, "help", false, "Display this message")

	arguments.argParse.BoolVar(&arguments.DisplayVersion, "version", false, "Display the version")

	if len(os.Args) < 2 {
		arguments.ShowHelp()
	}

	arguments.argParse.Parse(os.Args[1:])

	if len(os.Args) > 2 {
		arguments.config()
	}

	if arguments.DisplayHelp {
		arguments.ShowHelp()
	} else if arguments.DisplayVersion {
		arguments.ShowVersion()
	}

	return arguments
}

func (arguments *Arguments) config() {
	arguments.Subcommand = os.Args[1]

	arguments.argParse.StringVar(&arguments.Domain, "d", "", "Specify the domain")
	arguments.argParse.StringVar(&arguments.Domain, "domain", "", "Specify the domain")

	arguments.argParse.StringVar(&arguments.URL, "u", "", "Specify an URL")
	arguments.argParse.StringVar(&arguments.URL, "url", "", "Specify an URL")

	arguments.argParse.BoolVar(&arguments.Verbose, "v", false, "Verbose output")
	arguments.argParse.BoolVar(&arguments.Verbose, "verbose", false, "Verbose output")

	_ = arguments.argParse.Parse(os.Args[2:])
}

// ShowHelp TODO: Doc
func (arguments Arguments) ShowHelp() {
	fmt.Printf("%s v%s - Just another hacking framework\n\n", arguments.ProgramName, arguments.Version)
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
func (arguments Arguments) ShowVersion() {
	fmt.Printf("%s v%s\n", arguments.ProgramName, arguments.Version)
	os.Exit(0)
}
