package error

import (
	"fmt"
	"os"
)

type ErrorLevel int

const (
	CRITICAL 	ErrorLevel = 10
	WARNING		ErrorLevel = 20
)

type MrRobotError struct {
	message string
	level	ErrorLevel
}

func NewError(message string, level ErrorLevel) *MrRobotError {
	return &MrRobotError{ message: message, level: level }
}

func (error MrRobotError) Show() {
	if error.level == CRITICAL {
		fmt.Printf("[!] %s\n", error.message)
		os.Exit(1)
	} else { fmt.Printf("[-] %s\n", error.message) }
}
