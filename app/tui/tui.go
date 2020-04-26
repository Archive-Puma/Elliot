package tui

import (
	"fmt"
	"time"

	"github.com/cosasdepuma/elliot/app/config"
)

// Separator TODO: Doc
func Separator() {
	fmt.Println("===============================================================")
}

// Banner TODO: Doc
func Banner() {
	Separator()
	fmt.Printf("%s v%s\n", config.Args.ProgramName, config.Args.Version)
}

// StartTime TODO: Doc
func StartTime(process *string) time.Time {
	now := time.Now()
	Separator()
	fmt.Printf("%s\tStarting %s process\n", now.Format(time.Kitchen), *process)
	Separator()
	return now
}

// EndTime TODO: Doc
func EndTime(start *time.Time) time.Time {
	now := time.Now()
	Separator()
	fmt.Printf("%s\tFinished in %s\n", time.Now().Format(time.Kitchen), time.Since(*start))
	Separator()
	return now
}

// PrintInfo TODO: Doc
func PrintInfo(info string, message string) {
	fmt.Printf("[+] %-10s : %s\n", info, message)
}
