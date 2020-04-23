package tui

import (
	"fmt"
	"time"
)

func Separator() {
	fmt.Println("===============================================================")
}

func Banner(name *string, version *string) {
	Separator()
	fmt.Printf("%s v%s\n", *name, *version)
	Separator()
}

func StartTime(process *string) time.Time {
	now := time.Now()
	Separator()
	fmt.Printf("%s\tStarting %s process\n", now.Format(time.Kitchen), *process)
	Separator()
	return now
}

func EndTime(start *time.Time) time.Time {
	now := time.Now()
	Separator()
	fmt.Printf("%s\tFinished in %s\n", time.Now().Format(time.Kitchen), time.Since(*start))
	Separator()
	return now
}

func PrintInfo(info string, message string) {
	fmt.Printf("[+] %-10s : %s\n", info, message)
}
