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

func NewError(level ErrorLevel, message string, a ...interface{}) *MrRobotError {
	msg := fmt.Sprintf(message, a...)
	return &MrRobotError{ message: msg, level: level }
}

func NewWarning(message string, a ...interface{}) *MrRobotError {
	return NewError(WARNING, message, a...)
}

func NewCritical(message string, a ...interface{}) *MrRobotError {
	return NewError(CRITICAL, message, a...)
}

func (error MrRobotError) Resolve() {
	if error.level == CRITICAL {
		fmt.Printf("[!] %s\n", error.message)
		os.Exit(1)
	} else { fmt.Printf("[-] %s\n", error.message) }
}
