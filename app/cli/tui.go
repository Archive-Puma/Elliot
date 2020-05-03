package cli

import (
	"fmt"
	"time"

	"github.com/cosasdepuma/elliot/app/config"
)

// Separator TODO: Doc
func Separator() {
	if !config.Args.Bare {
		fmt.Println("===============================================================")
	}
}

// Banner TODO: Doc
func Banner() {
	if !config.Args.Bare {
		Separator()
		fmt.Printf("[Ɇ] %s v%s\n", config.Args.ProgramName, config.Args.Version)
	}
}

// StartTime TODO: Doc
func StartTime(process *string) time.Time {
	now := time.Now()
	if !config.Args.Bare {
		Separator()
		fmt.Printf("[◷] %-8sStarting %s process\n", now.Format(time.Kitchen), *process)
		Separator()
	}
	return now
}

// EndTime TODO: Doc
func EndTime(start *time.Time) time.Time {
	now := time.Now()
	if !config.Args.Bare {
		Separator()
		fmt.Printf("[◷] %-8sFinished in %s\n", time.Now().Format(time.Kitchen), time.Since(*start))
		Separator()
	}
	return now
}

// PrintInfo TODO: Doc
func PrintInfo(info string, message string) {
	if !config.Args.Bare {
		fmt.Printf("[+] %-10s : %s\n", info, message)
	}
}
